package framework

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb/proto"
	"context"
)

// GRPCClient is an implementation of FrameworkVersions that talks over RPC.
type GRPCClient struct {
	Client proto.FrameworkServiceClient
}

func (m *GRPCClient) GetVersionName() (string, error) {
	ret, err := m.Client.GetVersionName(context.Background(), &proto.FrameworkServiceEmpty{})
	if err != nil {
		return "", err
	}
	return ret.Value, err
}
func (m *GRPCClient) GetSemver() (string, error) {
	ret, err := m.Client.GetSemver(context.Background(), &proto.FrameworkServiceEmpty{})
	if err != nil {
		return "", err
	}
	return ret.Value, err
}
func (m *GRPCClient) Detect(runtimeVersion, root string) (bool, error) {
	ret, err := m.Client.Detect(context.Background(), &proto.FrameworkServiceInput{
		Root:           root,
		RuntimeVersion: runtimeVersion,
	})
	if err != nil {
		return false, err
	}
	return ret.Value, err
}
func (m *GRPCClient) IsFrameworkFound(runtimeVersion, root string) (bool, error) {
	ret, err := m.Client.IsFrameworkFound(context.Background(), &proto.FrameworkServiceInput{
		Root:           root,
		RuntimeVersion: runtimeVersion,
	})
	if err != nil {
		return false, err
	}
	return ret.Value, err
}
func (m *GRPCClient) IsFrameworkUsed(runtimeVersion, root string) (bool, error) {
	ret, err := m.Client.IsFrameworkUsed(context.Background(), &proto.FrameworkServiceInput{
		Root:           root,
		RuntimeVersion: runtimeVersion,
	})
	if err != nil {
		return false, err
	}
	return ret.Value, err
}
func (m *GRPCClient) PercentOfFrameworkUsed(runtimeVersion, root string) (int64, error) {
	ret, err := m.Client.PercentOfFrameworkUsed(context.Background(), &proto.FrameworkServiceInput{
		Root:           root,
		RuntimeVersion: runtimeVersion,
	})
	if err != nil {
		return 0, err
	}
	return ret.Value, err
}
func (m *GRPCClient) GetFrameworkName() (string, error) {
	ret, err := m.Client.GetFrameworkName(context.Background(), &proto.FrameworkServiceEmpty{})
	if err != nil {
		return "", err
	}
	return ret.Value, err
}

type GRPCServer struct {
	// This is the real implementation
	Impl interfaces.FrameworkVersions
}



func (m *GRPCServer) GetVersionName(context.Context, *proto.FrameworkServiceEmpty) (*proto.FrameworkServiceOutputString, error) {
	v, err := m.Impl.GetVersionName()
	return &proto.FrameworkServiceOutputString{Value: v}, err
}
func (m *GRPCServer) GetSemver(context.Context, *proto.FrameworkServiceEmpty) (*proto.FrameworkServiceOutputString, error) {
	v, err := m.Impl.GetSemver()
	return &proto.FrameworkServiceOutputString{Value: v}, err
}
func (m *GRPCServer) Detect(ctx context.Context, req *proto.FrameworkServiceInput) (*proto.FrameworkServiceOutputBool, error) {
	v, err := m.Impl.Detect(req.RuntimeVersion, req.Root)
	return &proto.FrameworkServiceOutputBool{Value: v}, err
}
func (m *GRPCServer) IsFrameworkFound(ctx context.Context, req *proto.FrameworkServiceInput) (*proto.FrameworkServiceOutputBool, error) {
	v, err := m.Impl.IsFrameworkFound(req.RuntimeVersion, req.Root)
	return &proto.FrameworkServiceOutputBool{Value: v}, err
}
func (m *GRPCServer) IsFrameworkUsed(ctx context.Context, req *proto.FrameworkServiceInput) (*proto.FrameworkServiceOutputBool, error) {
	v, err := m.Impl.IsFrameworkUsed(req.RuntimeVersion, req.Root)
	return &proto.FrameworkServiceOutputBool{Value: v}, err
}
func (m *GRPCServer) PercentOfFrameworkUsed(ctx context.Context, req *proto.FrameworkServiceInput) (*proto.FrameworkServiceOutputInt, error) {
	v, err := m.Impl.PercentOfFrameworkUsed(req.RuntimeVersion, req.Root)
	return &proto.FrameworkServiceOutputInt{Value: v}, err
}
func (m *GRPCServer) GetFrameworkName(context.Context, *proto.FrameworkServiceEmpty) (*proto.FrameworkServiceOutputString, error) {
	v, err := m.Impl.GetSemver()
	return &proto.FrameworkServiceOutputString{Value: v}, err
}
