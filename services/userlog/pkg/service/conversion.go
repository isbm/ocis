package service

import (
	"bytes"
	"context"
	"errors"
	"text/template"
	"time"

	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	user "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	collaboration "github.com/cs3org/go-cs3apis/cs3/sharing/collaboration/v1beta1"
	storageprovider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	"github.com/cs3org/reva/v2/pkg/events"
	"github.com/cs3org/reva/v2/pkg/storagespace"
	"github.com/cs3org/reva/v2/pkg/utils"
	"github.com/leonelquinteros/gotext"
	ehmsg "github.com/owncloud/ocis/v2/protogen/gen/ocis/messages/eventhistory/v0"
)

var (
	_resourceTypeSpace = "storagespace"
	_resourceTypeShare = "share"

	// TODO: from config
	_pathToLocales = "/home/jkoberg/ocis/services/userlog/pkg/service/locales"
)

// OC10Notification is the oc10 style representation of an event
// some fields are left out for simplicity
type OC10Notification struct {
	EventID        string                 `json:"notification_id"`
	Service        string                 `json:"app"`
	UserName       string                 `json:"user"`
	Timestamp      string                 `json:"datetime"`
	ResourceID     string                 `json:"object_id"`
	ResourceType   string                 `json:"object_type"`
	Subject        string                 `json:"subject"`
	SubjectRaw     string                 `json:"subjectRich"`
	Message        string                 `json:"message"`
	MessageRaw     string                 `json:"messageRich"`
	MessageDetails map[string]interface{} `json:"messageRichParameters"`
}

// Converter is responsible for converting eventhistory events to OC10Notifications
type Converter struct {
	locale            string
	gwClient          gateway.GatewayAPIClient
	machineAuthAPIKey string
	serviceName       string
	registeredEvents  map[string]events.Unmarshaller

	// cached within one request not query other service too much
	spaces    map[string]*storageprovider.StorageSpace
	users     map[string]*user.User
	resources map[string]*storageprovider.ResourceInfo
	contexts  map[string]context.Context
}

// NewConverter returns a new Converter
func NewConverter(loc string, gwc gateway.GatewayAPIClient, machineAuthAPIKey string, name string, registeredEvents map[string]events.Unmarshaller) *Converter {
	return &Converter{
		locale:            loc,
		gwClient:          gwc,
		machineAuthAPIKey: machineAuthAPIKey,
		serviceName:       name,
		registeredEvents:  registeredEvents,
		spaces:            make(map[string]*storageprovider.StorageSpace),
		users:             make(map[string]*user.User),
		resources:         make(map[string]*storageprovider.ResourceInfo),
		contexts:          make(map[string]context.Context),
	}
}

// ConvertEvent converts an eventhistory event to an OC10Notification
func (c *Converter) ConvertEvent(event *ehmsg.Event) (OC10Notification, error) {
	etype, ok := c.registeredEvents[event.Type]
	if !ok {
		// this should not happen
		return OC10Notification{}, errors.New("eventtype not registered")
	}

	einterface, err := etype.Unmarshal(event.Event)
	if err != nil {
		// this shouldn't happen either
		return OC10Notification{}, errors.New("cant unmarshal event")
	}

	switch ev := einterface.(type) {
	default:
		return OC10Notification{}, errors.New("unknown event type")
	// space related
	case events.SpaceDisabled:
		return c.spaceMessage(event.Id, SpaceDisabled, ev.Executant, ev.ID.GetOpaqueId(), ev.Timestamp)
	case events.SpaceDeleted:
		return c.spaceDeletedMessage(event.Id, ev.Executant, ev.ID.GetOpaqueId(), ev.SpaceName, ev.Timestamp)
	case events.SpaceShared:
		return c.spaceMessage(event.Id, SpaceShared, ev.Executant, ev.ID.GetOpaqueId(), ev.Timestamp)
	case events.SpaceUnshared:
		return c.spaceMessage(event.Id, SpaceUnshared, ev.Executant, ev.ID.GetOpaqueId(), ev.Timestamp)
	case events.SpaceMembershipExpired:
		return c.spaceMessage(event.Id, SpaceMembershipExpired, ev.SpaceOwner, ev.SpaceID.GetOpaqueId(), ev.ExpiredAt)

	// share related
	case events.ShareCreated:
		return c.shareMessage(event.Id, ShareCreated, ev.Executant, ev.ItemID, ev.ShareID, utils.TSToTime(ev.CTime))
	case events.ShareExpired:
		return c.shareMessage(event.Id, ShareExpired, ev.ShareOwner, ev.ItemID, ev.ShareID, ev.ExpiredAt)
	case events.ShareRemoved:
		return c.shareMessage(event.Id, ShareRemoved, ev.Executant, ev.ItemID, ev.ShareID, ev.Timestamp)
	}
}

