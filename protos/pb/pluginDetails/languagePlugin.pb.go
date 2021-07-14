// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/pluginDetails/languagePlugin.proto

package pluginDetails

import (
	utils "code-analyser/protos/pb/output/utils"
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

type LanguagePlugins struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DetectRuntimePluginPath    string                               `protobuf:"bytes,7,opt,name=detectRuntimePluginPath,proto3" json:"detectRuntimePluginPath,omitempty"`
	RuntimeVersions            map[string]*DependencyVersionDetails `protobuf:"bytes,1,rep,name=runtimeVersions,proto3" json:"runtimeVersions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Frameworks                 map[string]*DependencyDetails        `protobuf:"bytes,2,rep,name=frameworks,proto3" json:"frameworks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Databases                  map[string]*DependencyDetails        `protobuf:"bytes,3,rep,name=databases,proto3" json:"databases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Orms                       map[string]*DependencyDetails        `protobuf:"bytes,4,rep,name=orms,proto3" json:"orms,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Libraries                  map[string]*DependencyDetails        `protobuf:"bytes,5,rep,name=libraries,proto3" json:"libraries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Dependencies               map[string]*DependencyDetails        `protobuf:"bytes,6,rep,name=dependencies,proto3" json:"dependencies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EnvPluginPath              string                               `protobuf:"bytes,8,opt,name=EnvPluginPath,proto3" json:"EnvPluginPath,omitempty"`
	PreDetectCommandPluginPath string                               `protobuf:"bytes,9,opt,name=preDetectCommandPluginPath,proto3" json:"preDetectCommandPluginPath,omitempty"`
	StaticAssetsPluginPath     string                               `protobuf:"bytes,10,opt,name=staticAssetsPluginPath,proto3" json:"staticAssetsPluginPath,omitempty"`
	BuildDirectoryPluginPath   string                               `protobuf:"bytes,11,opt,name=buildDirectoryPluginPath,proto3" json:"buildDirectoryPluginPath,omitempty"`
	TestCasesCommandPluginPath string                               `protobuf:"bytes,12,opt,name=TestCasesCommandPluginPath,proto3" json:"TestCasesCommandPluginPath,omitempty"`
	CommandsPluginPath         string                               `protobuf:"bytes,13,opt,name=commandsPluginPath,proto3" json:"commandsPluginPath,omitempty"`
}

func (x *LanguagePlugins) Reset() {
	*x = LanguagePlugins{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pluginDetails_languagePlugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguagePlugins) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguagePlugins) ProtoMessage() {}

func (x *LanguagePlugins) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pluginDetails_languagePlugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguagePlugins.ProtoReflect.Descriptor instead.
func (*LanguagePlugins) Descriptor() ([]byte, []int) {
	return file_protos_pluginDetails_languagePlugin_proto_rawDescGZIP(), []int{0}
}

func (x *LanguagePlugins) GetDetectRuntimePluginPath() string {
	if x != nil {
		return x.DetectRuntimePluginPath
	}
	return ""
}

func (x *LanguagePlugins) GetRuntimeVersions() map[string]*DependencyVersionDetails {
	if x != nil {
		return x.RuntimeVersions
	}
	return nil
}

func (x *LanguagePlugins) GetFrameworks() map[string]*DependencyDetails {
	if x != nil {
		return x.Frameworks
	}
	return nil
}

func (x *LanguagePlugins) GetDatabases() map[string]*DependencyDetails {
	if x != nil {
		return x.Databases
	}
	return nil
}

func (x *LanguagePlugins) GetOrms() map[string]*DependencyDetails {
	if x != nil {
		return x.Orms
	}
	return nil
}

func (x *LanguagePlugins) GetLibraries() map[string]*DependencyDetails {
	if x != nil {
		return x.Libraries
	}
	return nil
}

func (x *LanguagePlugins) GetDependencies() map[string]*DependencyDetails {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *LanguagePlugins) GetEnvPluginPath() string {
	if x != nil {
		return x.EnvPluginPath
	}
	return ""
}

func (x *LanguagePlugins) GetPreDetectCommandPluginPath() string {
	if x != nil {
		return x.PreDetectCommandPluginPath
	}
	return ""
}

func (x *LanguagePlugins) GetStaticAssetsPluginPath() string {
	if x != nil {
		return x.StaticAssetsPluginPath
	}
	return ""
}

func (x *LanguagePlugins) GetBuildDirectoryPluginPath() string {
	if x != nil {
		return x.BuildDirectoryPluginPath
	}
	return ""
}

