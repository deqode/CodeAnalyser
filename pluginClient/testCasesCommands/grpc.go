package testCasesCommands

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/plugin"
	"context"
)

type GRPCClient struct {
	Client pb.TestCaseCommandsClient
}

func (g *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputTestCommand, error) {
	res, err := g.Client.Detect(context.Background(), input)
	return res, err
}

type GRPCServer struct {
	Impl interfaces.TestCasesRunCommands
}

func (m *GRPCServer) Detect(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputTestCommand, error) {
	res, err := m.Impl.Detect(input)
	return res, err
}
