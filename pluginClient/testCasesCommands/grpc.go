package testCasesCommands

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	"code-analyser/protos/pb/plugin"
	"context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client plugin.TestCaseCommandsClient
}

//Detect method call on client side over gRPC
func (g *GRPCClient) Detect(input *pb.Input) (*languageSpecific.TestCasesCommand, error) {
	res, err := g.Client.Detect(context.Background(), input)
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl interfaces.TestCasesRunCommands
}

//Detect will detect test commands
func (m *GRPCServer) Detect(ctx context.Context, input *pb.Input) (*languageSpecific.TestCasesCommand, error) {
	res, err := m.Impl.Detect(input)
	return res, err
}
