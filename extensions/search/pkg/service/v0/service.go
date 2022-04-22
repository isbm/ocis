package service

import (
	"context"
	"errors"

	"github.com/asim/go-micro/plugins/events/natsjs/v4"
	"github.com/blevesearch/bleve/v2"
	revactx "github.com/cs3org/reva/v2/pkg/ctx"
	"github.com/cs3org/reva/v2/pkg/events"
	"github.com/cs3org/reva/v2/pkg/events/server"
	"github.com/cs3org/reva/v2/pkg/rgrpc/todo/pool"
	"go-micro.dev/v4/metadata"
	grpcmetadata "google.golang.org/grpc/metadata"

	"github.com/owncloud/ocis/extensions/audit/pkg/types"
	"github.com/owncloud/ocis/extensions/search/pkg/config"
	"github.com/owncloud/ocis/extensions/search/pkg/search"
	"github.com/owncloud/ocis/extensions/search/pkg/search/index"
	searchprovider "github.com/owncloud/ocis/extensions/search/pkg/search/provider"
	"github.com/owncloud/ocis/ocis-pkg/log"
	searchsvc "github.com/owncloud/ocis/protogen/gen/ocis/services/search/v0"
)

// NewHandler returns a service implementation for Service.
func NewHandler(opts ...Option) (searchsvc.SearchProviderHandler, error) {
	options := newOptions(opts...)
	logger := options.Logger
	cfg := options.Config

	// Connect to nats to listen for changes that need to trigger an index update
	evtsCfg := cfg.Events
	client, err := server.NewNatsStream(
		natsjs.Address(evtsCfg.Endpoint),
		natsjs.ClusterID(evtsCfg.Cluster),
	)
	if err != nil {
		return nil, err
	}
	evts, err := events.Consume(client, evtsCfg.ConsumerGroup, types.RegisteredEvents()...)
	if err != nil {
		return nil, err
	}

	bleveIndex, err := bleve.NewMemOnly(index.BuildMapping())
	if err != nil {
		return nil, err
	}
	index, err := index.New(bleveIndex)
	if err != nil {
		return nil, err
	}

	gwclient, err := pool.GetGatewayServiceClient(cfg.Reva.Address)
	if err != nil {
		logger.Fatal().Err(err).Str("addr", cfg.Reva.Address).Msg("could not get reva client")
	}

	provider := searchprovider.New(gwclient, index, cfg.MachineAuthAPIKey, evts)

	return &Service{
		id:       cfg.GRPC.Namespace + "." + cfg.Service.Name,
		log:      logger,
		Config:   cfg,
		provider: provider,
	}, nil
}

// Service implements the searchServiceHandler interface
type Service struct {
	id       string
	log      log.Logger
	Config   *config.Config
	provider search.ProviderClient
}

func (s Service) Search(ctx context.Context, in *searchsvc.SearchRequest, out *searchsvc.SearchResponse) error {
	// Get token from the context (go-micro) and make it known to the reva client too (grpc)
	t, ok := metadata.Get(ctx, revactx.TokenHeader)
	if !ok {
		s.log.Error().Msg("Could not get token from context")
		return errors.New("could not get token from context")
	}
	ctx = grpcmetadata.AppendToOutgoingContext(ctx, revactx.TokenHeader, t)

	res, err := s.provider.Search(ctx, &searchsvc.SearchRequest{
		Query: in.Query,
	})
	if err != nil {
		return err
	}

	out.Matches = res.Matches
	out.NextPageToken = res.NextPageToken
	return nil
}
