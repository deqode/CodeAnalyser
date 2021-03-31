// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/outputs/languageSpecific/appserver.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AppserverOutput struct {
	// check appserver
	Used bool `protobuf:"varint,1,opt,name=used" json:"used,omitempty"`
	// list of appserver
	AppServers []*Appserver `protobuf:"bytes,2,rep,name=appServers" json:"appServers,omitempty"`
}

func (m *AppserverOutput) Reset()                    { *m = AppserverOutput{} }
func (m *AppserverOutput) String() string            { return proto.CompactTextString(m) }
func (*AppserverOutput) ProtoMessage()               {}
func (*AppserverOutput) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *AppserverOutput) GetUsed() bool {
	if m != nil {
		return m.Used
	}
	return false
}

func (m *AppserverOutput) GetAppServers() []*Appserver {
	if m != nil {
		return m.AppServers
	}
	return nil
}

type Appserver struct {
	//  name of appserver
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	//  version of appserver
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *Appserver) Reset()                    { *m = Appserver{} }
func (m *Appserver) String() string            { return proto.CompactTextString(m) }
func (*Appserver) ProtoMessage()               {}
func (*Appserver) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *Appserver) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Appserver) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*AppserverOutput)(nil), "AppserverOutput")
	proto.RegisterType((*Appserver)(nil), "Appserver")
}

func init() { proto.RegisterFile("protos/outputs/languageSpecific/appserver.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0xcf, 0x2f, 0x2d, 0x29, 0x28, 0x2d, 0x29, 0xd6, 0xcf, 0x49, 0xcc, 0x4b, 0x2f,
	0x4d, 0x4c, 0x4f, 0x0d, 0x2e, 0x48, 0x4d, 0xce, 0x4c, 0xcb, 0x4c, 0xd6, 0x4f, 0x2c, 0x28, 0x28,
	0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x03, 0xab, 0x54, 0x0a, 0xe4, 0xe2, 0x77, 0x84, 0x09, 0xf9,
	0x83, 0x35, 0x09, 0x09, 0x71, 0xb1, 0x94, 0x16, 0xa7, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x04, 0x81, 0xd9, 0x42, 0x5a, 0x5c, 0x5c, 0x89, 0x05, 0x05, 0xc1, 0x60, 0x65, 0xc5, 0x12, 0x4c,
	0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x5c, 0x7a, 0x70, 0x9d, 0x41, 0x48, 0xb2, 0x4a, 0x96, 0x5c, 0x9c,
	0x70, 0x09, 0x90, 0x61, 0x79, 0x89, 0xb9, 0xa9, 0x60, 0xc3, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x09,
	0x2e, 0x76, 0x90, 0xc2, 0xcc, 0xfc, 0x3c, 0x09, 0x26, 0xb0, 0x30, 0x8c, 0xeb, 0xc4, 0x15, 0xc5,
	0xa1, 0x07, 0xf5, 0x42, 0x12, 0x1b, 0x98, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xd1,
	0xe8, 0xa5, 0xd3, 0x00, 0x00, 0x00,
}
