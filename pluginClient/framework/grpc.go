package framework

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient/pb"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of FrameworkVersions that talks over RPC.
type GRPCClient struct{ Client pb.FrameworkServiceClient }

//Detect will detect the usage of framework
func (m *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	ret, err := m.Client.Detect(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return ret, err
}

//IsFrameworkUsed returns true if that framework used
func (m *GRPCClient) IsFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	ret, err := m.Client.IsFrameworkUsed(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return ret, err
}

//PercentOfFrameworkUsed returns percentage of framework used
func (m *GRPCClient) PercentOfFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	ret, err := m.Client.PercentOfFrameworkUsed(context.Background(), &pb.ServiceInput{})
	return ret, err
}

//GRPCServer  is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl interfaces.FrameworkVersions
}

//Detect will detect the usage of framework
func (m *GRPCServer) Detect(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	v, err := m.Impl.Detect(req)
	return &pb.ServiceOutputBool{Value: v.Value, Error: v.Error}, err
}

//IsFrameworkUsed returns true if that framework used
func (m *GRPCServer) IsFrameworkUsed(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	v, err := m.Impl.IsFrameworkUsed(req)
	return &pb.ServiceOutputBool{Value: v.Value, Error: v.Error}, err
}

//PercentOfFrameworkUsed returns percentage of framework used
func (m *GRPCServer) PercentOfFrameworkUsed(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	v, err := m.Impl.PercentOfFrameworkUsed(req)
	return &pb.ServiceOutputFloat{Value: v.Value, Error: v.Error}, err
}
