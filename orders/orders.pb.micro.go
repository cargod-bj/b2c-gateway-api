// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orders/orders.proto

package Orders

import (
	fmt "fmt"
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Orders service

func NewOrdersEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Orders service

type OrdersService interface {
	AutoOrderStateClosed(ctx context.Context, in *common.EmptyDto, opts ...client.CallOption) (*common.Response, error)
	CheckOrderCancel(ctx context.Context, in *common.EmptyDto, opts ...client.CallOption) (*common.Response, error)
	ExpiredCoupon(ctx context.Context, in *common.EmptyDto, opts ...client.CallOption) (*common.Response, error)
}

type ordersService struct {
	c    client.Client
	name string
}

func NewOrdersService(name string, c client.Client) OrdersService {
	return &ordersService{
		c:    c,
		name: name,
	}
}

func (c *ordersService) AutoOrderStateClosed(ctx context.Context, in *common.EmptyDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Orders.AutoOrderStateClosed", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersService) CheckOrderCancel(ctx context.Context, in *common.EmptyDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Orders.CheckOrderCancel", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersService) ExpiredCoupon(ctx context.Context, in *common.EmptyDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Orders.ExpiredCoupon", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Orders service

type OrdersHandler interface {
	AutoOrderStateClosed(context.Context, *common.EmptyDto, *common.Response) error
	CheckOrderCancel(context.Context, *common.EmptyDto, *common.Response) error
	ExpiredCoupon(context.Context, *common.EmptyDto, *common.Response) error
}

func RegisterOrdersHandler(s server.Server, hdlr OrdersHandler, opts ...server.HandlerOption) error {
	type orders interface {
		AutoOrderStateClosed(ctx context.Context, in *common.EmptyDto, out *common.Response) error
		CheckOrderCancel(ctx context.Context, in *common.EmptyDto, out *common.Response) error
		ExpiredCoupon(ctx context.Context, in *common.EmptyDto, out *common.Response) error
	}
	type Orders struct {
		orders
	}
	h := &ordersHandler{hdlr}
	return s.Handle(s.NewHandler(&Orders{h}, opts...))
}

type ordersHandler struct {
	OrdersHandler
}

func (h *ordersHandler) AutoOrderStateClosed(ctx context.Context, in *common.EmptyDto, out *common.Response) error {
	return h.OrdersHandler.AutoOrderStateClosed(ctx, in, out)
}

func (h *ordersHandler) CheckOrderCancel(ctx context.Context, in *common.EmptyDto, out *common.Response) error {
	return h.OrdersHandler.CheckOrderCancel(ctx, in, out)
}

func (h *ordersHandler) ExpiredCoupon(ctx context.Context, in *common.EmptyDto, out *common.Response) error {
	return h.OrdersHandler.ExpiredCoupon(ctx, in, out)
}
