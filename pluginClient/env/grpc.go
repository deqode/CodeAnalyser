package env

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	"code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

//GRPCClient struct for grpc client call
type GRPCClient struct {
	Client plugin.EnvClient
}

//Detect method implemented from Env interface (interface implemented by us)
func (m *GRPCClient) Detect(input *pb.Input) (*languageSpecific.Envs, error) {
	res, err := m.Client.Detect(context.Background(), &pb.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return res, err
}

//GRPCServer struct for plugin call
type GRPCServer struct {
	Impl interfaces.Env
}

//Detect implemented function from interface EnvServiceServer(auto generated from proto)
func (m *GRPCServer) Detect(ctx context.Context, input *pb.Input) (*languageSpecific.Envs, error) {
	res, err := m.Impl.Detect(input)
	return &languageSpecific.Envs{
		EnvKeyValues: res.EnvKeyValues,
		Keys:         res.Keys,
		Error:        res.Error,
	}, err
}
