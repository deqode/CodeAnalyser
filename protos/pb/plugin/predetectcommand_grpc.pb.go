// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package plugin

import (
	helpers "code-analyser/protos/pb/helpers"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PreDetectCommandClient is the client API for PreDetectCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PreDetectCommandClient interface {
	RunPreDetect(ctx context.Context, in *helpers.Input, opts ...grpc.CallOption) (*helpers.EmptyOutput, error)
}

type preDetectCommandClient struct {
	cc grpc.ClientConnInterface
}

func NewPreDetectCommandClient(cc grpc.ClientConnInterface) PreDetectCommandClient {
	return &preDetectCommandClient{cc}
}

func (c *preDetectCommandClient) RunPreDetect(ctx context.Context, in *helpers.Input, opts ...grpc.CallOption) (*helpers.EmptyOutput, error) {
	out := new(helpers.EmptyOutput)
	err := c.cc.Invoke(ctx, "/proto.PreDetectCommand/RunPreDetect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PreDetectCommandServer is the server API for PreDetectCommand service.
// All implementations should embed UnimplementedPreDetectCommandServer
// for forward compatibility
type PreDetectCommandServer interface {
	RunPreDetect(context.Context, *helpers.Input) (*helpers.EmptyOutput, error)
}

// UnimplementedPreDetectCommandServer should be embedded to have forward compatible implementations.
type UnimplementedPreDetectCommandServer struct {
}

func (UnimplementedPreDetectCommandServer) RunPreDetect(context.Context, *helpers.Input) (*helpers.EmptyOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunPreDetect not implemented")
}

// UnsafePreDetectCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PreDetectCommandServer will
// result in compilation errors.
type UnsafePreDetectCommandServer interface {
	mustEmbedUnimplementedPreDetectCommandServer()
}

func RegisterPreDetectCommandServer(s grpc.ServiceRegistrar, srv PreDetectCommandServer) {
	s.RegisterService(&PreDetectCommand_ServiceDesc, srv)
}

func _PreDetectCommand_RunPreDetect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helpers.Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreDetectCommandServer).RunPreDetect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PreDetectCommand/RunPreDetect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreDetectCommandServer).RunPreDetect(ctx, req.(*helpers.Input))
	}
	return interceptor(ctx, in, info, handler)
}

// PreDetectCommand_ServiceDesc is the grpc.ServiceDesc for PreDetectCommand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PreDetectCommand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PreDetectCommand",
	HandlerType: (*PreDetectCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunPreDetect",
			Handler:    _PreDetectCommand_RunPreDetect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/plugin/predetectcommand.proto",
}
