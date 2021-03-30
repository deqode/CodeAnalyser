package framework

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of FrameworkVersions that talks over RPC.
type GRPCClient struct{ Client pb.FrameworkServiceClient }

func (m *GRPCClient) GetVersionName() (*pb.ServiceOutputString, error) {
	ret, err := m.Client.GetVersionName(context.Background(), &pb.ServiceEmpty{})
	return ret, err
}
func (m *GRPCClient) GetSemver() (*pb.ServiceOutputString, error) {
	ret, err := m.Client.GetSemver(context.Background(), &pb.ServiceEmpty{
	})

	return ret, err
}
func (m *GRPCClient) GetFrameworkName() (*pb.ServiceOutputString, error) {
	ret, err := m.Client.GetFrameworkName(context.Background(), &pb.ServiceEmpty{
	})
	return ret, err
}

func (m *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	ret, err := m.Client.Detect(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return ret, err
}
func (m *GRPCClient) IsFrameworkFound(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	ret, err := m.Client.IsFrameworkFound(context.Background(), &pb.ServiceInput{
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

func (m *GRPCServer) GetVersionName(context.Context, *pb.ServiceEmpty) (*pb.ServiceOutputString, error) {
	v, err := m.Impl.GetVersionName()
	return &pb.ServiceOutputString{Value: v.Value, Error: v.Error}, err
}

func (m *GRPCServer) GetSemver(context.Context, *pb.ServiceEmpty) (*pb.ServiceOutputString, error) {
	v, err := m.Impl.GetSemver()
	return &pb.ServiceOutputString{Value: v.Value, Error: v.Error}, err
}

func (m *GRPCServer) Detect(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	v, err := m.Impl.Detect(req)
	return &pb.ServiceOutputBool{Value: v.Value, Error: v.Error}, err
}

func (m *GRPCServer) GetFrameworkName(context.Context, *pb.ServiceEmpty) (*pb.ServiceOutputString, error) {
	v, err := m.Impl.GetFrameworkName()
	return &pb.ServiceOutputString{Value: v.Value, Error: v.Error}, err
}
func (m *GRPCServer) IsFrameworkUsed(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	v, err := m.Impl.IsFrameworkUsed(req)
	return &pb.ServiceOutputBool{Value: v.Value, Error: v.Error}, err
}
func (m *GRPCServer) IsFrameworkFound(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	v, err := m.Impl.IsFrameworkFound(req)
	return &pb.ServiceOutputBool{Value: v.Value, Error: v.Error}, err
}
func (m *GRPCServer) PercentOfFrameworkUsed(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	v, err := m.Impl.PercentOfFrameworkUsed(req)
	return &pb.ServiceOutputFloat{Value: v.Value, Error: v.Error}, err
}
