package commands

import (
	"code-analyser/GlobalFiles"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client pb.CommandsServiceClient
}

func (g *GRPCClient) DetectBuildCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceOutputDetectBuildCommands, error) {
	res, err := g.Client.DetectBuildCommands(context.Background(), inputString)
	return res, err
}

func (g *GRPCClient) DetectStartUpCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceOutputStartUpCommands, error) {
	res, err := g.Client.DetectStartUpCommands(context.Background(), inputString)
	return res, err
}

func (g *GRPCClient) DetectSeedCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceOutputSeedCommands, error) {
	res, err := g.Client.DetectSeedCommands(context.Background(), inputString)
	return res, err
}

func (g *GRPCClient) DetectMigrationCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceMigrationCommands, error) {
	res, err := g.Client.DetectMigrationCommands(context.Background(), inputString)
	return res, err
}

func (g *GRPCClient) DetectAdHocScripts(inputString *pb.ServiceCommandsInput) (*pb.ServiceAdHocScripts, error) {
	res, err := g.Client.DetectAdHocScripts(context.Background(), inputString)
	return res, err
}

type GRPCServer struct {
	Impl GlobalFiles.Commands
}

func (g *GRPCServer) DetectBuildCommands(ctx context.Context, input *pb.ServiceCommandsInput) (*pb.ServiceOutputDetectBuildCommands, error) {
	res, err := g.Impl.DetectBuildCommands(input)
	return res, err
}

func (g *GRPCServer) DetectStartUpCommands(ctx context.Context, input *pb.ServiceCommandsInput) (*pb.ServiceOutputStartUpCommands, error) {
	res, err := g.Impl.DetectStartUpCommands(input)
	return res, err
}

func (g *GRPCServer) DetectSeedCommands(ctx context.Context, input *pb.ServiceCommandsInput) (*pb.ServiceOutputSeedCommands, error) {
	res, err := g.Impl.DetectSeedCommands(input)
	return res, err
}

func (g *GRPCServer) DetectMigrationCommands(ctx context.Context, input *pb.ServiceCommandsInput) (*pb.ServiceMigrationCommands, error) {
	res, err := g.Impl.DetectMigrationCommands(input)
	return res, err
}

func (g *GRPCServer) DetectAdHocScripts(ctx context.Context, input *pb.ServiceCommandsInput) (*pb.ServiceAdHocScripts, error) {
	res, err := g.Impl.DetectAdHocScripts(input)
	return res, err
}
