// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: protos/plugin/makeFile.proto

package plugin

import (
	global "code-analyser/protos/pb/output/global"
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

type ServiceOutputMakeFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    *ServiceError            `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	MakeFile []*global.MakefileOutput `protobuf:"bytes,2,rep,name=makeFile,proto3" json:"makeFile,omitempty"`
}

func (x *ServiceOutputMakeFile) Reset() {
	*x = ServiceOutputMakeFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_makeFile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOutputMakeFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOutputMakeFile) ProtoMessage() {}

func (x *ServiceOutputMakeFile) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_makeFile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOutputMakeFile.ProtoReflect.Descriptor instead.
func (*ServiceOutputMakeFile) Descriptor() ([]byte, []int) {
	return file_protos_plugin_makeFile_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceOutputMakeFile) GetError() *ServiceError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ServiceOutputMakeFile) GetMakeFile() []*global.MakefileOutput {
	if x != nil {
		return x.MakeFile
	}
	return nil
}

var File_protos_plugin_makeFile_proto protoreflect.FileDescriptor

var file_protos_plugin_makeFile_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x6d, 0x61, 0x6b, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d, 0x61, 0x6b, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x08, 0x6d,
	0x61, 0x6b, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x52, 0x08, 0x6d, 0x61, 0x6b, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x32, 0x54, 0x0a,
	0x0f, 0x4d, 0x61, 0x6b, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d, 0x61, 0x6b, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_plugin_makeFile_proto_rawDescOnce sync.Once
	file_protos_plugin_makeFile_proto_rawDescData = file_protos_plugin_makeFile_proto_rawDesc
)

func file_protos_plugin_makeFile_proto_rawDescGZIP() []byte {
	file_protos_plugin_makeFile_proto_rawDescOnce.Do(func() {
		file_protos_plugin_makeFile_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_plugin_makeFile_proto_rawDescData)
	})
	return file_protos_plugin_makeFile_proto_rawDescData
}

var file_protos_plugin_makeFile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_plugin_makeFile_proto_goTypes = []interface{}{
	(*ServiceOutputMakeFile)(nil), // 0: proto.ServiceOutputMakeFile
	(*ServiceError)(nil),          // 1: proto.ServiceError
	(*global.MakefileOutput)(nil), // 2: protos.output.global.MakefileOutput
	(*ServiceInputString)(nil),    // 3: proto.ServiceInputString
}
var file_protos_plugin_makeFile_proto_depIdxs = []int32{
	1, // 0: proto.ServiceOutputMakeFile.error:type_name -> proto.ServiceError
	2, // 1: proto.ServiceOutputMakeFile.makeFile:type_name -> protos.output.global.MakefileOutput
	3, // 2: proto.MakeFileService.Detect:input_type -> proto.ServiceInputString
	0, // 3: proto.MakeFileService.Detect:output_type -> proto.ServiceOutputMakeFile
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_plugin_makeFile_proto_init() }
func file_protos_plugin_makeFile_proto_init() {
	if File_protos_plugin_makeFile_proto != nil {
		return
	}
	file_protos_plugin_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_plugin_makeFile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOutputMakeFile); i {
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
			RawDescriptor: file_protos_plugin_makeFile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_plugin_makeFile_proto_goTypes,
		DependencyIndexes: file_protos_plugin_makeFile_proto_depIdxs,
		MessageInfos:      file_protos_plugin_makeFile_proto_msgTypes,
	}.Build()
	File_protos_plugin_makeFile_proto = out.File
	file_protos_plugin_makeFile_proto_rawDesc = nil
	file_protos_plugin_makeFile_proto_goTypes = nil
	file_protos_plugin_makeFile_proto_depIdxs = nil
}
