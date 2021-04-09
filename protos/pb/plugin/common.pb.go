// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/plugin/common.proto

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

type ServiceError_ErrorCode int32

const (
	ServiceError_Error ServiceError_ErrorCode = 0 //TODO discuss error codes and types
)

// Enum value maps for ServiceError_ErrorCode.
var (
	ServiceError_ErrorCode_name = map[int32]string{
		0: "Error",
	}
	ServiceError_ErrorCode_value = map[string]int32{
		"Error": 0,
	}
)

func (x ServiceError_ErrorCode) Enum() *ServiceError_ErrorCode {
	p := new(ServiceError_ErrorCode)
	*p = x
	return p
}

func (x ServiceError_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceError_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_plugin_common_proto_enumTypes[0].Descriptor()
}

func (ServiceError_ErrorCode) Type() protoreflect.EnumType {
	return &file_protos_plugin_common_proto_enumTypes[0]
}

func (x ServiceError_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceError_ErrorCode.Descriptor instead.
func (ServiceError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_protos_plugin_common_proto_rawDescGZIP(), []int{9, 0}
}

type ServiceEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceEmpty) Reset() {
	*x = ServiceEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceEmpty) ProtoMessage() {}

func (x *ServiceEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceEmpty.ProtoReflect.Descriptor instead.
func (*ServiceEmpty) Descriptor() ([]byte, []int) {
	return file_protos_plugin_common_proto_rawDescGZIP(), []int{0}
}

type ServiceOutputString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string        `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Error *ServiceError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ServiceOutputString) Reset() {
	*x = ServiceOutputString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOutputString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOutputString) ProtoMessage() {}

func (x *ServiceOutputString) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOutputString.ProtoReflect.Descriptor instead.
func (*ServiceOutputString) Descriptor() ([]byte, []int) {
	return file_protos_plugin_common_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceOutputString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ServiceOutputString) GetError() *ServiceError {
	if x != nil {
		return x.Error
	}
	return nil
}

type ServiceInputString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ServiceInputString) Reset() {
	*x = ServiceInputString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInputString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInputString) ProtoMessage() {}

func (x *ServiceInputString) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInputString.ProtoReflect.Descriptor instead.
func (*ServiceInputString) Descriptor() ([]byte, []int) {
	return file_protos_plugin_common_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceInputString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ServiceOutputInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64         `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Error *ServiceError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ServiceOutputInt) Reset() {
	*x = ServiceOutputInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOutputInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOutputInt) ProtoMessage() {}

func (x *ServiceOutputInt) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOutputInt.ProtoReflect.Descriptor instead.
func (*ServiceOutputInt) Descriptor() ([]byte, []int) {
	return file_protos_plugin_common_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceOutputInt) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ServiceOutputInt) GetError() *ServiceError {
	if x != nil {
		return x.Error
	}
	return nil
}

type ServiceEmptyOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *ServiceError `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ServiceEmptyOutput) Reset() {
	*x = ServiceEmptyOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceEmptyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceEmptyOutput) ProtoMessage() {}

func (x *ServiceEmptyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceEmptyOutput.ProtoReflect.Descriptor instead.
func (*ServiceEmptyOutput) Descriptor() ([]byte, []int) {
	return file_protos_plugin_common_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceEmptyOutput) GetError() *ServiceError {
	if x != nil {
		return x.Error
	}
	return nil
}

type ServiceOutputFloat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float32       `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	Error *ServiceError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ServiceOutputFloat) Reset() {
	*x = ServiceOutputFloat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOutputFloat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOutputFloat) ProtoMessage() {}

func (x *ServiceOutputFloat) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOutputFloat.ProtoReflect.Descriptor instead.
func (*ServiceOutputFloat) Descriptor() ([]byte, []int) {
	return file_protos_plugin_common_proto_rawDescGZIP(), []int{5}
}

func (x *ServiceOutputFloat) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ServiceOutputFloat) GetError() *ServiceError {
	if x != nil {
		return x.Error
	}
	return nil
}

type ServiceOutputBool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool          `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Error *ServiceError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ServiceOutputBool) Reset() {
	*x = ServiceOutputBool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOutputBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOutputBool) ProtoMessage() {}

func (x *ServiceOutputBool) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOutputBool.ProtoReflect.Descriptor instead.
func (*ServiceOutputBool) Descriptor() ([]byte, []int) {
	return file_protos_plugin_common_proto_rawDescGZIP(), []int{6}
}

func (x *ServiceOutputBool) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

func (x *ServiceOutputBool) GetError() *ServiceError {
	if x != nil {
		return x.Error
	}
	return nil
}

type ServiceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuntimeVersion string `protobuf:"bytes,1,opt,name=runtimeVersion,proto3" json:"runtimeVersion,omitempty"`
	Root           string `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
}

func (x *ServiceInput) Reset() {
	*x = ServiceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInput) ProtoMessage() {}

func (x *ServiceInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInput.ProtoReflect.Descriptor instead.
func (*ServiceInput) Descriptor() ([]byte, []int) {
	return file_protos_plugin_common_proto_rawDescGZIP(), []int{7}
}

func (x *ServiceInput) GetRuntimeVersion() string {
	if x != nil {
		return x.RuntimeVersion
	}
	return ""
}

func (x *ServiceInput) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

type ServiceOutputStringMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Error *ServiceError     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ServiceOutputStringMap) Reset() {
	*x = ServiceOutputStringMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOutputStringMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOutputStringMap) ProtoMessage() {}

