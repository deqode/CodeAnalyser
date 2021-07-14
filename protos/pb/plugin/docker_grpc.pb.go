// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package plugin

import (
	helpers "code-analyser/protos/pb/helpers"
	globalFiles "code-analyser/protos/pb/output/globalFiles"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DockerFileClient is the client API for DockerFile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DockerFileClient interface {
	DetectDockerFiles(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*globalFiles.DockerFile, error)
	DetectDockerComposeFiles(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*globalFiles.DockerCompose, error)
}

type dockerFileClient struct {
	cc grpc.ClientConnInterface
}

func NewDockerFileClient(cc grpc.ClientConnInterface) DockerFileClient {
	return &dockerFileClient{cc}
}

func (c *dockerFileClient) DetectDockerFiles(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*globalFiles.DockerFile, error) {
	out := new(globalFiles.DockerFile)
	err := c.cc.Invoke(ctx, "/proto.DockerFile/DetectDockerFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerFileClient) DetectDockerComposeFiles(ctx context.Context, in *helpers.StringInput, opts ...grpc.CallOption) (*globalFiles.DockerCompose, error) {
	out := new(globalFiles.DockerCompose)
	err := c.cc.Invoke(ctx, "/proto.DockerFile/DetectDockerComposeFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockerFileServer is the server API for DockerFile service.
// All implementations should embed UnimplementedDockerFileServer
// for forward compatibility
type DockerFileServer interface {
	DetectDockerFiles(context.Context, *helpers.StringInput) (*globalFiles.DockerFile, error)
	DetectDockerComposeFiles(context.Context, *helpers.StringInput) (*globalFiles.DockerCompose, error)
}

// UnimplementedDockerFileServer should be embedded to have forward compatible implementations.
type UnimplementedDockerFileServer struct {
}

func (UnimplementedDockerFileServer) DetectDockerFiles(context.Context, *helpers.StringInput) (*globalFiles.DockerFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectDockerFiles not implemented")
}
func (UnimplementedDockerFileServer) DetectDockerComposeFiles(context.Context, *helpers.StringInput) (*globalFiles.DockerCompose, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectDockerComposeFiles not implemented")
}

// UnsafeDockerFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DockerFileServer will
// result in compilation errors.
type UnsafeDockerFileServer interface {
	mustEmbedUnimplementedDockerFileServer()
}

func RegisterDockerFileServer(s grpc.ServiceRegistrar, srv DockerFileServer) {
	s.RegisterService(&DockerFile_ServiceDesc, srv)
}

func _DockerFile_DetectDockerFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helpers.StringInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerFileServer).DetectDockerFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DockerFile/DetectDockerFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerFileServer).DetectDockerFiles(ctx, req.(*helpers.StringInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerFile_DetectDockerComposeFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helpers.StringInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerFileServer).DetectDockerComposeFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DockerFile/DetectDockerComposeFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerFileServer).DetectDockerComposeFiles(ctx, req.(*helpers.StringInput))
	}
	return interceptor(ctx, in, info, handler)
}

// DockerFile_ServiceDesc is the grpc.ServiceDesc for DockerFile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DockerFile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DockerFile",
	HandlerType: (*DockerFileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DetectDockerFiles",
			Handler:    _DockerFile_DetectDockerFiles_Handler,
		},
		{
			MethodName: "DetectDockerComposeFiles",
			Handler:    _DockerFile_DetectDockerComposeFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/plugin/docker.proto",
}
