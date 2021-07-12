package commands

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client pb.CommandsClient
}

func (g *GRPCClient) DetectBuildCommands(projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Client.DetectBuildCommands(context.Background(), projectRootPath)
	return res, err
}

func (g *GRPCClient) DetectStartUpCommands(projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Client.DetectStartUpCommands(context.Background(), projectRootPath)
	return res, err
}

func (g *GRPCClient) DetectSeedCommands(projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Client.DetectSeedCommands(context.Background(), projectRootPath)
	return res, err
}

func (g *GRPCClient) DetectMigrationCommands(projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Client.DetectMigrationCommands(context.Background(), projectRootPath)
	return res, err
}

func (g *GRPCClient) DetectAdHocScripts(projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Client.DetectAdHocScripts(context.Background(), projectRootPath)
	return res, err
}

type GRPCServer struct {
	Impl interfaces.Commands
}

func (g *GRPCServer) DetectBuildCommands(ctx context.Context, projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Impl.DetectBuildCommands(projectRootPath)
	return res, err
}

func (g *GRPCServer) DetectStartUpCommands(ctx context.Context, projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Impl.DetectStartUpCommands(projectRootPath)
	return res, err
}

func (g *GRPCServer) DetectSeedCommands(ctx context.Context, projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Impl.DetectSeedCommands(projectRootPath)
	return res, err
}

func (g *GRPCServer) DetectMigrationCommands(ctx context.Context, projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Impl.DetectMigrationCommands(projectRootPath)
	return res, err
}

func (g *GRPCServer) DetectAdHocScripts(ctx context.Context, projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Impl.DetectAdHocScripts(projectRootPath)
	return res, err
}