func (x *ServiceOutputStringMap) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOutputStringMap.ProtoReflect.Descriptor instead.
func (*ServiceOutputStringMap) Descriptor() ([]byte, []int) {
	return file_protos_plugin_common_proto_rawDescGZIP(), []int{8}
}

func (x *ServiceOutputStringMap) GetValue() map[string]string {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ServiceOutputStringMap) GetError() *ServiceError {
	if x != nil {
		return x.Error
	}
	return nil
}

type ServiceError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ErrorType ServiceError_ErrorCode `protobuf:"varint,2,opt,name=errorType,proto3,enum=proto.ServiceError_ErrorCode" json:"errorType,omitempty"`
	Cause     string                 `protobuf:"bytes,3,opt,name=cause,proto3" json:"cause,omitempty"`
}

func (x *ServiceError) Reset() {
	*x = ServiceError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceError) ProtoMessage() {}

func (x *ServiceError) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceError.ProtoReflect.Descriptor instead.
func (*ServiceError) Descriptor() ([]byte, []int) {
	return file_protos_plugin_common_proto_rawDescGZIP(), []int{9}
}

func (x *ServiceError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ServiceError) GetErrorType() ServiceError_ErrorCode {
	if x != nil {
		return x.ErrorType
	}
	return ServiceError_Error
}

func (x *ServiceError) GetCause() string {
	if x != nil {
		return x.Cause
	}
	return ""
}

var File_protos_plugin_common_proto protoreflect.FileDescriptor

var file_protos_plugin_common_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x56, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2a, 0x0a, 0x12, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x53, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x12,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x55, 0x0a,
	0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x54, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4a, 0x0a, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x22, 0xbd, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61,
	0x70, 0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x38, 0x0a, 0x0a,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x93, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x3b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x61, 0x75, 0x73, 0x65, 0x22, 0x16, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x00, 0x42, 0x20, 0x5a, 0x1e,
	0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_plugin_common_proto_rawDescOnce sync.Once
	file_protos_plugin_common_proto_rawDescData = file_protos_plugin_common_proto_rawDesc
)

func file_protos_plugin_common_proto_rawDescGZIP() []byte {
	file_protos_plugin_common_proto_rawDescOnce.Do(func() {
		file_protos_plugin_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_plugin_common_proto_rawDescData)
	})
	return file_protos_plugin_common_proto_rawDescData
}

var file_protos_plugin_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_plugin_common_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protos_plugin_common_proto_goTypes = []interface{}{
	(ServiceError_ErrorCode)(0),    // 0: proto.ServiceError.ErrorCode
	(*ServiceEmpty)(nil),           // 1: proto.ServiceEmpty
	(*ServiceOutputString)(nil),    // 2: proto.ServiceOutputString
	(*ServiceInputString)(nil),     // 3: proto.ServiceInputString
	(*ServiceOutputInt)(nil),       // 4: proto.ServiceOutputInt
	(*ServiceEmptyOutput)(nil),     // 5: proto.ServiceEmptyOutput
	(*ServiceOutputFloat)(nil),     // 6: proto.ServiceOutputFloat
	(*ServiceOutputBool)(nil),      // 7: proto.ServiceOutputBool
	(*ServiceInput)(nil),           // 8: proto.ServiceInput
	(*ServiceOutputStringMap)(nil), // 9: proto.ServiceOutputStringMap
	(*ServiceError)(nil),           // 10: proto.ServiceError
	nil,                            // 11: proto.ServiceOutputStringMap.ValueEntry
}
var file_protos_plugin_common_proto_depIdxs = []int32{
	10, // 0: proto.ServiceOutputString.error:type_name -> proto.ServiceError
	10, // 1: proto.ServiceOutputInt.error:type_name -> proto.ServiceError
	10, // 2: proto.ServiceEmptyOutput.error:type_name -> proto.ServiceError
	10, // 3: proto.ServiceOutputFloat.error:type_name -> proto.ServiceError
	10, // 4: proto.ServiceOutputBool.error:type_name -> proto.ServiceError
	11, // 5: proto.ServiceOutputStringMap.value:type_name -> proto.ServiceOutputStringMap.ValueEntry
	10, // 6: proto.ServiceOutputStringMap.error:type_name -> proto.ServiceError
	0,  // 7: proto.ServiceError.errorType:type_name -> proto.ServiceError.ErrorCode
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_protos_plugin_common_proto_init() }
func file_protos_plugin_common_proto_init() {
	if File_protos_plugin_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_plugin_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceEmpty); i {
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
		file_protos_plugin_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOutputString); i {
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
		file_protos_plugin_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInputString); i {
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
		file_protos_plugin_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOutputInt); i {
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
		file_protos_plugin_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceEmptyOutput); i {
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
		file_protos_plugin_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOutputFloat); i {
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
		file_protos_plugin_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOutputBool); i {
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
		file_protos_plugin_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInput); i {
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
		file_protos_plugin_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOutputStringMap); i {
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
		file_protos_plugin_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceError); i {
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
			RawDescriptor: file_protos_plugin_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_plugin_common_proto_goTypes,
		DependencyIndexes: file_protos_plugin_common_proto_depIdxs,
		EnumInfos:         file_protos_plugin_common_proto_enumTypes,
		MessageInfos:      file_protos_plugin_common_proto_msgTypes,
	}.Build()
	File_protos_plugin_common_proto = out.File
	file_protos_plugin_common_proto_rawDesc = nil
	file_protos_plugin_common_proto_goTypes = nil
	file_protos_plugin_common_proto_depIdxs = nil
}
