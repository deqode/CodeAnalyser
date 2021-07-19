package env

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	"code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client plugin.EnvClient
}

//Detect method call on client side over gRPC
func (m *GRPCClient) Detect(input *pb.Input) (*languageSpecific.Envs, error) {
	res, err := m.Client.Detect(context.Background(), &pb.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl interfaces.Env
}

//Detect will detect and parse environment variables
func (m *GRPCServer) Detect(ctx context.Context, input *pb.Input) (*languageSpecific.Envs, error) {
	res, err := m.Impl.Detect(input)
	return &languageSpecific.Envs{
		EnvKeyValues: res.EnvKeyValues,
		Keys:         res.Keys,
		Error:        res.Error,
	}, err
}