func (x *LanguagePlugins) GetTestCasesCommandPluginPath() string {
	if x != nil {
		return x.TestCasesCommandPluginPath
	}
	return ""
}

func (x *LanguagePlugins) GetCommandsPluginPath() string {
	if x != nil {
		return x.CommandsPluginPath
	}
	return ""
}

type DependencyDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version map[string]*DependencyVersionDetails `protobuf:"bytes,1,rep,name=version,proto3" json:"version,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DependencyDetails) Reset() {
	*x = DependencyDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pluginDetails_languagePlugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependencyDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyDetails) ProtoMessage() {}

func (x *DependencyDetails) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pluginDetails_languagePlugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyDetails.ProtoReflect.Descriptor instead.
func (*DependencyDetails) Descriptor() ([]byte, []int) {
	return file_protos_pluginDetails_languagePlugin_proto_rawDescGZIP(), []int{1}
}

func (x *DependencyDetails) GetVersion() map[string]*DependencyVersionDetails {
	if x != nil {
		return x.Version
	}
	return nil
}

type DependencyVersionDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Semver     []string         `protobuf:"bytes,2,rep,name=semver,proto3" json:"semver,omitempty"`
	Libraries  []*utils.Library `protobuf:"bytes,1,rep,name=libraries,proto3" json:"libraries,omitempty"`
	PluginPath string           `protobuf:"bytes,3,opt,name=pluginPath,proto3" json:"pluginPath,omitempty"`
}

func (x *DependencyVersionDetails) Reset() {
	*x = DependencyVersionDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pluginDetails_languagePlugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependencyVersionDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyVersionDetails) ProtoMessage() {}

func (x *DependencyVersionDetails) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pluginDetails_languagePlugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyVersionDetails.ProtoReflect.Descriptor instead.
func (*DependencyVersionDetails) Descriptor() ([]byte, []int) {
	return file_protos_pluginDetails_languagePlugin_proto_rawDescGZIP(), []int{2}
}

func (x *DependencyVersionDetails) GetSemver() []string {
	if x != nil {
		return x.Semver
	}
	return nil
}

func (x *DependencyVersionDetails) GetLibraries() []*utils.Library {
	if x != nil {
		return x.Libraries
	}
	return nil
}

func (x *DependencyVersionDetails) GetPluginPath() string {
	if x != nil {
		return x.PluginPath
	}
	return ""
}

var File_protos_pluginDetails_languagePlugin_proto protoreflect.FileDescriptor

var file_protos_pluginDetails_languagePlugin_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x59, 0x61, 0x6d,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x0c, 0x0a, 0x0f, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x17, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x64, 0x0a, 0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x55, 0x0a, 0x0a, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x12, 0x52, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x04, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x4f, 0x72, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x52, 0x0a, 0x09, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x5b, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x45, 0x6e, 0x76, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x45, 0x6e, 0x76, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x3e, 0x0a, 0x1a, 0x70, 0x72, 0x65, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x70, 0x72, 0x65, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x16, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3a, 0x0a, 0x18, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x1a, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x72, 0x0a, 0x14, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x44, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x66, 0x0a, 0x0f, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x65, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x60, 0x0a, 0x09, 0x4f, 0x72, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x44,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x65, 0x0a, 0x0e, 0x4c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x68, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e,
	0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcf, 0x01, 0x0a,
	0x11, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x4e, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0x6a, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8e,
	0x01, 0x0a, 0x18, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x6d, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6d,
	0x76, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x42,
	0x27, 0x5a, 0x25, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_pluginDetails_languagePlugin_proto_rawDescOnce sync.Once
	file_protos_pluginDetails_languagePlugin_proto_rawDescData = file_protos_pluginDetails_languagePlugin_proto_rawDesc
)

func file_protos_pluginDetails_languagePlugin_proto_rawDescGZIP() []byte {
	file_protos_pluginDetails_languagePlugin_proto_rawDescOnce.Do(func() {
		file_protos_pluginDetails_languagePlugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_pluginDetails_languagePlugin_proto_rawDescData)
	})
	return file_protos_pluginDetails_languagePlugin_proto_rawDescData
}

var file_protos_pluginDetails_languagePlugin_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_pluginDetails_languagePlugin_proto_goTypes = []interface{}{
	(*LanguagePlugins)(nil),          // 0: protos.pluginDetails.LanguagePlugins
	(*DependencyDetails)(nil),        // 1: protos.pluginDetails.DependencyDetails
	(*DependencyVersionDetails)(nil), // 2: protos.pluginDetails.DependencyVersionDetails
	nil,                              // 3: protos.pluginDetails.LanguagePlugins.RuntimeVersionsEntry
	nil,                              // 4: protos.pluginDetails.LanguagePlugins.FrameworksEntry
	nil,                              // 5: protos.pluginDetails.LanguagePlugins.DatabasesEntry
	nil,                              // 6: protos.pluginDetails.LanguagePlugins.OrmsEntry
	nil,                              // 7: protos.pluginDetails.LanguagePlugins.LibrariesEntry
	nil,                              // 8: protos.pluginDetails.LanguagePlugins.DependenciesEntry
	nil,                              // 9: protos.pluginDetails.DependencyDetails.VersionEntry
	(*utils.Library)(nil),            // 10: protos.output.utils.Library
}
var file_protos_pluginDetails_languagePlugin_proto_depIdxs = []int32{
	3,  // 0: protos.pluginDetails.LanguagePlugins.runtimeVersions:type_name -> protos.pluginDetails.LanguagePlugins.RuntimeVersionsEntry
	4,  // 1: protos.pluginDetails.LanguagePlugins.frameworks:type_name -> protos.pluginDetails.LanguagePlugins.FrameworksEntry
	5,  // 2: protos.pluginDetails.LanguagePlugins.databases:type_name -> protos.pluginDetails.LanguagePlugins.DatabasesEntry
	6,  // 3: protos.pluginDetails.LanguagePlugins.orms:type_name -> protos.pluginDetails.LanguagePlugins.OrmsEntry
	7,  // 4: protos.pluginDetails.LanguagePlugins.libraries:type_name -> protos.pluginDetails.LanguagePlugins.LibrariesEntry
	8,  // 5: protos.pluginDetails.LanguagePlugins.dependencies:type_name -> protos.pluginDetails.LanguagePlugins.DependenciesEntry
	9,  // 6: protos.pluginDetails.DependencyDetails.version:type_name -> protos.pluginDetails.DependencyDetails.VersionEntry
	10, // 7: protos.pluginDetails.DependencyVersionDetails.libraries:type_name -> protos.output.utils.Library
	2,  // 8: protos.pluginDetails.LanguagePlugins.RuntimeVersionsEntry.value:type_name -> protos.pluginDetails.DependencyVersionDetails
	1,  // 9: protos.pluginDetails.LanguagePlugins.FrameworksEntry.value:type_name -> protos.pluginDetails.DependencyDetails
	1,  // 10: protos.pluginDetails.LanguagePlugins.DatabasesEntry.value:type_name -> protos.pluginDetails.DependencyDetails
	1,  // 11: protos.pluginDetails.LanguagePlugins.OrmsEntry.value:type_name -> protos.pluginDetails.DependencyDetails
	1,  // 12: protos.pluginDetails.LanguagePlugins.LibrariesEntry.value:type_name -> protos.pluginDetails.DependencyDetails
	1,  // 13: protos.pluginDetails.LanguagePlugins.DependenciesEntry.value:type_name -> protos.pluginDetails.DependencyDetails
	2,  // 14: protos.pluginDetails.DependencyDetails.VersionEntry.value:type_name -> protos.pluginDetails.DependencyVersionDetails
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_protos_pluginDetails_languagePlugin_proto_init() }
func file_protos_pluginDetails_languagePlugin_proto_init() {
	if File_protos_pluginDetails_languagePlugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_pluginDetails_languagePlugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguagePlugins); i {
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
		file_protos_pluginDetails_languagePlugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependencyDetails); i {
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
		file_protos_pluginDetails_languagePlugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependencyVersionDetails); i {
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
			RawDescriptor: file_protos_pluginDetails_languagePlugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_pluginDetails_languagePlugin_proto_goTypes,
		DependencyIndexes: file_protos_pluginDetails_languagePlugin_proto_depIdxs,
		MessageInfos:      file_protos_pluginDetails_languagePlugin_proto_msgTypes,
	}.Build()
	File_protos_pluginDetails_languagePlugin_proto = out.File
	file_protos_pluginDetails_languagePlugin_proto_rawDesc = nil
	file_protos_pluginDetails_languagePlugin_proto_goTypes = nil
	file_protos_pluginDetails_languagePlugin_proto_depIdxs = nil
}
