// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/output/globalFiles/dockercompose.proto

package globalFiles

import (
	helpers "code-analyser/protos/pb/helpers"
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

type DockerCompose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    *helpers.Error     `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Used     bool               `protobuf:"varint,2,opt,name=used,proto3" json:"used,omitempty"`
	FilePath string             `protobuf:"bytes,3,opt,name=filePath,proto3" json:"filePath,omitempty"`
	Commands []*helpers.Command `protobuf:"bytes,4,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *DockerCompose) Reset() {
	*x = DockerCompose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_output_globalFiles_dockercompose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerCompose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerCompose) ProtoMessage() {}

func (x *DockerCompose) ProtoReflect() protoreflect.Message {
	mi := &file_protos_output_globalFiles_dockercompose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerCompose.ProtoReflect.Descriptor instead.
func (*DockerCompose) Descriptor() ([]byte, []int) {
	return file_protos_output_globalFiles_dockercompose_proto_rawDescGZIP(), []int{0}
}

func (x *DockerCompose) GetError() *helpers.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *DockerCompose) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *DockerCompose) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *DockerCompose) GetCommands() []*helpers.Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

var File_protos_output_globalFiles_dockercompose_proto protoreflect.FileDescriptor

var file_protos_output_globalFiles_dockercompose_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x0d, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x63, 0x6f, 0x64,
	0x65, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_output_globalFiles_dockercompose_proto_rawDescOnce sync.Once
	file_protos_output_globalFiles_dockercompose_proto_rawDescData = file_protos_output_globalFiles_dockercompose_proto_rawDesc
)

func file_protos_output_globalFiles_dockercompose_proto_rawDescGZIP() []byte {
	file_protos_output_globalFiles_dockercompose_proto_rawDescOnce.Do(func() {
		file_protos_output_globalFiles_dockercompose_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_output_globalFiles_dockercompose_proto_rawDescData)
	})
	return file_protos_output_globalFiles_dockercompose_proto_rawDescData
}

var file_protos_output_globalFiles_dockercompose_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_output_globalFiles_dockercompose_proto_goTypes = []interface{}{
	(*DockerCompose)(nil),   // 0: protos.output.globalFiles.DockerCompose
	(*helpers.Error)(nil),   // 1: protos.helpers.Error
	(*helpers.Command)(nil), // 2: protos.helpers.Command
}
var file_protos_output_globalFiles_dockercompose_proto_depIdxs = []int32{
	1, // 0: protos.output.globalFiles.DockerCompose.error:type_name -> protos.helpers.Error
	2, // 1: protos.output.globalFiles.DockerCompose.commands:type_name -> protos.helpers.Command
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_output_globalFiles_dockercompose_proto_init() }
func file_protos_output_globalFiles_dockercompose_proto_init() {
	if File_protos_output_globalFiles_dockercompose_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_output_globalFiles_dockercompose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerCompose); i {
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
			RawDescriptor: file_protos_output_globalFiles_dockercompose_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_output_globalFiles_dockercompose_proto_goTypes,
		DependencyIndexes: file_protos_output_globalFiles_dockercompose_proto_depIdxs,
		MessageInfos:      file_protos_output_globalFiles_dockercompose_proto_msgTypes,
	}.Build()
	File_protos_output_globalFiles_dockercompose_proto = out.File
	file_protos_output_globalFiles_dockercompose_proto_rawDesc = nil
	file_protos_output_globalFiles_dockercompose_proto_goTypes = nil
	file_protos_output_globalFiles_dockercompose_proto_depIdxs = nil
}