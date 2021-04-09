// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/plugin/orm.proto

package plugin

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

type ServiceOutputDetectOrm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Used               bool          `protobuf:"varint,1,opt,name=used,proto3" json:"used,omitempty"`
	SupportedDbName    string        `protobuf:"bytes,2,opt,name=supportedDbName,proto3" json:"supportedDbName,omitempty"`
	SupportedDbVersion string        `protobuf:"bytes,3,opt,name=supportedDbVersion,proto3" json:"supportedDbVersion,omitempty"`
	Error              *ServiceError `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ServiceOutputDetectOrm) Reset() {
	*x = ServiceOutputDetectOrm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_orm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOutputDetectOrm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOutputDetectOrm) ProtoMessage() {}

func (x *ServiceOutputDetectOrm) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_orm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOutputDetectOrm.ProtoReflect.Descriptor instead.
func (*ServiceOutputDetectOrm) Descriptor() ([]byte, []int) {
	return file_protos_plugin_orm_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceOutputDetectOrm) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *ServiceOutputDetectOrm) GetSupportedDbName() string {
	if x != nil {
		return x.SupportedDbName
	}
	return ""
}

func (x *ServiceOutputDetectOrm) GetSupportedDbVersion() string {
	if x != nil {
		return x.SupportedDbVersion
	}
	return ""
}

func (x *ServiceOutputDetectOrm) GetError() *ServiceError {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_protos_plugin_orm_proto protoreflect.FileDescriptor

var file_protos_plugin_orm_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a,
	0x16, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x4f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x44, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x44,
	0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x44, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x44, 0x62, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0xca, 0x01, 0x0a, 0x0a, 0x4f, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3c, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x4f, 0x72, 0x6d, 0x12, 0x3a, 0x0a,
	0x09, 0x49, 0x73, 0x4f, 0x52, 0x4d, 0x55, 0x73, 0x65, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x42, 0x0a, 0x10, 0x50, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x4f, 0x66, 0x4f, 0x52, 0x4d, 0x55, 0x73, 0x65, 0x64, 0x12, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x42, 0x20, 0x5a,
	0x1e, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_plugin_orm_proto_rawDescOnce sync.Once
	file_protos_plugin_orm_proto_rawDescData = file_protos_plugin_orm_proto_rawDesc
)

func file_protos_plugin_orm_proto_rawDescGZIP() []byte {
	file_protos_plugin_orm_proto_rawDescOnce.Do(func() {
		file_protos_plugin_orm_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_plugin_orm_proto_rawDescData)
	})
	return file_protos_plugin_orm_proto_rawDescData
}

var file_protos_plugin_orm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_plugin_orm_proto_goTypes = []interface{}{
	(*ServiceOutputDetectOrm)(nil), // 0: proto.ServiceOutputDetectOrm
	(*ServiceError)(nil),           // 1: proto.ServiceError
	(*ServiceInput)(nil),           // 2: proto.ServiceInput
	(*ServiceOutputBool)(nil),      // 3: proto.ServiceOutputBool
	(*ServiceOutputFloat)(nil),     // 4: proto.ServiceOutputFloat
}
var file_protos_plugin_orm_proto_depIdxs = []int32{
	1, // 0: proto.ServiceOutputDetectOrm.error:type_name -> proto.ServiceError
	2, // 1: proto.OrmService.Detect:input_type -> proto.ServiceInput
	2, // 2: proto.OrmService.IsORMUsed:input_type -> proto.ServiceInput
	2, // 3: proto.OrmService.PercentOfORMUsed:input_type -> proto.ServiceInput
	0, // 4: proto.OrmService.Detect:output_type -> proto.ServiceOutputDetectOrm
	3, // 5: proto.OrmService.IsORMUsed:output_type -> proto.ServiceOutputBool
	4, // 6: proto.OrmService.PercentOfORMUsed:output_type -> proto.ServiceOutputFloat
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_plugin_orm_proto_init() }
func file_protos_plugin_orm_proto_init() {
	if File_protos_plugin_orm_proto != nil {
		return
	}
	file_protos_plugin_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_plugin_orm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOutputDetectOrm); i {
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
			RawDescriptor: file_protos_plugin_orm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_plugin_orm_proto_goTypes,
		DependencyIndexes: file_protos_plugin_orm_proto_depIdxs,
		MessageInfos:      file_protos_plugin_orm_proto_msgTypes,
	}.Build()
	File_protos_plugin_orm_proto = out.File
	file_protos_plugin_orm_proto_rawDesc = nil
	file_protos_plugin_orm_proto_goTypes = nil
	file_protos_plugin_orm_proto_depIdxs = nil
}
