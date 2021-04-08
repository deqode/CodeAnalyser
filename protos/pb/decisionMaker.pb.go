// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: protos/decisionMaker/decisionMaker.proto

package pb

import (
	global "code-analyser/protos/pb/output/global"
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

type LanguageSpecificDetections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string                                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RuntimeVersion string                                   `protobuf:"bytes,2,opt,name=runtimeVersion,proto3" json:"runtimeVersion,omitempty"`
	Env            *languageSpecific.EnvOutput              `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	Framework      []*languageSpecific.FrameworkOutput      `protobuf:"bytes,4,rep,name=framework,proto3" json:"framework,omitempty"`
	Db             *languageSpecific.DBOutput               `protobuf:"bytes,5,opt,name=db,proto3" json:"db,omitempty"`
	Orm            *languageSpecific.OrmOutput              `protobuf:"bytes,6,opt,name=orm,proto3" json:"orm,omitempty"`
	Dependencies   []*languageSpecific.DependenciesOutput   `protobuf:"bytes,7,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Libraries      []*languageSpecific.LibraryOutput        `protobuf:"bytes,8,rep,name=libraries,proto3" json:"libraries,omitempty"`
	StaticAssets   *languageSpecific.StaticAssetsOutput     `protobuf:"bytes,9,opt,name=staticAssets,proto3" json:"staticAssets,omitempty"`
	StackOutput    []*languageSpecific.StackOutput          `protobuf:"bytes,10,rep,name=stackOutput,proto3" json:"stackOutput,omitempty"`
	Appserver      []*languageSpecific.AppserverOutput      `protobuf:"bytes,11,rep,name=appserver,proto3" json:"appserver,omitempty"`
	BuildDirectory []*languageSpecific.BuildDirectoryOutput `protobuf:"bytes,12,rep,name=buildDirectory,proto3" json:"buildDirectory,omitempty"`
	Commands       *Commands                                `protobuf:"bytes,13,opt,name=commands,proto3" json:"commands,omitempty"` // todo: DetectTestCasesRunCommands
}

func (x *LanguageSpecificDetections) Reset() {
	*x = LanguageSpecificDetections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_decisionMaker_decisionMaker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageSpecificDetections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageSpecificDetections) ProtoMessage() {}

func (x *LanguageSpecificDetections) ProtoReflect() protoreflect.Message {
	mi := &file_protos_decisionMaker_decisionMaker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageSpecificDetections.ProtoReflect.Descriptor instead.
func (*LanguageSpecificDetections) Descriptor() ([]byte, []int) {
	return file_protos_decisionMaker_decisionMaker_proto_rawDescGZIP(), []int{0}
}

func (x *LanguageSpecificDetections) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LanguageSpecificDetections) GetRuntimeVersion() string {
	if x != nil {
		return x.RuntimeVersion
	}
	return ""
}

func (x *LanguageSpecificDetections) GetEnv() *languageSpecific.EnvOutput {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *LanguageSpecificDetections) GetFramework() []*languageSpecific.FrameworkOutput {
	if x != nil {
		return x.Framework
	}
	return nil
}

func (x *LanguageSpecificDetections) GetDb() *languageSpecific.DBOutput {
	if x != nil {
		return x.Db
	}
	return nil
}

func (x *LanguageSpecificDetections) GetOrm() *languageSpecific.OrmOutput {
	if x != nil {
		return x.Orm
	}
	return nil
}

func (x *LanguageSpecificDetections) GetDependencies() []*languageSpecific.DependenciesOutput {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *LanguageSpecificDetections) GetLibraries() []*languageSpecific.LibraryOutput {
	if x != nil {
		return x.Libraries
	}
	return nil
}

func (x *LanguageSpecificDetections) GetStaticAssets() *languageSpecific.StaticAssetsOutput {
	if x != nil {
		return x.StaticAssets
	}
	return nil
}

func (x *LanguageSpecificDetections) GetStackOutput() []*languageSpecific.StackOutput {
	if x != nil {
		return x.StackOutput
	}
	return nil
}

func (x *LanguageSpecificDetections) GetAppserver() []*languageSpecific.AppserverOutput {
	if x != nil {
		return x.Appserver
	}
	return nil
}

func (x *LanguageSpecificDetections) GetBuildDirectory() []*languageSpecific.BuildDirectoryOutput {
	if x != nil {
		return x.BuildDirectory
	}
	return nil
}