func (c *Converter) spaceDeletedMessage(eventid string, executant *user.UserId, spaceid string, spacename string, ts time.Time) (OC10Notification, error) {
	usr, err := c.getUser(context.Background(), executant)
	if err != nil {
		return OC10Notification{}, err
	}

	subj, subjraw, msg, msgraw, err := composeMessage(SpaceDeleted, c.locale, map[string]interface{}{
		"username":  usr.GetDisplayName(),
		"spacename": spacename,
	})
	if err != nil {
		return OC10Notification{}, err
	}

	space := &storageprovider.StorageSpace{Id: &storageprovider.StorageSpaceId{OpaqueId: spaceid}, Name: spacename}

	return OC10Notification{
		EventID:        eventid,
		Service:        c.serviceName,
		UserName:       usr.GetUsername(),
		Timestamp:      ts.Format(time.RFC3339Nano),
		ResourceID:     spaceid,
		ResourceType:   _resourceTypeSpace,
		Subject:        subj,
		SubjectRaw:     subjraw,
		Message:        msg,
		MessageRaw:     msgraw,
		MessageDetails: generateDetails(usr, space, nil, nil),
	}, nil
}

func (c *Converter) spaceMessage(eventid string, nt NotificationTemplate, executant *user.UserId, spaceid string, ts time.Time) (OC10Notification, error) {
	usr, err := c.getUser(context.Background(), executant)
	if err != nil {
		return OC10Notification{}, err
	}

	ctx, err := c.authenticate(usr)
	if err != nil {
		return OC10Notification{}, err
	}

	space, err := c.getSpace(ctx, spaceid)
	if err != nil {
		return OC10Notification{}, err
	}

	subj, subjraw, msg, msgraw, err := composeMessage(nt, c.locale, map[string]interface{}{
		"username":  usr.GetDisplayName(),
		"spacename": space.GetName(),
	})
	if err != nil {
		return OC10Notification{}, err
	}

	return OC10Notification{
		EventID:        eventid,
		Service:        c.serviceName,
		UserName:       usr.GetUsername(),
		Timestamp:      ts.Format(time.RFC3339Nano),
		ResourceID:     spaceid,
		ResourceType:   _resourceTypeSpace,
		Subject:        subj,
		SubjectRaw:     subjraw,
		Message:        msg,
		MessageRaw:     msgraw,
		MessageDetails: generateDetails(usr, space, nil, nil),
	}, nil
}

func (c *Converter) shareMessage(eventid string, nt NotificationTemplate, executant *user.UserId, resourceid *storageprovider.ResourceId, shareid *collaboration.ShareId, ts time.Time) (OC10Notification, error) {
	usr, err := c.getUser(context.Background(), executant)
	if err != nil {
		return OC10Notification{}, err
	}

	ctx, err := c.authenticate(usr)
	if err != nil {
		return OC10Notification{}, err
	}

	info, err := c.getResource(ctx, resourceid)
	if err != nil {
		return OC10Notification{}, err
	}

	subj, subjraw, msg, msgraw, err := composeMessage(nt, c.locale, map[string]interface{}{
		"username":     usr.GetDisplayName(),
		"resourcename": info.GetName(),
	})
	if err != nil {
		return OC10Notification{}, err
	}

	return OC10Notification{
		EventID:        eventid,
		Service:        c.serviceName,
		UserName:       usr.GetUsername(),
		Timestamp:      ts.Format(time.RFC3339Nano),
		ResourceID:     storagespace.FormatResourceID(*info.GetId()),
		ResourceType:   _resourceTypeShare,
		Subject:        subj,
		SubjectRaw:     subjraw,
		Message:        msg,
		MessageRaw:     msgraw,
		MessageDetails: generateDetails(usr, nil, info, shareid),
	}, nil
}

