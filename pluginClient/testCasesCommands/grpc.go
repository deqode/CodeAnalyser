package testCasesCommands

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	"code-analyser/protos/pb/plugin"
	"context"
)

type GRPCClient struct {
	Client plugin.TestCaseCommandsClient
}

func (g *GRPCClient) Detect(input *pb.Input) (*languageSpecific.TestCasesCommand, error) {
	res, err := g.Client.Detect(context.Background(), input)
	return res, err
}

type GRPCServer struct {
	Impl interfaces.TestCasesRunCommands
}

func (m *GRPCServer) Detect(ctx context.Context, input *pb.Input) (*languageSpecific.TestCasesCommand, error) {
	res, err := m.Impl.Detect(input)
	return res, err
}
