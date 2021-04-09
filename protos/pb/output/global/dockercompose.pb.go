// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/output/global/dockercompose.proto

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

type DockerComposeFileOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Used               bool                 `protobuf:"varint,1,opt,name=used,proto3" json:"used,omitempty"`
	DockerComposeFiles []*DockerComposeFile `protobuf:"bytes,2,rep,name=dockerComposeFiles,proto3" json:"dockerComposeFiles,omitempty"`
}

func (x *DockerComposeFileOutput) Reset() {
	*x = DockerComposeFileOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_output_global_dockercompose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerComposeFileOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerComposeFileOutput) ProtoMessage() {}

func (x *DockerComposeFileOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_output_global_dockercompose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerComposeFileOutput.ProtoReflect.Descriptor instead.
func (*DockerComposeFileOutput) Descriptor() ([]byte, []int) {
	return file_protos_output_global_dockercompose_proto_rawDescGZIP(), []int{0}
}

func (x *DockerComposeFileOutput) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *DockerComposeFileOutput) GetDockerComposeFiles() []*DockerComposeFile {
	if x != nil {
		return x.DockerComposeFiles
	}
	return nil
}

type DockerComposeFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName              string     `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	DockerComposeCommands []*Command `protobuf:"bytes,2,rep,name=dockerComposeCommands,proto3" json:"dockerComposeCommands,omitempty"`
}

func (x *DockerComposeFile) Reset() {
	*x = DockerComposeFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_output_global_dockercompose_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerComposeFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerComposeFile) ProtoMessage() {}

func (x *DockerComposeFile) ProtoReflect() protoreflect.Message {
	mi := &file_protos_output_global_dockercompose_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerComposeFile.ProtoReflect.Descriptor instead.
func (*DockerComposeFile) Descriptor() ([]byte, []int) {
	return file_protos_output_global_dockercompose_proto_rawDescGZIP(), []int{1}
}

func (x *DockerComposeFile) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DockerComposeFile) GetDockerComposeCommands() []*Command {
	if x != nil {
		return x.DockerComposeCommands
	}
	return nil
}

var File_protos_output_global_dockercompose_proto protoreflect.FileDescriptor

var file_protos_output_global_dockercompose_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x17, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x64, 0x12, 0x57, 0x0a, 0x12, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x12, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x84, 0x01,
	0x0a, 0x11, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x53, 0x0a, 0x15, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x15, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_output_global_dockercompose_proto_rawDescOnce sync.Once
	file_protos_output_global_dockercompose_proto_rawDescData = file_protos_output_global_dockercompose_proto_rawDesc
)

func file_protos_output_global_dockercompose_proto_rawDescGZIP() []byte {
	file_protos_output_global_dockercompose_proto_rawDescOnce.Do(func() {
		file_protos_output_global_dockercompose_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_output_global_dockercompose_proto_rawDescData)
	})
	return file_protos_output_global_dockercompose_proto_rawDescData
}

var file_protos_output_global_dockercompose_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_output_global_dockercompose_proto_goTypes = []interface{}{
	(*DockerComposeFileOutput)(nil), // 0: protos.output.global.DockerComposeFileOutput
	(*DockerComposeFile)(nil),       // 1: protos.output.global.DockerComposeFile
	(*Command)(nil),                 // 2: protos.output.global.Command
}
var file_protos_output_global_dockercompose_proto_depIdxs = []int32{
	1, // 0: protos.output.global.DockerComposeFileOutput.dockerComposeFiles:type_name -> protos.output.global.DockerComposeFile
	2, // 1: protos.output.global.DockerComposeFile.dockerComposeCommands:type_name -> protos.output.global.Command
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_output_global_dockercompose_proto_init() }
func file_protos_output_global_dockercompose_proto_init() {
	if File_protos_output_global_dockercompose_proto != nil {
		return
	}
	file_protos_output_global_command_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_output_global_dockercompose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerComposeFileOutput); i {
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
		file_protos_output_global_dockercompose_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerComposeFile); i {
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
			RawDescriptor: file_protos_output_global_dockercompose_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_output_global_dockercompose_proto_goTypes,
		DependencyIndexes: file_protos_output_global_dockercompose_proto_depIdxs,
		MessageInfos:      file_protos_output_global_dockercompose_proto_msgTypes,
	}.Build()
	File_protos_output_global_dockercompose_proto = out.File
	file_protos_output_global_dockercompose_proto_rawDesc = nil
	file_protos_output_global_dockercompose_proto_goTypes = nil
	file_protos_output_global_dockercompose_proto_depIdxs = nil
}
