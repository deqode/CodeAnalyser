package preDetectCommand

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client plugin.PreDetectCommandClient
}

//RunPreDetect method call on client side over gRPC
func (g *GRPCClient) RunPreDetect(input *pb.Input) (*pb.EmptyOutput, error) {
	res, err := g.Client.RunPreDetect(context.Background(), input)
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl interfaces.PreDetectCommands
}

//RunPreDetect will execute all required commands (e.g. pruning of unused dependency) before all others detection
func (g *GRPCServer) RunPreDetect(ctx context.Context, input *pb.Input) (*pb.EmptyOutput, error) {
	res, err := g.Impl.RunPreDetect(input)
	return res, err
}
