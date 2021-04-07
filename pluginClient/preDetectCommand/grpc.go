package preDetectCommand

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient/pb"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client pb.PreDetectCommandClient
}

func (g *GRPCClient) RunPreDetect(input *pb.ServiceInput) (*pb.ServiceEmptyOutput, error) {
	res, err := g.Client.RunPreDetect(context.Background(), input)
	return res, err
}

type GRPCServer struct {
	Impl interfaces.PreDetectCommands
}

func (g *GRPCServer) RunPreDetect(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceEmptyOutput, error) {
	res, err := g.Impl.RunPreDetect(input)
	return res, err
}
