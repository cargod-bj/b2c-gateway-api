// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: scheduleNotice/shcedule_notice.proto

package scheduleNotice

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

// Api Endpoints for ScheduleNotice service

func NewScheduleNoticeEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ScheduleNotice service

type ScheduleNoticeService interface {
	ScheduleNotice(ctx context.Context, in *common.EmptyDto, opts ...client.CallOption) (*common.Response, error)
	ScheduleReport2DingTalkForCarMaintenanceInfo(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error)
}

type scheduleNoticeService struct {
	c    client.Client
	name string
}

func NewScheduleNoticeService(name string, c client.Client) ScheduleNoticeService {
	return &scheduleNoticeService{
		c:    c,
		name: name,
	}
}

func (c *scheduleNoticeService) ScheduleNotice(ctx context.Context, in *common.EmptyDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "ScheduleNotice.ScheduleNotice", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleNoticeService) ScheduleReport2DingTalkForCarMaintenanceInfo(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "ScheduleNotice.ScheduleReport2DingTalkForCarMaintenanceInfo", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ScheduleNotice service

type ScheduleNoticeHandler interface {
	ScheduleNotice(context.Context, *common.EmptyDto, *common.Response) error
	ScheduleReport2DingTalkForCarMaintenanceInfo(context.Context, *common.Page, *common.Response) error
}

func RegisterScheduleNoticeHandler(s server.Server, hdlr ScheduleNoticeHandler, opts ...server.HandlerOption) error {
	type scheduleNotice interface {
		ScheduleNotice(ctx context.Context, in *common.EmptyDto, out *common.Response) error
		ScheduleReport2DingTalkForCarMaintenanceInfo(ctx context.Context, in *common.Page, out *common.Response) error
	}
	type ScheduleNotice struct {
		scheduleNotice
	}
	h := &scheduleNoticeHandler{hdlr}
	return s.Handle(s.NewHandler(&ScheduleNotice{h}, opts...))
}

type scheduleNoticeHandler struct {
	ScheduleNoticeHandler
}

func (h *scheduleNoticeHandler) ScheduleNotice(ctx context.Context, in *common.EmptyDto, out *common.Response) error {
	return h.ScheduleNoticeHandler.ScheduleNotice(ctx, in, out)
}

func (h *scheduleNoticeHandler) ScheduleReport2DingTalkForCarMaintenanceInfo(ctx context.Context, in *common.Page, out *common.Response) error {
	return h.ScheduleNoticeHandler.ScheduleReport2DingTalkForCarMaintenanceInfo(ctx, in, out)
}
