package dependencies

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client pb.DependenciesServiceClient
}

type GRPCServer struct {
	Impl interfaces.Dependencies
}

func (m *GRPCClient) GetDependencies(inputString *pb.ServiceInput) (*pb.ServiceOutputStringMap, error) {
	res, err := m.Client.GetDependencies(context.Background(), inputString)
	if res != nil {
		return res, err
	}
	return res, err
}

func (m *GRPCServer) GetDependencies(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputStringMap, error) {
	res, err := m.Impl.GetDependencies(input)
	return res, err
}
