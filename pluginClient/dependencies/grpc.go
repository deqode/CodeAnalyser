package dependencies

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/protos/pb/helpers"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of Dependencies that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client pb.DependenciesClient
}

//GetDependencies method call on client side over gRPC
func (m *GRPCClient) GetDependencies(inputString *helpers.Input) (*helpers.StringMapOutput, error) {
	res, err := m.Client.GetDependencies(context.Background(), inputString)
	if res != nil {
		return res, err
	}
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl interfaces.Dependencies
}

//GetDependencies will return all the dependencies used
func (m *GRPCServer) GetDependencies(ctx context.Context, input *helpers.Input) (*helpers.StringMapOutput, error) {
	res, err := m.Impl.GetDependencies(input)
	return res, err
}
