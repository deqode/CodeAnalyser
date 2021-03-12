// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: protos/outputs/languageSpecific/libraries.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LibrariesOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// check library
	Used bool `protobuf:"varint,1,opt,name=used,proto3" json:"used,omitempty"`
	// list of all library
	Libraries []*Library `protobuf:"bytes,2,rep,name=libraries,proto3" json:"libraries,omitempty"`
}

func (x *LibrariesOutput) Reset() {
	*x = LibrariesOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_outputs_languageSpecific_libraries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LibrariesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LibrariesOutput) ProtoMessage() {}

func (x *LibrariesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_outputs_languageSpecific_libraries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LibrariesOutput.ProtoReflect.Descriptor instead.
func (*LibrariesOutput) Descriptor() ([]byte, []int) {
	return file_protos_outputs_languageSpecific_libraries_proto_rawDescGZIP(), []int{0}
}

func (x *LibrariesOutput) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *LibrariesOutput) GetLibraries() []*Library {
	if x != nil {
		return x.Libraries
	}
	return nil
}

type Library struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  name of library
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//  type of library
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Library) Reset() {
	*x = Library{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_outputs_languageSpecific_libraries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Library) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Library) ProtoMessage() {}

func (x *Library) ProtoReflect() protoreflect.Message {
	mi := &file_protos_outputs_languageSpecific_libraries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Library.ProtoReflect.Descriptor instead.
func (*Library) Descriptor() ([]byte, []int) {
	return file_protos_outputs_languageSpecific_libraries_proto_rawDescGZIP(), []int{1}
}

func (x *Library) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Library) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_protos_outputs_languageSpecific_libraries_proto protoreflect.FileDescriptor

var file_protos_outputs_languageSpecific_libraries_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4d, 0x0a, 0x0f, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x09, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x31, 0x0a, 0x07, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_outputs_languageSpecific_libraries_proto_rawDescOnce sync.Once
	file_protos_outputs_languageSpecific_libraries_proto_rawDescData = file_protos_outputs_languageSpecific_libraries_proto_rawDesc
)

func file_protos_outputs_languageSpecific_libraries_proto_rawDescGZIP() []byte {
	file_protos_outputs_languageSpecific_libraries_proto_rawDescOnce.Do(func() {
		file_protos_outputs_languageSpecific_libraries_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_outputs_languageSpecific_libraries_proto_rawDescData)
	})
	return file_protos_outputs_languageSpecific_libraries_proto_rawDescData
}

var file_protos_outputs_languageSpecific_libraries_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_outputs_languageSpecific_libraries_proto_goTypes = []interface{}{
	(*LibrariesOutput)(nil), // 0: LibrariesOutput
	(*Library)(nil),         // 1: Library
}
var file_protos_outputs_languageSpecific_libraries_proto_depIdxs = []int32{
	1, // 0: LibrariesOutput.libraries:type_name -> Library
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_outputs_languageSpecific_libraries_proto_init() }
func file_protos_outputs_languageSpecific_libraries_proto_init() {
	if File_protos_outputs_languageSpecific_libraries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_outputs_languageSpecific_libraries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LibrariesOutput); i {
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
		file_protos_outputs_languageSpecific_libraries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Library); i {
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
			RawDescriptor: file_protos_outputs_languageSpecific_libraries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_outputs_languageSpecific_libraries_proto_goTypes,
		DependencyIndexes: file_protos_outputs_languageSpecific_libraries_proto_depIdxs,
		MessageInfos:      file_protos_outputs_languageSpecific_libraries_proto_msgTypes,
	}.Build()
	File_protos_outputs_languageSpecific_libraries_proto = out.File
	file_protos_outputs_languageSpecific_libraries_proto_rawDesc = nil
	file_protos_outputs_languageSpecific_libraries_proto_goTypes = nil
	file_protos_outputs_languageSpecific_libraries_proto_depIdxs = nil
}
