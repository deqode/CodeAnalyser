package dependencies

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of Framework that talks over RPC.
type GRPCClient struct {
	Client pb.DependenciesServiceClient
}

//GRPCServer  is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	Impl interfaces.Dependencies
}

//GetDependencies will return all the dependencies used
func (m *GRPCClient) GetDependencies(inputString *pb.ServiceInput) (*pb.ServiceOutputStringMap, error) {
	res, err := m.Client.GetDependencies(context.Background(), inputString)
	if res != nil {
		return res, err
	}
	return res, err
}

//GetDependencies will return all the dependencies used
func (m *GRPCServer) GetDependencies(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputStringMap, error) {
	res, err := m.Impl.GetDependencies(input)
	return res, err
}
