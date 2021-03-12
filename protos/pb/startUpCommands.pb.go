// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: protos/outputs/global/startUpCommands.proto

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

type StartUpCommandsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// check startUpCommands
	Used bool `protobuf:"varint,1,opt,name=used,proto3" json:"used,omitempty"`
	// list of all startUpCommands
	StartUpCommands []*StartUpCommand `protobuf:"bytes,2,rep,name=startUpCommands,proto3" json:"startUpCommands,omitempty"`
}

func (x *StartUpCommandsOutput) Reset() {
	*x = StartUpCommandsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_outputs_global_startUpCommands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartUpCommandsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartUpCommandsOutput) ProtoMessage() {}

func (x *StartUpCommandsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_outputs_global_startUpCommands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartUpCommandsOutput.ProtoReflect.Descriptor instead.
func (*StartUpCommandsOutput) Descriptor() ([]byte, []int) {
	return file_protos_outputs_global_startUpCommands_proto_rawDescGZIP(), []int{0}
}

func (x *StartUpCommandsOutput) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *StartUpCommandsOutput) GetStartUpCommands() []*StartUpCommand {
	if x != nil {
		return x.StartUpCommands
	}
	return nil
}

type StartUpCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  name of StartUpCommand
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//  version of StartUpCommand
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *StartUpCommand) Reset() {
	*x = StartUpCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_outputs_global_startUpCommands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartUpCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartUpCommand) ProtoMessage() {}

func (x *StartUpCommand) ProtoReflect() protoreflect.Message {
	mi := &file_protos_outputs_global_startUpCommands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartUpCommand.ProtoReflect.Descriptor instead.
func (*StartUpCommand) Descriptor() ([]byte, []int) {
	return file_protos_outputs_global_startUpCommands_proto_rawDescGZIP(), []int{1}
}

func (x *StartUpCommand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StartUpCommand) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

var File_protos_outputs_global_startUpCommands_proto protoreflect.FileDescriptor

var file_protos_outputs_global_startUpCommands_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a,
	0x15, 0x53, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x55, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x38, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_outputs_global_startUpCommands_proto_rawDescOnce sync.Once
	file_protos_outputs_global_startUpCommands_proto_rawDescData = file_protos_outputs_global_startUpCommands_proto_rawDesc
)

func file_protos_outputs_global_startUpCommands_proto_rawDescGZIP() []byte {
	file_protos_outputs_global_startUpCommands_proto_rawDescOnce.Do(func() {
		file_protos_outputs_global_startUpCommands_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_outputs_global_startUpCommands_proto_rawDescData)
	})
	return file_protos_outputs_global_startUpCommands_proto_rawDescData
}

var file_protos_outputs_global_startUpCommands_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_outputs_global_startUpCommands_proto_goTypes = []interface{}{
	(*StartUpCommandsOutput)(nil), // 0: StartUpCommandsOutput
	(*StartUpCommand)(nil),        // 1: StartUpCommand
}
var file_protos_outputs_global_startUpCommands_proto_depIdxs = []int32{
	1, // 0: StartUpCommandsOutput.startUpCommands:type_name -> StartUpCommand
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_outputs_global_startUpCommands_proto_init() }
func file_protos_outputs_global_startUpCommands_proto_init() {
	if File_protos_outputs_global_startUpCommands_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_outputs_global_startUpCommands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartUpCommandsOutput); i {
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
		file_protos_outputs_global_startUpCommands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartUpCommand); i {
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
			RawDescriptor: file_protos_outputs_global_startUpCommands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_outputs_global_startUpCommands_proto_goTypes,
		DependencyIndexes: file_protos_outputs_global_startUpCommands_proto_depIdxs,
		MessageInfos:      file_protos_outputs_global_startUpCommands_proto_msgTypes,
	}.Build()
	File_protos_outputs_global_startUpCommands_proto = out.File
	file_protos_outputs_global_startUpCommands_proto_rawDesc = nil
	file_protos_outputs_global_startUpCommands_proto_goTypes = nil
	file_protos_outputs_global_startUpCommands_proto_depIdxs = nil
}
