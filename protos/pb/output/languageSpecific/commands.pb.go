// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/output/language_specific/commands.proto

package languageSpecific

import (
	plugin "code-analyser/protos/pb/plugin"
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

type Commands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildCommands     *plugin.CommandsOutput `protobuf:"bytes,1,opt,name=buildCommands,proto3" json:"buildCommands,omitempty"`
	StartUpCommands   *plugin.CommandsOutput `protobuf:"bytes,2,opt,name=startUpCommands,proto3" json:"startUpCommands,omitempty"`
	SeedCommands      *plugin.CommandsOutput `protobuf:"bytes,3,opt,name=seedCommands,proto3" json:"seedCommands,omitempty"`
	MigrationCommands *plugin.CommandsOutput `protobuf:"bytes,4,opt,name=MigrationCommands,proto3" json:"MigrationCommands,omitempty"`
	AdHocScripts      *plugin.CommandsOutput `protobuf:"bytes,5,opt,name=adHocScripts,proto3" json:"adHocScripts,omitempty"`
}

func (x *Commands) Reset() {
	*x = Commands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_output_language_specific_commands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commands) ProtoMessage() {}

func (x *Commands) ProtoReflect() protoreflect.Message {
	mi := &file_protos_output_language_specific_commands_proto_msgTypes[0]
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
	return file_protos_output_language_specific_commands_proto_rawDescGZIP(), []int{0}
}

func (x *Commands) GetBuildCommands() *plugin.CommandsOutput {
	if x != nil {
		return x.BuildCommands
	}
	return nil
}

func (x *Commands) GetStartUpCommands() *plugin.CommandsOutput {
	if x != nil {
		return x.StartUpCommands
	}
	return nil
}

func (x *Commands) GetSeedCommands() *plugin.CommandsOutput {
	if x != nil {
		return x.SeedCommands
	}
	return nil
}

func (x *Commands) GetMigrationCommands() *plugin.CommandsOutput {
	if x != nil {
		return x.MigrationCommands
	}
	return nil
}

func (x *Commands) GetAdHocScripts() *plugin.CommandsOutput {
	if x != nil {
		return x.AdHocScripts
	}
	return nil
}

var File_protos_output_language_specific_commands_proto protoreflect.FileDescriptor

var file_protos_output_language_specific_commands_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc3, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x3b, 0x0a, 0x0d,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0d, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x3f, 0x0a, 0x0f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x55, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x55, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0c, 0x73, 0x65,
	0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0c, 0x73, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x43, 0x0a, 0x11, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x11, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0c, 0x61, 0x64,
	0x48, 0x6f, 0x63, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0c, 0x61, 0x64, 0x48, 0x6f, 0x63, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62,
	0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_output_language_specific_commands_proto_rawDescOnce sync.Once
	file_protos_output_language_specific_commands_proto_rawDescData = file_protos_output_language_specific_commands_proto_rawDesc
)

func file_protos_output_language_specific_commands_proto_rawDescGZIP() []byte {
	file_protos_output_language_specific_commands_proto_rawDescOnce.Do(func() {
		file_protos_output_language_specific_commands_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_output_language_specific_commands_proto_rawDescData)
	})
	return file_protos_output_language_specific_commands_proto_rawDescData
}

var file_protos_output_language_specific_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_output_language_specific_commands_proto_goTypes = []interface{}{
	(*Commands)(nil),              // 0: protos.output.language_specific.Commands
	(*plugin.CommandsOutput)(nil), // 1: proto.CommandsOutput
}
var file_protos_output_language_specific_commands_proto_depIdxs = []int32{
	1, // 0: protos.output.language_specific.Commands.buildCommands:type_name -> proto.CommandsOutput
	1, // 1: protos.output.language_specific.Commands.startUpCommands:type_name -> proto.CommandsOutput
	1, // 2: protos.output.language_specific.Commands.seedCommands:type_name -> proto.CommandsOutput
	1, // 3: protos.output.language_specific.Commands.MigrationCommands:type_name -> proto.CommandsOutput
	1, // 4: protos.output.language_specific.Commands.adHocScripts:type_name -> proto.CommandsOutput
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_output_language_specific_commands_proto_init() }
func file_protos_output_language_specific_commands_proto_init() {
	if File_protos_output_language_specific_commands_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_output_language_specific_commands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_output_language_specific_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_output_language_specific_commands_proto_goTypes,
		DependencyIndexes: file_protos_output_language_specific_commands_proto_depIdxs,
		MessageInfos:      file_protos_output_language_specific_commands_proto_msgTypes,
	}.Build()
	File_protos_output_language_specific_commands_proto = out.File
	file_protos_output_language_specific_commands_proto_rawDesc = nil
	file_protos_output_language_specific_commands_proto_goTypes = nil
	file_protos_output_language_specific_commands_proto_depIdxs = nil
}