func (x *LanguageSpecificDetections) GetCommands() *Commands {
	if x != nil {
		return x.Commands
	}
	return nil
}

type GlobalDetections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcFile          []*global.ProcFileOutput          `protobuf:"bytes,1,rep,name=procFile,proto3" json:"procFile,omitempty"`
	Makefile          []*global.MakefileOutput          `protobuf:"bytes,2,rep,name=makefile,proto3" json:"makefile,omitempty"`
	DockerFile        []*global.DockerFileOutput        `protobuf:"bytes,3,rep,name=dockerFile,proto3" json:"dockerFile,omitempty"`
	DockerComposeFile []*global.DockerComposeFileOutput `protobuf:"bytes,4,rep,name=dockerComposeFile,proto3" json:"dockerComposeFile,omitempty"`
}

func (x *GlobalDetections) Reset() {
	*x = GlobalDetections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_decisionMaker_decisionMaker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalDetections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalDetections) ProtoMessage() {}

func (x *GlobalDetections) ProtoReflect() protoreflect.Message {
	mi := &file_protos_decisionMaker_decisionMaker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalDetections.ProtoReflect.Descriptor instead.
func (*GlobalDetections) Descriptor() ([]byte, []int) {
	return file_protos_decisionMaker_decisionMaker_proto_rawDescGZIP(), []int{1}
}

func (x *GlobalDetections) GetProcFile() []*global.ProcFileOutput {
	if x != nil {
		return x.ProcFile
	}
	return nil
}

func (x *GlobalDetections) GetMakefile() []*global.MakefileOutput {
	if x != nil {
		return x.Makefile
	}
	return nil
}

func (x *GlobalDetections) GetDockerFile() []*global.DockerFileOutput {
	if x != nil {
		return x.DockerFile
	}
	return nil
}

func (x *GlobalDetections) GetDockerComposeFile() []*global.DockerComposeFileOutput {
	if x != nil {
		return x.DockerComposeFile
	}
	return nil
}

type Commands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildCommands     []*global.BuildCommandsOutput     `protobuf:"bytes,1,rep,name=buildCommands,proto3" json:"buildCommands,omitempty"`
	StartUpCommands   []*global.StartUpCommandsOutput   `protobuf:"bytes,2,rep,name=startUpCommands,proto3" json:"startUpCommands,omitempty"`
	SeedCommands      []*global.SeedCommandsOutput      `protobuf:"bytes,3,rep,name=seedCommands,proto3" json:"seedCommands,omitempty"`
	MigrationCommands []*global.MigrationCommandsOutput `protobuf:"bytes,4,rep,name=MigrationCommands,proto3" json:"MigrationCommands,omitempty"` //     todo:   DetectAdHocScripts
}

func (x *Commands) Reset() {
	*x = Commands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_decisionMaker_decisionMaker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commands) ProtoMessage() {}

func (x *Commands) ProtoReflect() protoreflect.Message {
	mi := &file_protos_decisionMaker_decisionMaker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commands.ProtoReflect.Descriptor instead.
func (*Commands) Descriptor() ([]byte, []int) {
	return file_protos_decisionMaker_decisionMaker_proto_rawDescGZIP(), []int{2}
}

func (x *Commands) GetBuildCommands() []*global.BuildCommandsOutput {
	if x != nil {
		return x.BuildCommands
	}
	return nil
}

func (x *Commands) GetStartUpCommands() []*global.StartUpCommandsOutput {
	if x != nil {
		return x.StartUpCommands
	}
	return nil
}

func (x *Commands) GetSeedCommands() []*global.SeedCommandsOutput {
	if x != nil {
		return x.SeedCommands
	}
	return nil
}

func (x *Commands) GetMigrationCommands() []*global.MigrationCommandsOutput {
	if x != nil {
		return x.MigrationCommands
	}
	return nil
}

type DecisionMakerInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LanguageSpecificDetection []*LanguageSpecificDetections `protobuf:"bytes,1,rep,name=languageSpecificDetection,proto3" json:"languageSpecificDetection,omitempty"`
	GloabalDetections         *GlobalDetections             `protobuf:"bytes,2,opt,name=gloabalDetections,proto3" json:"gloabalDetections,omitempty"`
}

func (x *DecisionMakerInput) Reset() {
	*x = DecisionMakerInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_decisionMaker_decisionMaker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecisionMakerInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionMakerInput) ProtoMessage() {}

