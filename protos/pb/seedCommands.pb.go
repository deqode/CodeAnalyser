// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: protos/outputs/global/seedCommands.proto

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

type SeedCommandsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// check SeedCommand
	Used bool `protobuf:"varint,1,opt,name=used,proto3" json:"used,omitempty"`
	// list of all SeedCommand
	SeedCommandCommand []*SeedCommand `protobuf:"bytes,2,rep,name=SeedCommandCommand,json=seedCommandCommand,proto3" json:"SeedCommandCommand,omitempty"`
}

func (x *SeedCommandsOutput) Reset() {
	*x = SeedCommandsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_outputs_global_seedCommands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeedCommandsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeedCommandsOutput) ProtoMessage() {}

func (x *SeedCommandsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_outputs_global_seedCommands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeedCommandsOutput.ProtoReflect.Descriptor instead.
func (*SeedCommandsOutput) Descriptor() ([]byte, []int) {
	return file_protos_outputs_global_seedCommands_proto_rawDescGZIP(), []int{0}
}

func (x *SeedCommandsOutput) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *SeedCommandsOutput) GetSeedCommandCommand() []*SeedCommand {
	if x != nil {
		return x.SeedCommandCommand
	}
	return nil
}

type SeedCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  name of MigrationCommand
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//  version of MigrationCommand
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *SeedCommand) Reset() {
	*x = SeedCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_outputs_global_seedCommands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeedCommand) ProtoMessage() {}

func (x *SeedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_protos_outputs_global_seedCommands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeedCommand.ProtoReflect.Descriptor instead.
func (*SeedCommand) Descriptor() ([]byte, []int) {
	return file_protos_outputs_global_seedCommands_proto_rawDescGZIP(), []int{1}
}

func (x *SeedCommand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SeedCommand) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

var File_protos_outputs_global_seedCommands_proto protoreflect.FileDescriptor

var file_protos_outputs_global_seedCommands_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x12, 0x53, 0x65,
	0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x12, 0x53, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x53, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x12,
	0x73, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x22, 0x35, 0x0a, 0x0b, 0x53, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_outputs_global_seedCommands_proto_rawDescOnce sync.Once
	file_protos_outputs_global_seedCommands_proto_rawDescData = file_protos_outputs_global_seedCommands_proto_rawDesc
)

func file_protos_outputs_global_seedCommands_proto_rawDescGZIP() []byte {
	file_protos_outputs_global_seedCommands_proto_rawDescOnce.Do(func() {
		file_protos_outputs_global_seedCommands_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_outputs_global_seedCommands_proto_rawDescData)
	})
	return file_protos_outputs_global_seedCommands_proto_rawDescData
}

var file_protos_outputs_global_seedCommands_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_outputs_global_seedCommands_proto_goTypes = []interface{}{
	(*SeedCommandsOutput)(nil), // 0: SeedCommandsOutput
	(*SeedCommand)(nil),        // 1: SeedCommand
}
var file_protos_outputs_global_seedCommands_proto_depIdxs = []int32{
	1, // 0: SeedCommandsOutput.SeedCommandCommand:type_name -> SeedCommand
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_outputs_global_seedCommands_proto_init() }
func file_protos_outputs_global_seedCommands_proto_init() {
	if File_protos_outputs_global_seedCommands_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_outputs_global_seedCommands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeedCommandsOutput); i {
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
		file_protos_outputs_global_seedCommands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeedCommand); i {
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
			RawDescriptor: file_protos_outputs_global_seedCommands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_outputs_global_seedCommands_proto_goTypes,
		DependencyIndexes: file_protos_outputs_global_seedCommands_proto_depIdxs,
		MessageInfos:      file_protos_outputs_global_seedCommands_proto_msgTypes,
	}.Build()
	File_protos_outputs_global_seedCommands_proto = out.File
	file_protos_outputs_global_seedCommands_proto_rawDesc = nil
	file_protos_outputs_global_seedCommands_proto_goTypes = nil
	file_protos_outputs_global_seedCommands_proto_depIdxs = nil
}
