package framework

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of FrameworkVersions that talks over RPC.
type GRPCClient struct{ Client pb.FrameworkServiceClient }

func (m *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	ret, err := m.Client.Detect(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return ret, err
}

func (m *GRPCClient) IsFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	ret, err := m.Client.IsFrameworkUsed(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return ret, err
}

func (m *GRPCClient) PercentOfFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	ret, err := m.Client.PercentOfFrameworkUsed(context.Background(), &pb.ServiceInput{})
	return ret, err
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl interfaces.FrameworkVersions
}

func (m *GRPCServer) Detect(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	v, err := m.Impl.Detect(req)
	return &pb.ServiceOutputBool{Value: v.Value, Error: v.Error}, err
}

func (m *GRPCServer) IsFrameworkUsed(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	v, err := m.Impl.IsFrameworkUsed(req)
	return &pb.ServiceOutputBool{Value: v.Value, Error: v.Error}, err
}

func (m *GRPCServer) PercentOfFrameworkUsed(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	v, err := m.Impl.PercentOfFrameworkUsed(req)
	return &pb.ServiceOutputFloat{Value: v.Value, Error: v.Error}, err
}
