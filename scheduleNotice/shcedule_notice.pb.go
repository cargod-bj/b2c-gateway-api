// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: scheduleNotice/shcedule_notice.proto

package scheduleNotice

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_scheduleNotice_shcedule_notice_proto protoreflect.FileDescriptor

var file_scheduleNotice_shcedule_notice_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x2f, 0x73, 0x68, 0x63, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xdd, 0x01, 0x0a, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x2c,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x44,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c, 0x6b, 0x46, 0x6f, 0x72, 0x43, 0x61, 0x72, 0x4d, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x19, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x6d, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_scheduleNotice_shcedule_notice_proto_goTypes = []interface{}{
	(*common.EmptyDto)(nil), // 0: common.EmptyDto
	(*common.Page)(nil),     // 1: common.Page
	(*common.Response)(nil), // 2: common.Response
}
var file_scheduleNotice_shcedule_notice_proto_depIdxs = []int32{
	0, // 0: scheduleNotice.ScheduleNotice.ScheduleNotice:input_type -> common.EmptyDto
	1, // 1: scheduleNotice.ScheduleNotice.ScheduleReport2DingTalkForCarMaintenanceInfo:input_type -> common.Page
	0, // 2: scheduleNotice.ScheduleNotice.ScheduleOrderCancelledSms:input_type -> common.EmptyDto
	2, // 3: scheduleNotice.ScheduleNotice.ScheduleNotice:output_type -> common.Response
	2, // 4: scheduleNotice.ScheduleNotice.ScheduleReport2DingTalkForCarMaintenanceInfo:output_type -> common.Response
	2, // 5: scheduleNotice.ScheduleNotice.ScheduleOrderCancelledSms:output_type -> common.Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scheduleNotice_shcedule_notice_proto_init() }
func file_scheduleNotice_shcedule_notice_proto_init() {
	if File_scheduleNotice_shcedule_notice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scheduleNotice_shcedule_notice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scheduleNotice_shcedule_notice_proto_goTypes,
		DependencyIndexes: file_scheduleNotice_shcedule_notice_proto_depIdxs,
	}.Build()
	File_scheduleNotice_shcedule_notice_proto = out.File
	file_scheduleNotice_shcedule_notice_proto_rawDesc = nil
	file_scheduleNotice_shcedule_notice_proto_goTypes = nil
	file_scheduleNotice_shcedule_notice_proto_depIdxs = nil
}
