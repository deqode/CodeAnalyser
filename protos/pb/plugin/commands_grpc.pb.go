// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package plugin

import (
	helpers "code-analyser/protos/pb/helpers"
	languageSpecific "code-analyser/protos/pb/output/languageSpecific"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommandsClient is the client API for Commands service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandsClient interface {
	DetectBuildCommands(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*languageSpecific.CommandOutput, error)
	DetectStartUpCommands(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*languageSpecific.CommandOutput, error)
	DetectSeedCommands(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*languageSpecific.CommandOutput, error)
	DetectMigrationCommands(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*languageSpecific.CommandOutput, error)
	DetectAdHocScripts(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*languageSpecific.CommandOutput, error)
}

type commandsClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandsClient(cc grpc.ClientConnInterface) CommandsClient {
	return &commandsClient{cc}
}

func (c *commandsClient) DetectBuildCommands(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*languageSpecific.CommandOutput, error) {
	out := new(languageSpecific.CommandOutput)
	err := c.cc.Invoke(ctx, "/proto.Commands/DetectBuildCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsClient) DetectStartUpCommands(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*languageSpecific.CommandOutput, error) {
	out := new(languageSpecific.CommandOutput)
	err := c.cc.Invoke(ctx, "/proto.Commands/DetectStartUpCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsClient) DetectSeedCommands(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*languageSpecific.CommandOutput, error) {
	out := new(languageSpecific.CommandOutput)
	err := c.cc.Invoke(ctx, "/proto.Commands/DetectSeedCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsClient) DetectMigrationCommands(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*languageSpecific.CommandOutput, error) {
	out := new(languageSpecific.CommandOutput)
	err := c.cc.Invoke(ctx, "/proto.Commands/DetectMigrationCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsClient) DetectAdHocScripts(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*languageSpecific.CommandOutput, error) {
	out := new(languageSpecific.CommandOutput)
	err := c.cc.Invoke(ctx, "/proto.Commands/DetectAdHocScripts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandsServer is the server API for Commands service.
// All implementations should embed UnimplementedCommandsServer
// for forward compatibility
type CommandsServer interface {
	DetectBuildCommands(context.Context, *helpers.StringInput) (*languageSpecific.CommandOutput, error)
	DetectStartUpCommands(context.Context, *helpers.StringInput) (*languageSpecific.CommandOutput, error)
	DetectSeedCommands(context.Context, *helpers.StringInput) (*languageSpecific.CommandOutput, error)
	DetectMigrationCommands(context.Context, *helpers.StringInput) (*languageSpecific.CommandOutput, error)
	DetectAdHocScripts(context.Context, *helpers.StringInput) (*languageSpecific.CommandOutput, error)
}

// UnimplementedCommandsServer should be embedded to have forward compatible implementations.
type UnimplementedCommandsServer struct {
}

func (UnimplementedCommandsServer) DetectBuildCommands(context.Context, *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectBuildCommands not implemented")
}
func (UnimplementedCommandsServer) DetectStartUpCommands(context.Context, *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectStartUpCommands not implemented")
}
func (UnimplementedCommandsServer) DetectSeedCommands(context.Context, *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectSeedCommands not implemented")
}
func (UnimplementedCommandsServer) DetectMigrationCommands(context.Context, *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectMigrationCommands not implemented")
}
func (UnimplementedCommandsServer) DetectAdHocScripts(context.Context, *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectAdHocScripts not implemented")
}

// UnsafeCommandsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandsServer will
// result in compilation errors.
type UnsafeCommandsServer interface {
	mustEmbedUnimplementedCommandsServer()
}

func RegisterCommandsServer(s grpc.ServiceRegistrar, srv CommandsServer) {
	s.RegisterService(&Commands_ServiceDesc, srv)
}

func _Commands_DetectBuildCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helpers.StringInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).DetectBuildCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Commands/DetectBuildCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).DetectBuildCommands(ctx, req.(*helpers.StringInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commands_DetectStartUpCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helpers.StringInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).DetectStartUpCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Commands/DetectStartUpCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).DetectStartUpCommands(ctx, req.(*helpers.StringInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commands_DetectSeedCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helpers.StringInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).DetectSeedCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Commands/DetectSeedCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).DetectSeedCommands(ctx, req.(*helpers.StringInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commands_DetectMigrationCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helpers.StringInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).DetectMigrationCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Commands/DetectMigrationCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).DetectMigrationCommands(ctx, req.(*helpers.StringInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commands_DetectAdHocScripts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helpers.StringInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).DetectAdHocScripts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Commands/DetectAdHocScripts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).DetectAdHocScripts(ctx, req.(*helpers.StringInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Commands_ServiceDesc is the grpc.ServiceDesc for Commands service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Commands_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Commands",
	HandlerType: (*CommandsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DetectBuildCommands",
			Handler:    _Commands_DetectBuildCommands_Handler,
		},
		{
			MethodName: "DetectStartUpCommands",
			Handler:    _Commands_DetectStartUpCommands_Handler,
		},
		{
			MethodName: "DetectSeedCommands",
			Handler:    _Commands_DetectSeedCommands_Handler,
		},
		{
			MethodName: "DetectMigrationCommands",
			Handler:    _Commands_DetectMigrationCommands_Handler,
		},
		{
			MethodName: "DetectAdHocScripts",
			Handler:    _Commands_DetectAdHocScripts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/plugin/commands.proto",
}
