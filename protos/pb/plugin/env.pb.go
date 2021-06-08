// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/plugin/env.proto

package plugin

import (
	languageSpecific "code-analyser/protos/pb/output/languageSpecific"
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

type EnvsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *languageSpecific.Envs `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Error *Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *EnvsOutput) Reset() {
	*x = EnvsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_env_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvsOutput) ProtoMessage() {}

func (x *EnvsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_env_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvsOutput.ProtoReflect.Descriptor instead.
func (*EnvsOutput) Descriptor() ([]byte, []int) {
	return file_protos_plugin_env_proto_rawDescGZIP(), []int{0}
}

func (x *EnvsOutput) GetValue() *languageSpecific.Envs {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *EnvsOutput) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_protos_plugin_env_proto protoreflect.FileDescriptor

var file_protos_plugin_env_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x65, 0x6e, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2f, 0x65, 0x6e,
	0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0a, 0x45, 0x6e, 0x76, 0x73, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e, 0x45, 0x6e, 0x76, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x30, 0x0a, 0x03, 0x45, 0x6e, 0x76, 0x12, 0x29, 0x0a,
	0x06, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e,
	0x76, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x20, 0x5a, 0x1e, 0x63, 0x6f, 0x64, 0x65,
	0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_plugin_env_proto_rawDescOnce sync.Once
	file_protos_plugin_env_proto_rawDescData = file_protos_plugin_env_proto_rawDesc
)

func file_protos_plugin_env_proto_rawDescGZIP() []byte {
	file_protos_plugin_env_proto_rawDescOnce.Do(func() {
		file_protos_plugin_env_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_plugin_env_proto_rawDescData)
	})
	return file_protos_plugin_env_proto_rawDescData
}

var file_protos_plugin_env_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_plugin_env_proto_goTypes = []interface{}{
	(*EnvsOutput)(nil),            // 0: proto.EnvsOutput
	(*languageSpecific.Envs)(nil), // 1: protos.output.language_specific.Envs
	(*Error)(nil),                 // 2: proto.Error
	(*Input)(nil),                 // 3: proto.Input
}
var file_protos_plugin_env_proto_depIdxs = []int32{
	1, // 0: proto.EnvsOutput.value:type_name -> protos.output.language_specific.Envs
	2, // 1: proto.EnvsOutput.error:type_name -> proto.Error
	3, // 2: proto.Env.Detect:input_type -> proto.Input
	0, // 3: proto.Env.Detect:output_type -> proto.EnvsOutput
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_plugin_env_proto_init() }
func file_protos_plugin_env_proto_init() {
	if File_protos_plugin_env_proto != nil {
		return
	}
	file_protos_plugin_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_plugin_env_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvsOutput); i {
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
			RawDescriptor: file_protos_plugin_env_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_plugin_env_proto_goTypes,
		DependencyIndexes: file_protos_plugin_env_proto_depIdxs,
		MessageInfos:      file_protos_plugin_env_proto_msgTypes,
	}.Build()
	File_protos_plugin_env_proto = out.File
	file_protos_plugin_env_proto_rawDesc = nil
	file_protos_plugin_env_proto_goTypes = nil
	file_protos_plugin_env_proto_depIdxs = nil
}
