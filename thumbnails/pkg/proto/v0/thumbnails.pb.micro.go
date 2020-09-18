// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pkg/proto/v0/thumbnails.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ThumbnailService service

type ThumbnailService interface {
	GetThumbnail(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type thumbnailService struct {
	c    client.Client
	name string
}

func NewThumbnailService(name string, c client.Client) ThumbnailService {
	return &thumbnailService{
		c:    c,
		name: name,
	}
}

func (c *thumbnailService) GetThumbnail(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "ThumbnailService.GetThumbnail", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ThumbnailService service

type ThumbnailServiceHandler interface {
	GetThumbnail(context.Context, *GetRequest, *GetResponse) error
}

func RegisterThumbnailServiceHandler(s server.Server, hdlr ThumbnailServiceHandler, opts ...server.HandlerOption) error {
	type thumbnailService interface {
		GetThumbnail(ctx context.Context, in *GetRequest, out *GetResponse) error
	}
	type ThumbnailService struct {
		thumbnailService
	}
	h := &thumbnailServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ThumbnailService{h}, opts...))
}

type thumbnailServiceHandler struct {
	ThumbnailServiceHandler
}

func (h *thumbnailServiceHandler) GetThumbnail(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.ThumbnailServiceHandler.GetThumbnail(ctx, in, out)
}
