// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/versions/language.proto

package protos

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

type LanguageVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runtimeversions *RuntimeVersions  `protobuf:"bytes,1,opt,name=runtimeversions,proto3" json:"runtimeversions,omitempty"`
	Framework       []*PluginVersions `protobuf:"bytes,2,rep,name=framework,proto3" json:"framework,omitempty"`
	Databases       []*PluginVersions `protobuf:"bytes,3,rep,name=databases,proto3" json:"databases,omitempty"`
	Orms            []*PluginVersions `protobuf:"bytes,4,rep,name=orms,proto3" json:"orms,omitempty"`
	Libraries       []*PluginVersions `protobuf:"bytes,5,rep,name=libraries,proto3" json:"libraries,omitempty"`
	Dependencies    []*PluginVersions `protobuf:"bytes,6,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
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

func (x *LanguageVersion) GetRuntimeversions() *RuntimeVersions {
	if x != nil {
		return x.Runtimeversions
	}
	return nil
}

func (x *LanguageVersion) GetFramework() []*PluginVersions {
	if x != nil {
		return x.Framework
	}
	return nil
}

func (x *LanguageVersion) GetDatabases() []*PluginVersions {
	if x != nil {
		return x.Databases
	}
	return nil
}

func (x *LanguageVersion) GetOrms() []*PluginVersions {
	if x != nil {
		return x.Orms
	}
	return nil
}

func (x *LanguageVersion) GetLibraries() []*PluginVersions {
	if x != nil {
		return x.Libraries
	}
	return nil
}

func (x *LanguageVersion) GetDependencies() []*PluginVersions {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

type RuntimeVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detector string          `protobuf:"bytes,1,opt,name=detector,proto3" json:"detector,omitempty"`
	Versions []*PluginSemver `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *RuntimeVersions) Reset() {
	*x = RuntimeVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_versions_language_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuntimeVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeVersions) ProtoMessage() {}

func (x *RuntimeVersions) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RuntimeVersions.ProtoReflect.Descriptor instead.
func (*RuntimeVersions) Descriptor() ([]byte, []int) {
	return file_protos_versions_language_proto_rawDescGZIP(), []int{1}
}

func (x *RuntimeVersions) GetDetector() string {
	if x != nil {
		return x.Detector
	}
	return ""
}

func (x *RuntimeVersions) GetVersions() []*PluginSemver {
	if x != nil {
		return x.Versions
	}
	return nil
}

type PluginVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Versions []*PluginSemver `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *PluginVersions) Reset() {
	*x = PluginVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_versions_language_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginVersions) ProtoMessage() {}

func (x *PluginVersions) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PluginVersions.ProtoReflect.Descriptor instead.
func (*PluginVersions) Descriptor() ([]byte, []int) {
	return file_protos_versions_language_proto_rawDescGZIP(), []int{2}
}

func (x *PluginVersions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginVersions) GetVersions() []*PluginSemver {
	if x != nil {
		return x.Versions
	}
	return nil
}

type PluginSemver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Semver        string `protobuf:"bytes,2,opt,name=semver,proto3" json:"semver,omitempty"`
	Plugincommand string `protobuf:"bytes,4,opt,name=plugincommand,proto3" json:"plugincommand,omitempty"`
}

func (x *PluginSemver) Reset() {
	*x = PluginSemver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_versions_language_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginSemver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginSemver) ProtoMessage() {}

func (x *PluginSemver) ProtoReflect() protoreflect.Message {
	mi := &file_protos_versions_language_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginSemver.ProtoReflect.Descriptor instead.
func (*PluginSemver) Descriptor() ([]byte, []int) {
	return file_protos_versions_language_proto_rawDescGZIP(), []int{3}
}

func (x *PluginSemver) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginSemver) GetSemver() string {
	if x != nil {
		return x.Semver
	}
	return ""
}

func (x *PluginSemver) GetPlugincommand() string {
	if x != nil {
		return x.Plugincommand
	}
	return ""
}

var File_protos_versions_language_proto protoreflect.FileDescriptor

var file_protos_versions_language_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb4, 0x02, 0x0a, 0x0f, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x2d, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x2d, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x23,
	0x0a, 0x04, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x6f,
	0x72, 0x6d, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x33, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x58, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x53, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x4f, 0x0a, 0x0e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x53, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x60, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x6d, 0x76,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_protos_versions_language_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_versions_language_proto_goTypes = []interface{}{
	(*LanguageVersion)(nil), // 0: LanguageVersion
	(*RuntimeVersions)(nil), // 1: RuntimeVersions
	(*PluginVersions)(nil),  // 2: PluginVersions
	(*PluginSemver)(nil),    // 3: PluginSemver
}
var file_protos_versions_language_proto_depIdxs = []int32{
	1, // 0: LanguageVersion.runtimeversions:type_name -> RuntimeVersions
	2, // 1: LanguageVersion.framework:type_name -> PluginVersions
	2, // 2: LanguageVersion.databases:type_name -> PluginVersions
	2, // 3: LanguageVersion.orms:type_name -> PluginVersions
	2, // 4: LanguageVersion.libraries:type_name -> PluginVersions
	2, // 5: LanguageVersion.dependencies:type_name -> PluginVersions
	3, // 6: RuntimeVersions.versions:type_name -> PluginSemver
	3, // 7: PluginVersions.versions:type_name -> PluginSemver
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
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
			switch v := v.(*RuntimeVersions); i {
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
			switch v := v.(*PluginVersions); i {
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
		file_protos_versions_language_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginSemver); i {
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
			NumMessages:   4,
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
