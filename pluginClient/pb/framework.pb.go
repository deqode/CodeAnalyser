// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: proto/framework.proto

package pb

import (
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

var File_proto_framework_proto protoreflect.FileDescriptor

var file_proto_framework_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xd7, 0x01, 0x0a, 0x10, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6c,
	0x12, 0x40, 0x0a, 0x0f, 0x49, 0x73, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x55,
	0x73, 0x65, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x6f,
	0x6f, 0x6c, 0x12, 0x48, 0x0a, 0x16, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x4f, 0x66, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x64, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_framework_proto_goTypes = []interface{}{
	(*ServiceInput)(nil),       // 0: proto.ServiceInput
	(*ServiceOutputBool)(nil),  // 1: proto.ServiceOutputBool
	(*ServiceOutputFloat)(nil), // 2: proto.ServiceOutputFloat
}
var file_proto_framework_proto_depIdxs = []int32{
	0, // 0: proto.FrameworkService.Detect:input_type -> proto.ServiceInput
	0, // 1: proto.FrameworkService.IsFrameworkUsed:input_type -> proto.ServiceInput
	0, // 2: proto.FrameworkService.PercentOfFrameworkUsed:input_type -> proto.ServiceInput
	1, // 3: proto.FrameworkService.Detect:output_type -> proto.ServiceOutputBool
	1, // 4: proto.FrameworkService.IsFrameworkUsed:output_type -> proto.ServiceOutputBool
	2, // 5: proto.FrameworkService.PercentOfFrameworkUsed:output_type -> proto.ServiceOutputFloat
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_framework_proto_init() }
func file_proto_framework_proto_init() {
	if File_proto_framework_proto != nil {
		return
	}
	file_proto_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_framework_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_framework_proto_goTypes,
		DependencyIndexes: file_proto_framework_proto_depIdxs,
	}.Build()
	File_proto_framework_proto = out.File
	file_proto_framework_proto_rawDesc = nil
	file_proto_framework_proto_goTypes = nil
	file_proto_framework_proto_depIdxs = nil
}
