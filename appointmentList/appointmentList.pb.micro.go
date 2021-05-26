// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: appointmentList/appointmentList.proto

package appointmentList

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/anypb"
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

// Api Endpoints for AppointmentList service

func NewAppointmentListEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AppointmentList service

type AppointmentListService interface {
	GetListForSMS(ctx context.Context, in *SmsCond, opts ...client.CallOption) (*Response, error)
}

type appointmentListService struct {
	c    client.Client
	name string
}

func NewAppointmentListService(name string, c client.Client) AppointmentListService {
	return &appointmentListService{
		c:    c,
		name: name,
	}
}

func (c *appointmentListService) GetListForSMS(ctx context.Context, in *SmsCond, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppointmentList.GetListForSMS", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppointmentList service

type AppointmentListHandler interface {
	GetListForSMS(context.Context, *SmsCond, *Response) error
}

func RegisterAppointmentListHandler(s server.Server, hdlr AppointmentListHandler, opts ...server.HandlerOption) error {
	type appointmentList interface {
		GetListForSMS(ctx context.Context, in *SmsCond, out *Response) error
	}
	type AppointmentList struct {
		appointmentList
	}
	h := &appointmentListHandler{hdlr}
	return s.Handle(s.NewHandler(&AppointmentList{h}, opts...))
}

type appointmentListHandler struct {
	AppointmentListHandler
}

func (h *appointmentListHandler) GetListForSMS(ctx context.Context, in *SmsCond, out *Response) error {
	return h.AppointmentListHandler.GetListForSMS(ctx, in, out)
}
