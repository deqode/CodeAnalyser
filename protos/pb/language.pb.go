// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: protos/versions/language.proto

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

type LanguageVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Versions     []string    `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
	Framework    []*Versions `protobuf:"bytes,3,rep,name=framework,proto3" json:"framework,omitempty"`
	Databases    []*Versions `protobuf:"bytes,4,rep,name=databases,proto3" json:"databases,omitempty"`
	Orms         []*Versions `protobuf:"bytes,5,rep,name=orms,proto3" json:"orms,omitempty"`
	Libraries    []*Versions `protobuf:"bytes,6,rep,name=libraries,proto3" json:"libraries,omitempty"`
	Dependencies []*Versions `protobuf:"bytes,7,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
}

func (x *LanguageVersion) Reset() {
	*x = LanguageVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_versions_language_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageVersion) ProtoMessage() {}

func (x *LanguageVersion) ProtoReflect() protoreflect.Message {
	mi := &file_protos_versions_language_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageVersion.ProtoReflect.Descriptor instead.
func (*LanguageVersion) Descriptor() ([]byte, []int) {
	return file_protos_versions_language_proto_rawDescGZIP(), []int{0}
}

func (x *LanguageVersion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LanguageVersion) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *LanguageVersion) GetFramework() []*Versions {
	if x != nil {
		return x.Framework
	}
	return nil
}

func (x *LanguageVersion) GetDatabases() []*Versions {
	if x != nil {
		return x.Databases
	}
	return nil
}

func (x *LanguageVersion) GetOrms() []*Versions {
	if x != nil {
		return x.Orms
	}
	return nil
}

func (x *LanguageVersion) GetLibraries() []*Versions {
	if x != nil {
		return x.Libraries
	}
	return nil
}

func (x *LanguageVersion) GetDependencies() []*Versions {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

type Versions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Versions []*Semver `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *Versions) Reset() {
	*x = Versions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_versions_language_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Versions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Versions) ProtoMessage() {}

func (x *Versions) ProtoReflect() protoreflect.Message {
	mi := &file_protos_versions_language_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Versions.ProtoReflect.Descriptor instead.
func (*Versions) Descriptor() ([]byte, []int) {
	return file_protos_versions_language_proto_rawDescGZIP(), []int{1}
}

func (x *Versions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Versions) GetVersions() []*Semver {
	if x != nil {
		return x.Versions
	}
	return nil
}

type Semver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Semver string `protobuf:"bytes,2,opt,name=semver,proto3" json:"semver,omitempty"`
}

func (x *Semver) Reset() {
	*x = Semver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_versions_language_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Semver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Semver) ProtoMessage() {}

func (x *Semver) ProtoReflect() protoreflect.Message {
	mi := &file_protos_versions_language_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Semver.ProtoReflect.Descriptor instead.
func (*Semver) Descriptor() ([]byte, []int) {
	return file_protos_versions_language_proto_rawDescGZIP(), []int{2}
}

func (x *Semver) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Semver) GetSemver() string {
	if x != nil {
		return x.Semver
	}
	return ""
}

var File_protos_versions_language_proto protoreflect.FileDescriptor

var file_protos_versions_language_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8a, 0x02, 0x0a, 0x0f, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x27, 0x0a,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x04, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x27, 0x0a, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2d,
	0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x43, 0x0a,
	0x08, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x53, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x34, 0x0a, 0x06, 0x53, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_versions_language_proto_rawDescOnce sync.Once
	file_protos_versions_language_proto_rawDescData = file_protos_versions_language_proto_rawDesc
)

func file_protos_versions_language_proto_rawDescGZIP() []byte {
	file_protos_versions_language_proto_rawDescOnce.Do(func() {
		file_protos_versions_language_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_versions_language_proto_rawDescData)
	})
	return file_protos_versions_language_proto_rawDescData
}

var file_protos_versions_language_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_versions_language_proto_goTypes = []interface{}{
	(*LanguageVersion)(nil), // 0: LanguageVersion
	(*Versions)(nil),        // 1: Versions
	(*Semver)(nil),          // 2: Semver
}
var file_protos_versions_language_proto_depIdxs = []int32{
	1, // 0: LanguageVersion.framework:type_name -> Versions
	1, // 1: LanguageVersion.databases:type_name -> Versions
	1, // 2: LanguageVersion.orms:type_name -> Versions
	1, // 3: LanguageVersion.libraries:type_name -> Versions
	1, // 4: LanguageVersion.dependencies:type_name -> Versions
	2, // 5: Versions.versions:type_name -> Semver
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protos_versions_language_proto_init() }
func file_protos_versions_language_proto_init() {
	if File_protos_versions_language_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_versions_language_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageVersion); i {
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
		file_protos_versions_language_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Versions); i {
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
		file_protos_versions_language_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Semver); i {
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
			RawDescriptor: file_protos_versions_language_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_versions_language_proto_goTypes,
		DependencyIndexes: file_protos_versions_language_proto_depIdxs,
		MessageInfos:      file_protos_versions_language_proto_msgTypes,
	}.Build()
	File_protos_versions_language_proto = out.File
	file_protos_versions_language_proto_rawDesc = nil
	file_protos_versions_language_proto_goTypes = nil
	file_protos_versions_language_proto_depIdxs = nil
}
