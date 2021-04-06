package env

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client pb.EnvServiceClient
}

func (m *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputEnv, error) {
	res, err := m.Client.Detect(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}

func (m *GRPCClient) IsUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := m.Client.IsEnvUsed(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}

type GRPCServer struct {
	Impl interfaces.Env
}

func (m *GRPCServer) Detect(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputEnv, error) {
	res, err := m.Impl.Detect(input)
	return &pb.ServiceOutputEnv{
		EnvKeyValues: res.EnvKeyValues,
		Keys:         res.Keys,
		Error:        res.Error,
	}, err
}

func (m *GRPCServer) IsEnvUsed(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := m.Impl.IsUsed(input)
	return &pb.ServiceOutputBool{
		Value: res.Value,
		Error: res.Error,
	}, err
}