func (x *DecisionMakerInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_decisionMaker_decisionMaker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionMakerInput.ProtoReflect.Descriptor instead.
func (*DecisionMakerInput) Descriptor() ([]byte, []int) {
	return file_protos_decisionMaker_decisionMaker_proto_rawDescGZIP(), []int{3}
}

func (x *DecisionMakerInput) GetLanguageSpecificDetection() []*LanguageSpecificDetections {
	if x != nil {
		return x.LanguageSpecificDetection
	}
	return nil
}

func (x *DecisionMakerInput) GetGloabalDetections() *GlobalDetections {
	if x != nil {
		return x.GloabalDetections
	}
	return nil
}

var File_protos_decisionMaker_decisionMaker_proto protoreflect.FileDescriptor

var file_protos_decisionMaker_decisionMaker_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x4d, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4d,
	0x61, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x6d, 0x61, 0x6b, 0x65,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f,
	0x73, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x34, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x2f, 0x65, 0x6e, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2f,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2f, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x07, 0x0a,
	0x1a, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e, 0x45, 0x6e, 0x76, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x4e, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x39, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x2e, 0x44, 0x42, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x02, 0x64, 0x62,
	0x12, 0x3c, 0x0a, 0x03, 0x6f, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e,
	0x4f, 0x72, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x03, 0x6f, 0x72, 0x6d, 0x12, 0x57,
	0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e, 0x4c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x4e,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x4e,
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x52, 0x09, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x5d,
	0x0a, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x25, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x22, 0xbb, 0x02, 0x0a, 0x10, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x63, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x6d,
	0x61, 0x6b, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x52, 0x08, 0x6d, 0x61, 0x6b, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x46, 0x0a,
	0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x46,
	0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x5b, 0x0a, 0x11, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52,
	0x11, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x22, 0xdd, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12,
	0x4f, 0x0a, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x52, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x12, 0x55, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x4c, 0x0a, 0x0c, 0x73, 0x65, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0c, 0x73, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x5b, 0x0a, 0x11, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52,
	0x11, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x22, 0xb0, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4d,
	0x61, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x59, 0x0a, 0x19, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x19, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x11, 0x67, 0x6c, 0x6f, 0x61, 0x62, 0x61, 0x6c, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x11, 0x67, 0x6c, 0x6f, 0x61, 0x62, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x19, 0x5a, 0x17, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_decisionMaker_decisionMaker_proto_rawDescOnce sync.Once
	file_protos_decisionMaker_decisionMaker_proto_rawDescData = file_protos_decisionMaker_decisionMaker_proto_rawDesc
)

func file_protos_decisionMaker_decisionMaker_proto_rawDescGZIP() []byte {
	file_protos_decisionMaker_decisionMaker_proto_rawDescOnce.Do(func() {
		file_protos_decisionMaker_decisionMaker_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_decisionMaker_decisionMaker_proto_rawDescData)
	})
	return file_protos_decisionMaker_decisionMaker_proto_rawDescData
}

