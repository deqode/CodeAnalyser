// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: protos/output/global/buildCommands.proto

package global

import (
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

type BuildCommandsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// check BuildCommands
	Used bool `protobuf:"varint,1,opt,name=used,proto3" json:"used,omitempty"`
	// list of all BuildCommands
	BuildCommands []*Command `protobuf:"bytes,2,rep,name=buildCommands,proto3" json:"buildCommands,omitempty"`
}

func (x *BuildCommandsOutput) Reset() {
	*x = BuildCommandsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_output_global_buildCommands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildCommandsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildCommandsOutput) ProtoMessage() {}

func (x *BuildCommandsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_output_global_buildCommands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildCommandsOutput.ProtoReflect.Descriptor instead.
func (*BuildCommandsOutput) Descriptor() ([]byte, []int) {
	return file_protos_output_global_buildCommands_proto_rawDescGZIP(), []int{0}
}

func (x *BuildCommandsOutput) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *BuildCommandsOutput) GetBuildCommands() []*Command {
	if x != nil {
		return x.BuildCommands
	}
	return nil
}

var File_protos_output_global_buildCommands_proto protoreflect.FileDescriptor

var file_protos_output_global_buildCommands_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x13, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12,
	0x43, 0x0a, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_output_global_buildCommands_proto_rawDescOnce sync.Once
	file_protos_output_global_buildCommands_proto_rawDescData = file_protos_output_global_buildCommands_proto_rawDesc
)

func file_protos_output_global_buildCommands_proto_rawDescGZIP() []byte {
	file_protos_output_global_buildCommands_proto_rawDescOnce.Do(func() {
		file_protos_output_global_buildCommands_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_output_global_buildCommands_proto_rawDescData)
	})
	return file_protos_output_global_buildCommands_proto_rawDescData
}

var file_protos_output_global_buildCommands_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_output_global_buildCommands_proto_goTypes = []interface{}{
	(*BuildCommandsOutput)(nil), // 0: protos.output.global.BuildCommandsOutput
	(*Command)(nil),             // 1: protos.output.global.Command
}
var file_protos_output_global_buildCommands_proto_depIdxs = []int32{
	1, // 0: protos.output.global.BuildCommandsOutput.buildCommands:type_name -> protos.output.global.Command
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_output_global_buildCommands_proto_init() }
func file_protos_output_global_buildCommands_proto_init() {
	if File_protos_output_global_buildCommands_proto != nil {
		return
	}
	file_protos_output_global_command_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_output_global_buildCommands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildCommandsOutput); i {
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
			RawDescriptor: file_protos_output_global_buildCommands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_output_global_buildCommands_proto_goTypes,
		DependencyIndexes: file_protos_output_global_buildCommands_proto_depIdxs,
		MessageInfos:      file_protos_output_global_buildCommands_proto_msgTypes,
	}.Build()
	File_protos_output_global_buildCommands_proto = out.File
	file_protos_output_global_buildCommands_proto_rawDesc = nil
	file_protos_output_global_buildCommands_proto_goTypes = nil
	file_protos_output_global_buildCommands_proto_depIdxs = nil
}
