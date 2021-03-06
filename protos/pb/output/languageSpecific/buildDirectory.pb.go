// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/output/languageSpecific/buildDirectory.proto

package languageSpecific

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

type BuildDirectoryOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// check BuildDirectory
	Used bool `protobuf:"varint,1,opt,name=used,proto3" json:"used,omitempty"`
	// list of all BuildDirectory
	BuildDirectory []*BuildDirectory `protobuf:"bytes,2,rep,name=buildDirectory,proto3" json:"buildDirectory,omitempty"`
	Error          *helpers.Error    `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BuildDirectoryOutput) Reset() {
	*x = BuildDirectoryOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_output_languageSpecific_buildDirectory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildDirectoryOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildDirectoryOutput) ProtoMessage() {}

func (x *BuildDirectoryOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_output_languageSpecific_buildDirectory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildDirectoryOutput.ProtoReflect.Descriptor instead.
func (*BuildDirectoryOutput) Descriptor() ([]byte, []int) {
	return file_protos_output_languageSpecific_buildDirectory_proto_rawDescGZIP(), []int{0}
}

func (x *BuildDirectoryOutput) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *BuildDirectoryOutput) GetBuildDirectory() []*BuildDirectory {
	if x != nil {
		return x.BuildDirectory
	}
	return nil
}

func (x *BuildDirectoryOutput) GetError() *helpers.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type BuildDirectory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  name of Dir
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//  version of root
	Root string `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
}

func (x *BuildDirectory) Reset() {
	*x = BuildDirectory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_output_languageSpecific_buildDirectory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildDirectory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildDirectory) ProtoMessage() {}

func (x *BuildDirectory) ProtoReflect() protoreflect.Message {
	mi := &file_protos_output_languageSpecific_buildDirectory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildDirectory.ProtoReflect.Descriptor instead.
func (*BuildDirectory) Descriptor() ([]byte, []int) {
	return file_protos_output_languageSpecific_buildDirectory_proto_rawDescGZIP(), []int{1}
}

func (x *BuildDirectory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BuildDirectory) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

var File_protos_output_languageSpecific_buildDirectory_proto protoreflect.FileDescriptor

var file_protos_output_languageSpecific_buildDirectory_proto_rawDesc = []byte{
	0x0a, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x65,
	0x6c, 0x70, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x14, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12,
	0x56, 0x0a, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x38, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x42, 0x31,
	0x5a, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_output_languageSpecific_buildDirectory_proto_rawDescOnce sync.Once
	file_protos_output_languageSpecific_buildDirectory_proto_rawDescData = file_protos_output_languageSpecific_buildDirectory_proto_rawDesc
)

func file_protos_output_languageSpecific_buildDirectory_proto_rawDescGZIP() []byte {
	file_protos_output_languageSpecific_buildDirectory_proto_rawDescOnce.Do(func() {
		file_protos_output_languageSpecific_buildDirectory_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_output_languageSpecific_buildDirectory_proto_rawDescData)
	})
	return file_protos_output_languageSpecific_buildDirectory_proto_rawDescData
}

var file_protos_output_languageSpecific_buildDirectory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_output_languageSpecific_buildDirectory_proto_goTypes = []interface{}{
	(*BuildDirectoryOutput)(nil), // 0: protos.output.languageSpecific.BuildDirectoryOutput
	(*BuildDirectory)(nil),       // 1: protos.output.languageSpecific.BuildDirectory
	(*helpers.Error)(nil),        // 2: protos.helpers.Error
}
var file_protos_output_languageSpecific_buildDirectory_proto_depIdxs = []int32{
	1, // 0: protos.output.languageSpecific.BuildDirectoryOutput.buildDirectory:type_name -> protos.output.languageSpecific.BuildDirectory
	2, // 1: protos.output.languageSpecific.BuildDirectoryOutput.error:type_name -> protos.helpers.Error
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_output_languageSpecific_buildDirectory_proto_init() }
func file_protos_output_languageSpecific_buildDirectory_proto_init() {
	if File_protos_output_languageSpecific_buildDirectory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_output_languageSpecific_buildDirectory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildDirectoryOutput); i {
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
		file_protos_output_languageSpecific_buildDirectory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildDirectory); i {
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
			RawDescriptor: file_protos_output_languageSpecific_buildDirectory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_output_languageSpecific_buildDirectory_proto_goTypes,
		DependencyIndexes: file_protos_output_languageSpecific_buildDirectory_proto_depIdxs,
		MessageInfos:      file_protos_output_languageSpecific_buildDirectory_proto_msgTypes,
	}.Build()
	File_protos_output_languageSpecific_buildDirectory_proto = out.File
	file_protos_output_languageSpecific_buildDirectory_proto_rawDesc = nil
	file_protos_output_languageSpecific_buildDirectory_proto_goTypes = nil
	file_protos_output_languageSpecific_buildDirectory_proto_depIdxs = nil
}
