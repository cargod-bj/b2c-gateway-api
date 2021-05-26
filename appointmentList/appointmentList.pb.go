// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: appointmentList/appointmentList.proto

package appointmentList

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SmsCond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendInterval uint32 `protobuf:"varint,1,opt,name=sendInterval,proto3" json:"sendInterval,omitempty"`
}

func (x *SmsCond) Reset() {
	*x = SmsCond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointmentList_appointmentList_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsCond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsCond) ProtoMessage() {}

func (x *SmsCond) ProtoReflect() protoreflect.Message {
	mi := &file_appointmentList_appointmentList_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsCond.ProtoReflect.Descriptor instead.
func (*SmsCond) Descriptor() ([]byte, []int) {
	return file_appointmentList_appointmentList_proto_rawDescGZIP(), []int{0}
}

func (x *SmsCond) GetSendInterval() uint32 {
	if x != nil {
		return x.SendInterval
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string     `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Code string     `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointmentList_appointmentList_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_appointmentList_appointmentList_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_appointmentList_appointmentList_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Response) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_appointmentList_appointmentList_proto protoreflect.FileDescriptor

var file_appointmentList_appointmentList_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x07, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x22, 0x5a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x59,
	0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x46, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x53,
	0x4d, 0x53, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x1a, 0x19, 0x2e, 0x61,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62,
	0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appointmentList_appointmentList_proto_rawDescOnce sync.Once
	file_appointmentList_appointmentList_proto_rawDescData = file_appointmentList_appointmentList_proto_rawDesc
)

func file_appointmentList_appointmentList_proto_rawDescGZIP() []byte {
	file_appointmentList_appointmentList_proto_rawDescOnce.Do(func() {
		file_appointmentList_appointmentList_proto_rawDescData = protoimpl.X.CompressGZIP(file_appointmentList_appointmentList_proto_rawDescData)
	})
	return file_appointmentList_appointmentList_proto_rawDescData
}

var file_appointmentList_appointmentList_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_appointmentList_appointmentList_proto_goTypes = []interface{}{
	(*SmsCond)(nil),   // 0: appointmentList.SmsCond
	(*Response)(nil),  // 1: appointmentList.Response
	(*anypb.Any)(nil), // 2: google.protobuf.Any
}
var file_appointmentList_appointmentList_proto_depIdxs = []int32{
	2, // 0: appointmentList.Response.data:type_name -> google.protobuf.Any
	0, // 1: appointmentList.AppointmentList.GetListForSMS:input_type -> appointmentList.SmsCond
	1, // 2: appointmentList.AppointmentList.GetListForSMS:output_type -> appointmentList.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_appointmentList_appointmentList_proto_init() }
func file_appointmentList_appointmentList_proto_init() {
	if File_appointmentList_appointmentList_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appointmentList_appointmentList_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsCond); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_appointmentList_appointmentList_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_appointmentList_appointmentList_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appointmentList_appointmentList_proto_goTypes,
		DependencyIndexes: file_appointmentList_appointmentList_proto_depIdxs,
		MessageInfos:      file_appointmentList_appointmentList_proto_msgTypes,
	}.Build()
	File_appointmentList_appointmentList_proto = out.File
	file_appointmentList_appointmentList_proto_rawDesc = nil
	file_appointmentList_appointmentList_proto_goTypes = nil
	file_appointmentList_appointmentList_proto_depIdxs = nil
}
