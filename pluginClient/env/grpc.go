package env

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient/pb"
	"golang.org/x/net/context"
)
//GRPCClient struct for grpc client call
type GRPCClient struct {
	Client pb.EnvServiceClient
}
//Detect method implemented from Env interface (interface implemented by us)
func (m *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputEnv, error) {
	res, err := m.Client.Detect(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}
//IsUsed method implemented from Env interface (interface implemented by us)
func (m *GRPCClient) IsUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := m.Client.IsEnvUsed(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}

//GRPCServer struct for plugin call
type GRPCServer struct {
	Impl interfaces.Env
}

//Detect implemented function from interface EnvServiceServer(auto generated from proto)
func (m *GRPCServer) Detect(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputEnv, error) {
	res, err := m.Impl.Detect(input)
	return &pb.ServiceOutputEnv{
		EnvKeyValues: res.EnvKeyValues,
		Keys:         res.Keys,
		Error:        res.Error,
	}, err
}

//IsEnvUsed implemented function from interface EnvServiceServer(auto generated from proto)
func (m *GRPCServer) IsEnvUsed(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := m.Impl.IsUsed(input)
	return &pb.ServiceOutputBool{
		Value: res.Value,
		Error: res.Error,
	}, err
}
