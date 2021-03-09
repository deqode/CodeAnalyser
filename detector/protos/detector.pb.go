// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: detector.proto

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

type DetectedOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// checkers....
	FrameworkUsed   bool `protobuf:"varint,1,opt,name=frameworkUsed,proto3" json:"frameworkUsed,omitempty"`
	EnvUsed         bool `protobuf:"varint,2,opt,name=envUsed,proto3" json:"envUsed,omitempty"`
	StartUpCommands bool `protobuf:"varint,3,opt,name=startUpCommands,proto3" json:"startUpCommands,omitempty"`
	DbUsed          bool `protobuf:"varint,4,opt,name=dbUsed,proto3" json:"dbUsed,omitempty"`
	ORMUsed         bool `protobuf:"varint,5,opt,name=ORMUsed,json=oRMUsed,proto3" json:"ORMUsed,omitempty"`
	//runtime version used
	RuntimeVersion string `protobuf:"bytes,6,opt,name=runtimeVersion,proto3" json:"runtimeVersion,omitempty"`
	// list of envs
	Envs []string `protobuf:"bytes,7,rep,name=envs,proto3" json:"envs,omitempty"`
	// list of start-up commands
	Commands []string `protobuf:"bytes,8,rep,name=commands,proto3" json:"commands,omitempty"`
	// list of libraries used
	Libraries []string `protobuf:"bytes,9,rep,name=libraries,proto3" json:"libraries,omitempty"`
	// list of port used
	Ports []*Port `protobuf:"bytes,10,rep,name=ports,proto3" json:"ports,omitempty"`
	// list of port databases
	Databases []*DBPort `protobuf:"bytes,11,rep,name=databases,proto3" json:"databases,omitempty"`
	// list of DB ORMs
	ORMs []string `protobuf:"bytes,12,rep,name=ORMs,json=oRMs,proto3" json:"ORMs,omitempty"`
}

func (x *DetectedOutput) Reset() {
	*x = DetectedOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectedOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectedOutput) ProtoMessage() {}

func (x *DetectedOutput) ProtoReflect() protoreflect.Message {
	mi := &file_detector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectedOutput.ProtoReflect.Descriptor instead.
func (*DetectedOutput) Descriptor() ([]byte, []int) {
	return file_detector_proto_rawDescGZIP(), []int{0}
}

func (x *DetectedOutput) GetFrameworkUsed() bool {
	if x != nil {
		return x.FrameworkUsed
	}
	return false
}

func (x *DetectedOutput) GetEnvUsed() bool {
	if x != nil {
		return x.EnvUsed
	}
	return false
}

func (x *DetectedOutput) GetStartUpCommands() bool {
	if x != nil {
		return x.StartUpCommands
	}
	return false
}

func (x *DetectedOutput) GetDbUsed() bool {
	if x != nil {
		return x.DbUsed
	}
	return false
}

func (x *DetectedOutput) GetORMUsed() bool {
	if x != nil {
		return x.ORMUsed
	}
	return false
}

func (x *DetectedOutput) GetRuntimeVersion() string {
	if x != nil {
		return x.RuntimeVersion
	}
	return ""
}

func (x *DetectedOutput) GetEnvs() []string {
	if x != nil {
		return x.Envs
	}
	return nil
}

func (x *DetectedOutput) GetCommands() []string {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *DetectedOutput) GetLibraries() []string {
	if x != nil {
		return x.Libraries
	}
	return nil
}

func (x *DetectedOutput) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *DetectedOutput) GetDatabases() []*DBPort {
	if x != nil {
		return x.Databases
	}
	return nil
}

func (x *DetectedOutput) GetORMs() []string {
	if x != nil {
		return x.ORMs
	}
	return nil
}

// exposed port
type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_detector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_detector_proto_rawDescGZIP(), []int{1}
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

// exposed DB port
type DBPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port    string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *DBPort) Reset() {
	*x = DBPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBPort) ProtoMessage() {}

func (x *DBPort) ProtoReflect() protoreflect.Message {
	mi := &file_detector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBPort.ProtoReflect.Descriptor instead.
func (*DBPort) Descriptor() ([]byte, []int) {
	return file_detector_proto_rawDescGZIP(), []int{2}
}

func (x *DBPort) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DBPort) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *DBPort) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_detector_proto protoreflect.FileDescriptor

var file_detector_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfa, 0x02, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x55, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x76,
	0x55, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x55,
	0x73, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x55, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x62, 0x55, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64,
	0x62, 0x55, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x52, 0x4d, 0x55, 0x73, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x52, 0x4d, 0x55, 0x73, 0x65, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x12, 0x25, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x44, 0x42, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x52, 0x4d,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x52, 0x4d, 0x73, 0x22, 0x2e, 0x0a,
	0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x4a, 0x0a,
	0x06, 0x44, 0x42, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_detector_proto_rawDescOnce sync.Once
	file_detector_proto_rawDescData = file_detector_proto_rawDesc
)

func file_detector_proto_rawDescGZIP() []byte {
	file_detector_proto_rawDescOnce.Do(func() {
		file_detector_proto_rawDescData = protoimpl.X.CompressGZIP(file_detector_proto_rawDescData)
	})
	return file_detector_proto_rawDescData
}

var file_detector_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_detector_proto_goTypes = []interface{}{
	(*DetectedOutput)(nil), // 0: DetectedOutput
	(*Port)(nil),           // 1: Port
	(*DBPort)(nil),         // 2: DBPort
}
var file_detector_proto_depIdxs = []int32{
	1, // 0: DetectedOutput.ports:type_name -> Port
	2, // 1: DetectedOutput.databases:type_name -> DBPort
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_detector_proto_init() }
func file_detector_proto_init() {
	if File_detector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_detector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectedOutput); i {
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
		file_detector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_detector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBPort); i {
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
			RawDescriptor: file_detector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_detector_proto_goTypes,
		DependencyIndexes: file_detector_proto_depIdxs,
		MessageInfos:      file_detector_proto_msgTypes,
	}.Build()
	File_detector_proto = out.File
	file_detector_proto_rawDesc = nil
	file_detector_proto_goTypes = nil
	file_detector_proto_depIdxs = nil
}