var file_protos_decisionMaker_decisionMaker_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_decisionMaker_decisionMaker_proto_goTypes = []interface{}{
	(*LanguageSpecificDetections)(nil),            // 0: LanguageSpecificDetections
	(*GlobalDetections)(nil),                      // 1: GlobalDetections
	(*Commands)(nil),                              // 2: Commands
	(*DecisionMakerInput)(nil),                    // 3: DecisionMakerInput
	(*languageSpecific.EnvOutput)(nil),            // 4: protos.output.language_specific.EnvOutput
	(*languageSpecific.FrameworkOutput)(nil),      // 5: protos.output.language_specific.FrameworkOutput
	(*languageSpecific.DBOutput)(nil),             // 6: protos.output.language_specific.DBOutput
	(*languageSpecific.OrmOutput)(nil),            // 7: protos.output.language_specific.OrmOutput
	(*languageSpecific.DependenciesOutput)(nil),   // 8: protos.output.language_specific.DependenciesOutput
	(*languageSpecific.LibraryOutput)(nil),        // 9: protos.output.language_specific.LibraryOutput
	(*languageSpecific.StaticAssetsOutput)(nil),   // 10: protos.output.language_specific.StaticAssetsOutput
	(*languageSpecific.StackOutput)(nil),          // 11: protos.output.language_specific.StackOutput
	(*languageSpecific.AppserverOutput)(nil),      // 12: protos.output.language_specific.AppserverOutput
	(*languageSpecific.BuildDirectoryOutput)(nil), // 13: protos.output.language_specific.BuildDirectoryOutput
	(*global.ProcFileOutput)(nil),                 // 14: protos.output.global.ProcFileOutput
	(*global.MakefileOutput)(nil),                 // 15: protos.output.global.MakefileOutput
	(*global.DockerFileOutput)(nil),               // 16: protos.output.global.DockerFileOutput
	(*global.DockerComposeFileOutput)(nil),        // 17: protos.output.global.DockerComposeFileOutput
	(*global.BuildCommandsOutput)(nil),            // 18: protos.output.global.BuildCommandsOutput
	(*global.StartUpCommandsOutput)(nil),          // 19: protos.output.global.StartUpCommandsOutput
	(*global.SeedCommandsOutput)(nil),             // 20: protos.output.global.SeedCommandsOutput
	(*global.MigrationCommandsOutput)(nil),        // 21: protos.output.global.MigrationCommandsOutput
}
var file_protos_decisionMaker_decisionMaker_proto_depIdxs = []int32{
	4,  // 0: LanguageSpecificDetections.env:type_name -> protos.output.language_specific.EnvOutput
	5,  // 1: LanguageSpecificDetections.framework:type_name -> protos.output.language_specific.FrameworkOutput
	6,  // 2: LanguageSpecificDetections.db:type_name -> protos.output.language_specific.DBOutput
	7,  // 3: LanguageSpecificDetections.orm:type_name -> protos.output.language_specific.OrmOutput
	8,  // 4: LanguageSpecificDetections.dependencies:type_name -> protos.output.language_specific.DependenciesOutput
	9,  // 5: LanguageSpecificDetections.libraries:type_name -> protos.output.language_specific.LibraryOutput
	10, // 6: LanguageSpecificDetections.staticAssets:type_name -> protos.output.language_specific.StaticAssetsOutput
	11, // 7: LanguageSpecificDetections.stackOutput:type_name -> protos.output.language_specific.StackOutput
	12, // 8: LanguageSpecificDetections.appserver:type_name -> protos.output.language_specific.AppserverOutput
	13, // 9: LanguageSpecificDetections.buildDirectory:type_name -> protos.output.language_specific.BuildDirectoryOutput
	2,  // 10: LanguageSpecificDetections.commands:type_name -> Commands
	14, // 11: GlobalDetections.procFile:type_name -> protos.output.global.ProcFileOutput
	15, // 12: GlobalDetections.makefile:type_name -> protos.output.global.MakefileOutput
	16, // 13: GlobalDetections.dockerFile:type_name -> protos.output.global.DockerFileOutput
	17, // 14: GlobalDetections.dockerComposeFile:type_name -> protos.output.global.DockerComposeFileOutput
	18, // 15: Commands.buildCommands:type_name -> protos.output.global.BuildCommandsOutput
	19, // 16: Commands.startUpCommands:type_name -> protos.output.global.StartUpCommandsOutput
	20, // 17: Commands.seedCommands:type_name -> protos.output.global.SeedCommandsOutput
	21, // 18: Commands.MigrationCommands:type_name -> protos.output.global.MigrationCommandsOutput
	0,  // 19: DecisionMakerInput.languageSpecificDetection:type_name -> LanguageSpecificDetections
	1,  // 20: DecisionMakerInput.gloabalDetections:type_name -> GlobalDetections
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_protos_decisionMaker_decisionMaker_proto_init() }
func file_protos_decisionMaker_decisionMaker_proto_init() {
	if File_protos_decisionMaker_decisionMaker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_decisionMaker_decisionMaker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageSpecificDetections); i {
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
		file_protos_decisionMaker_decisionMaker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalDetections); i {
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
		file_protos_decisionMaker_decisionMaker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commands); i {
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
		file_protos_decisionMaker_decisionMaker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecisionMakerInput); i {
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
			RawDescriptor: file_protos_decisionMaker_decisionMaker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_decisionMaker_decisionMaker_proto_goTypes,
		DependencyIndexes: file_protos_decisionMaker_decisionMaker_proto_depIdxs,
		MessageInfos:      file_protos_decisionMaker_decisionMaker_proto_msgTypes,
	}.Build()
	File_protos_decisionMaker_decisionMaker_proto = out.File
	file_protos_decisionMaker_decisionMaker_proto_rawDesc = nil
	file_protos_decisionMaker_decisionMaker_proto_goTypes = nil
	file_protos_decisionMaker_decisionMaker_proto_depIdxs = nil
}
