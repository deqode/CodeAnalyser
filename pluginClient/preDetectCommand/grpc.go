package preDetectCommand

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client plugin.PreDetectCommandClient
}

func (g *GRPCClient) RunPreDetect(input *pb.Input) (*pb.EmptyOutput, error) {
	res, err := g.Client.RunPreDetect(context.Background(), input)
	return res, err
}

type GRPCServer struct {
	Impl interfaces.PreDetectCommands
}

func (g *GRPCServer) RunPreDetect(ctx context.Context, input *pb.Input) (*pb.EmptyOutput, error) {
	res, err := g.Impl.RunPreDetect(input)
	return res, err
}
