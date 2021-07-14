package dependencies

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/protos/pb/helpers"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of Framework that talks over RPC.
type GRPCClient struct {
	Client pb.DependenciesClient
}

//GRPCServer  is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	Impl interfaces.Dependencies
}

//GetDependencies will return all the dependencies used
func (m *GRPCClient) GetDependencies(inputString *helpers.Input) (*helpers.StringMapOutput, error) {
	res, err := m.Client.GetDependencies(context.Background(), inputString)
	if res != nil {
		return res, err
	}
	return res, err
}

//GetDependencies will return all the dependencies used
func (m *GRPCServer) GetDependencies(ctx context.Context, input *helpers.Input) (*helpers.StringMapOutput, error) {
	res, err := m.Impl.GetDependencies(input)
	return res, err
}
