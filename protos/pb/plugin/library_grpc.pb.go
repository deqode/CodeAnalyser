// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package plugin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LibraryClient is the client API for Library service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryClient interface {
	Detect(ctx context.Context, in *Input, opts ...grpc.CallOption) (*LibraryType, error)
	IsUsed(ctx context.Context, in *Input, opts ...grpc.CallOption) (*BoolOutput, error)
	PercentOfUsed(ctx context.Context, in *Input, opts ...grpc.CallOption) (*FloatOutput, error)
}

type libraryClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryClient(cc grpc.ClientConnInterface) LibraryClient {
	return &libraryClient{cc}
}

func (c *libraryClient) Detect(ctx context.Context, in *Input, opts ...grpc.CallOption) (*LibraryType, error) {
	out := new(LibraryType)
	err := c.cc.Invoke(ctx, "/proto.Library/Detect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) IsUsed(ctx context.Context, in *Input, opts ...grpc.CallOption) (*BoolOutput, error) {
	out := new(BoolOutput)
	err := c.cc.Invoke(ctx, "/proto.Library/IsUsed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) PercentOfUsed(ctx context.Context, in *Input, opts ...grpc.CallOption) (*FloatOutput, error) {
	out := new(FloatOutput)
	err := c.cc.Invoke(ctx, "/proto.Library/PercentOfUsed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServer is the server API for Library service.
// All implementations should embed UnimplementedLibraryServer
// for forward compatibility
type LibraryServer interface {
	Detect(context.Context, *Input) (*LibraryType, error)
	IsUsed(context.Context, *Input) (*BoolOutput, error)
	PercentOfUsed(context.Context, *Input) (*FloatOutput, error)
}

// UnimplementedLibraryServer should be embedded to have forward compatible implementations.
type UnimplementedLibraryServer struct {
}

func (UnimplementedLibraryServer) Detect(context.Context, *Input) (*LibraryType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detect not implemented")
}
func (UnimplementedLibraryServer) IsUsed(context.Context, *Input) (*BoolOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUsed not implemented")
}
func (UnimplementedLibraryServer) PercentOfUsed(context.Context, *Input) (*FloatOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PercentOfUsed not implemented")
}

// UnsafeLibraryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServer will
// result in compilation errors.
type UnsafeLibraryServer interface {
	mustEmbedUnimplementedLibraryServer()
}

func RegisterLibraryServer(s grpc.ServiceRegistrar, srv LibraryServer) {
	s.RegisterService(&Library_ServiceDesc, srv)
}

func _Library_Detect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).Detect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Library/Detect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).Detect(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_IsUsed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).IsUsed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Library/IsUsed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).IsUsed(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_PercentOfUsed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).PercentOfUsed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Library/PercentOfUsed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).PercentOfUsed(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

// Library_ServiceDesc is the grpc.ServiceDesc for Library service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Library_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Library",
	HandlerType: (*LibraryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Detect",
			Handler:    _Library_Detect_Handler,
		},
		{
			MethodName: "IsUsed",
			Handler:    _Library_IsUsed_Handler,
		},
		{
			MethodName: "PercentOfUsed",
			Handler:    _Library_PercentOfUsed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/plugin/library.proto",
}