func (c *Converter) authenticate(usr *user.User) (context.Context, error) {
	if ctx, ok := c.contexts[usr.GetId().GetOpaqueId()]; ok {
		return ctx, nil
	}
	ctx, err := authenticate(usr, c.gwClient, c.machineAuthAPIKey)
	if err == nil {
		c.contexts[usr.GetId().GetOpaqueId()] = ctx
	}
	return ctx, err
}

func (c *Converter) getSpace(ctx context.Context, spaceID string) (*storageprovider.StorageSpace, error) {
	if space, ok := c.spaces[spaceID]; ok {
		return space, nil
	}
	space, err := getSpace(ctx, spaceID, c.gwClient)
	if err == nil {
		c.spaces[spaceID] = space
	}
	return space, err
}

func (c *Converter) getResource(ctx context.Context, resourceID *storageprovider.ResourceId) (*storageprovider.ResourceInfo, error) {
	if r, ok := c.resources[resourceID.GetOpaqueId()]; ok {
		return r, nil
	}
	resource, err := getResource(ctx, resourceID, c.gwClient)
	if err == nil {
		c.resources[resourceID.GetOpaqueId()] = resource
	}
	return resource, err
}

func (c *Converter) getUser(ctx context.Context, userID *user.UserId) (*user.User, error) {
	if u, ok := c.users[userID.GetOpaqueId()]; ok {
		return u, nil
	}
	usr, err := getUser(ctx, userID, c.gwClient)
	if err == nil {
		c.users[userID.GetOpaqueId()] = usr
	}
	return usr, err
}

func composeMessage(nt NotificationTemplate, locale string, vars map[string]interface{}) (string, string, string, string, error) {
	subj, msg, err := parseTemplate(nt, locale)
	if err != nil {
		return "", "", "", "", err
	}

	subject, err := executeTemplate(subj, vars)
	if err != nil {
		return "", "", "", "", err
	}

	subjectraw, err := executeTemplate(subj, _placeholders)
	if err != nil {
		return "", "", "", "", err
	}

	message, err := executeTemplate(msg, vars)
	if err != nil {
		return "", "", "", "", err
	}

	messageraw, err := executeTemplate(msg, _placeholders)
	return subject, subjectraw, message, messageraw, err

}

func parseTemplate(nt NotificationTemplate, locale string) (*template.Template, *template.Template, error) {
	// Create Locale with library path and language code and load domain '.../default.po'
	l := gotext.NewLocale(_pathToLocales, locale)
	l.AddDomain("default")

	subject, err := template.New("").Parse(l.Get(nt.Subject))
	if err != nil {
		return nil, nil, err
	}

	message, err := template.New("").Parse(l.Get(nt.Message))
	return subject, message, err

}

func executeTemplate(tpl *template.Template, vars map[string]interface{}) (string, error) {
	var writer bytes.Buffer
	if err := tpl.Execute(&writer, vars); err != nil {
		return "", err
	}

	return writer.String(), nil
}

func generateDetails(user *user.User, space *storageprovider.StorageSpace, item *storageprovider.ResourceInfo, shareid *collaboration.ShareId) map[string]interface{} {
	details := make(map[string]interface{})

	if user != nil {
		details["user"] = map[string]string{
			"id":          user.GetId().GetOpaqueId(),
			"name":        user.GetUsername(),
			"displayname": user.GetDisplayName(),
		}
	}

	if space != nil {
		details["space"] = map[string]string{
			"id":   space.GetId().GetOpaqueId(),
			"name": space.GetName(),
		}
	}

	if item != nil {
		details["resource"] = map[string]string{
			"id":   storagespace.FormatResourceID(*item.GetId()),
			"name": item.GetName(),
		}
	}

	if shareid != nil {
		details["share"] = map[string]string{
			"id": shareid.GetOpaqueId(),
		}
	}

	return details
}
