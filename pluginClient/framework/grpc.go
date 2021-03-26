package framework

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of FrameworkVersions that talks over RPC.
type GRPCClient struct{ Client pb.FrameworkServiceClient }

func (m *GRPCClient) GetVersionName() (string, error) {
	ret, err := m.Client.GetVersionName(context.Background(), &pb.ServiceEmpty{
	})
	return ret.Value, err
}
func (m *GRPCClient) GetSemver() (string, error) {
	ret, err := m.Client.GetSemver(context.Background(), &pb.ServiceEmpty{
	})
	return ret.Value, err
}
func (m *GRPCClient) GetFrameworkName() (string, error) {
	ret, err := m.Client.GetFrameworkName(context.Background(), &pb.ServiceEmpty{
	})
	return ret.Value, err
}

func (m *GRPCClient) Detect(runtimeVersion, root string) (bool, error) {
	ret, err := m.Client.Detect(context.Background(), &pb.ServiceInput{
		RuntimeVersion: runtimeVersion,
		Root:           root,
	})
	return ret.Value, err
}
func (m *GRPCClient) IsFrameworkFound(runtimeVersion, root string) (bool, error) {
	ret, err := m.Client.IsFrameworkFound(context.Background(), &pb.ServiceInput{
		RuntimeVersion: runtimeVersion,
		Root:           runtimeVersion,
	})
	return ret.Value, err
}
func (m *GRPCClient) IsFrameworkUsed(runtimeVersion, root string) (bool, error) {
	ret, err := m.Client.IsFrameworkUsed(context.Background(), &pb.ServiceInput{
		RuntimeVersion: runtimeVersion,
		Root:           root,
	})
	return ret.Value, err
}
func (m *GRPCClient) PercentOfFrameworkUsed(runtimeVersion, root string) (int64, error) {
	ret, err := m.Client.PercentOfFrameworkUsed(context.Background(), &pb.ServiceInput{})
	return ret.Value, err
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl interfaces.FrameworkVersions
}

func (m *GRPCServer) GetVersionName(context.Context, *pb.ServiceEmpty) (*pb.ServiceOutputString, error) {
	v, err := m.Impl.GetVersionName()
	return &pb.ServiceOutputString{Value: v}, err
}

func (m *GRPCServer) GetSemver(context.Context, *pb.ServiceEmpty) (*pb.ServiceOutputString, error) {
	v, err := m.Impl.GetSemver()
	return &pb.ServiceOutputString{Value: v}, err
}

func (m *GRPCServer) Detect(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	v, err := m.Impl.Detect(req.RuntimeVersion, req.Root)
	return &pb.ServiceOutputBool{Value: v}, err
}

func (m *GRPCServer) GetFrameworkName(context.Context, *pb.ServiceEmpty) (*pb.ServiceOutputString, error) {
	v, err := m.Impl.GetFrameworkName()
	return &pb.ServiceOutputString{Value: v}, err
}
func (m *GRPCServer) IsFrameworkUsed(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	v, err := m.Impl.IsFrameworkUsed(req.RuntimeVersion, req.Root)
	return &pb.ServiceOutputBool{Value: v}, err
}
func (m *GRPCServer) IsFrameworkFound(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	v, err := m.Impl.IsFrameworkFound(req.RuntimeVersion, req.Root)
	return &pb.ServiceOutputBool{Value: v}, err
}
func (m *GRPCServer) PercentOfFrameworkUsed(ctx context.Context, req *pb.ServiceInput) (*pb.ServiceOutputInt, error) {
	v, err := m.Impl.PercentOfFrameworkUsed(req.RuntimeVersion, req.Root)
	return &pb.ServiceOutputInt{Value: v}, err
}
