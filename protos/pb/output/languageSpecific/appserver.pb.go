// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: protos/output/language_specific/appserver.proto

package languageSpecific

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

type AppserverOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// check appserver
	Used bool `protobuf:"varint,1,opt,name=used,proto3" json:"used,omitempty"`
	// list of appserver
	AppServers []*Appserver `protobuf:"bytes,2,rep,name=appServers,proto3" json:"appServers,omitempty"`
}

func (x *AppserverOutput) Reset() {
	*x = AppserverOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_output_language_specific_appserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppserverOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppserverOutput) ProtoMessage() {}

func (x *AppserverOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_output_language_specific_appserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppserverOutput.ProtoReflect.Descriptor instead.
func (*AppserverOutput) Descriptor() ([]byte, []int) {
	return file_protos_output_language_specific_appserver_proto_rawDescGZIP(), []int{0}
}

func (x *AppserverOutput) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *AppserverOutput) GetAppServers() []*Appserver {
	if x != nil {
		return x.AppServers
	}
	return nil
}

type Appserver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  name of appserver
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//  version of appserver
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Appserver) Reset() {
	*x = Appserver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_output_language_specific_appserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Appserver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Appserver) ProtoMessage() {}

func (x *Appserver) ProtoReflect() protoreflect.Message {
	mi := &file_protos_output_language_specific_appserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Appserver.ProtoReflect.Descriptor instead.
func (*Appserver) Descriptor() ([]byte, []int) {
	return file_protos_output_language_specific_appserver_proto_rawDescGZIP(), []int{1}
}

func (x *Appserver) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Appserver) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_protos_output_language_specific_appserver_proto protoreflect.FileDescriptor

var file_protos_output_language_specific_appserver_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x22, 0x71, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x0a, 0x61, 0x70, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e,
	0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x39, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x31, 0x5a, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_output_language_specific_appserver_proto_rawDescOnce sync.Once
	file_protos_output_language_specific_appserver_proto_rawDescData = file_protos_output_language_specific_appserver_proto_rawDesc
)

func file_protos_output_language_specific_appserver_proto_rawDescGZIP() []byte {
	file_protos_output_language_specific_appserver_proto_rawDescOnce.Do(func() {
		file_protos_output_language_specific_appserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_output_language_specific_appserver_proto_rawDescData)
	})
	return file_protos_output_language_specific_appserver_proto_rawDescData
}

var file_protos_output_language_specific_appserver_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_output_language_specific_appserver_proto_goTypes = []interface{}{
	(*AppserverOutput)(nil), // 0: protos.output.language_specific.AppserverOutput
	(*Appserver)(nil),       // 1: protos.output.language_specific.Appserver
}
var file_protos_output_language_specific_appserver_proto_depIdxs = []int32{
	1, // 0: protos.output.language_specific.AppserverOutput.appServers:type_name -> protos.output.language_specific.Appserver
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_output_language_specific_appserver_proto_init() }
func file_protos_output_language_specific_appserver_proto_init() {
	if File_protos_output_language_specific_appserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_output_language_specific_appserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppserverOutput); i {
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
		file_protos_output_language_specific_appserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Appserver); i {
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
			RawDescriptor: file_protos_output_language_specific_appserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_output_language_specific_appserver_proto_goTypes,
		DependencyIndexes: file_protos_output_language_specific_appserver_proto_depIdxs,
		MessageInfos:      file_protos_output_language_specific_appserver_proto_msgTypes,
	}.Build()
	File_protos_output_language_specific_appserver_proto = out.File
	file_protos_output_language_specific_appserver_proto_rawDesc = nil
	file_protos_output_language_specific_appserver_proto_goTypes = nil
	file_protos_output_language_specific_appserver_proto_depIdxs = nil
}
