// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: systemService/system_service.proto

package systemService

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageAreaDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Location    string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	DeviceType  string `protobuf:"bytes,3,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	ShowStatus  uint64 `protobuf:"varint,4,opt,name=showStatus,proto3" json:"showStatus,omitempty"`
	CopyWriting string `protobuf:"bytes,5,opt,name=copyWriting,proto3" json:"copyWriting,omitempty"`
	ButtonText  string `protobuf:"bytes,6,opt,name=buttonText,proto3" json:"buttonText,omitempty"`
	JumpUrl     string `protobuf:"bytes,7,opt,name=jumpUrl,proto3" json:"jumpUrl,omitempty"`
}

func (x *MessageAreaDTO) Reset() {
	*x = MessageAreaDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemService_system_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAreaDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAreaDTO) ProtoMessage() {}

func (x *MessageAreaDTO) ProtoReflect() protoreflect.Message {
	mi := &file_systemService_system_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAreaDTO.ProtoReflect.Descriptor instead.
func (*MessageAreaDTO) Descriptor() ([]byte, []int) {
	return file_systemService_system_service_proto_rawDescGZIP(), []int{0}
}

func (x *MessageAreaDTO) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageAreaDTO) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *MessageAreaDTO) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *MessageAreaDTO) GetShowStatus() uint64 {
	if x != nil {
		return x.ShowStatus
	}
	return 0
}

func (x *MessageAreaDTO) GetCopyWriting() string {
	if x != nil {
		return x.CopyWriting
	}
	return ""
}

func (x *MessageAreaDTO) GetButtonText() string {
	if x != nil {
		return x.ButtonText
	}
	return ""
}

func (x *MessageAreaDTO) GetJumpUrl() string {
	if x != nil {
		return x.JumpUrl
	}
	return ""
}

type MessageAreaLIST struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*MessageAreaDTO `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *MessageAreaLIST) Reset() {
	*x = MessageAreaLIST{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemService_system_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAreaLIST) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAreaLIST) ProtoMessage() {}

func (x *MessageAreaLIST) ProtoReflect() protoreflect.Message {
	mi := &file_systemService_system_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAreaLIST.ProtoReflect.Descriptor instead.
func (*MessageAreaLIST) Descriptor() ([]byte, []int) {
	return file_systemService_system_service_proto_rawDescGZIP(), []int{1}
}

func (x *MessageAreaLIST) GetList() []*MessageAreaDTO {
	if x != nil {
		return x.List
	}
	return nil
}

var File_systemService_system_service_proto protoreflect.FileDescriptor

var file_systemService_system_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8,
	0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x72, 0x65, 0x61, 0x44, 0x54,
	0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x70, 0x79, 0x57, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x70, 0x79, 0x57, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6a, 0x75, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6a, 0x75, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x22, 0x44, 0x0a, 0x0f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x49, 0x53, 0x54, 0x12, 0x31, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x41, 0x72, 0x65, 0x61, 0x44, 0x54, 0x4f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32,
	0x54, 0x0a, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x72,
	0x65, 0x61, 0x12, 0x1d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x72, 0x65, 0x61, 0x44, 0x54,
	0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_systemService_system_service_proto_rawDescOnce sync.Once
	file_systemService_system_service_proto_rawDescData = file_systemService_system_service_proto_rawDesc
)

func file_systemService_system_service_proto_rawDescGZIP() []byte {
	file_systemService_system_service_proto_rawDescOnce.Do(func() {
		file_systemService_system_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_systemService_system_service_proto_rawDescData)
	})
	return file_systemService_system_service_proto_rawDescData
}

var file_systemService_system_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_systemService_system_service_proto_goTypes = []interface{}{
	(*MessageAreaDTO)(nil),  // 0: systemService.MessageAreaDTO
	(*MessageAreaLIST)(nil), // 1: systemService.MessageAreaLIST
	(*common.Response)(nil), // 2: common.Response
}
var file_systemService_system_service_proto_depIdxs = []int32{
	0, // 0: systemService.MessageAreaLIST.list:type_name -> systemService.MessageAreaDTO
	0, // 1: systemService.systemService.GetMessageArea:input_type -> systemService.MessageAreaDTO
	2, // 2: systemService.systemService.GetMessageArea:output_type -> common.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_systemService_system_service_proto_init() }
func file_systemService_system_service_proto_init() {
	if File_systemService_system_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_systemService_system_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAreaDTO); i {
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
		file_systemService_system_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAreaLIST); i {
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
			RawDescriptor: file_systemService_system_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_systemService_system_service_proto_goTypes,
		DependencyIndexes: file_systemService_system_service_proto_depIdxs,
		MessageInfos:      file_systemService_system_service_proto_msgTypes,
	}.Build()
	File_systemService_system_service_proto = out.File
	file_systemService_system_service_proto_rawDesc = nil
	file_systemService_system_service_proto_goTypes = nil
	file_systemService_system_service_proto_depIdxs = nil
}
