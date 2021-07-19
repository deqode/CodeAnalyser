package commands

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of BuildDirectory that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client pb.CommandsClient
}

//DetectBuildCommands method call on client side over gRPC
func (g *GRPCClient) DetectBuildCommands(projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Client.DetectBuildCommands(context.Background(), projectRootPath)
	return res, err
}

//DetectStartUpCommands method call on client side over gRPC
func (g *GRPCClient) DetectStartUpCommands(projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Client.DetectStartUpCommands(context.Background(), projectRootPath)
	return res, err
}

//DetectSeedCommands method call on client side over gRPC
func (g *GRPCClient) DetectSeedCommands(projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Client.DetectSeedCommands(context.Background(), projectRootPath)
	return res, err
}

//DetectMigrationCommands method call on client side over gRPC
func (g *GRPCClient) DetectMigrationCommands(projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Client.DetectMigrationCommands(context.Background(), projectRootPath)
	return res, err
}

//DetectAdHocScripts method call on client side over gRPC
func (g *GRPCClient) DetectAdHocScripts(projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Client.DetectAdHocScripts(context.Background(), projectRootPath)
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl interfaces.Commands
}

//DetectBuildCommands will detect build commands in directory
func (g *GRPCServer) DetectBuildCommands(ctx context.Context, projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Impl.DetectBuildCommands(projectRootPath)
	return res, err
}

// DetectStartUpCommands will detect startup commands in directory
func (g *GRPCServer) DetectStartUpCommands(ctx context.Context, projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Impl.DetectStartUpCommands(projectRootPath)
	return res, err
}

//DetectSeedCommands will detect seed commands in directory
func (g *GRPCServer) DetectSeedCommands(ctx context.Context, projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Impl.DetectSeedCommands(projectRootPath)
	return res, err
}

//DetectMigrationCommands will detect migration commands in directory
func (g *GRPCServer) DetectMigrationCommands(ctx context.Context, projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Impl.DetectMigrationCommands(projectRootPath)
	return res, err
}

//DetectAdHocScripts will detect adhoc scripts in directory
func (g *GRPCServer) DetectAdHocScripts(ctx context.Context, projectRootPath *helpers.StringInput) (*languageSpecific.CommandOutput, error) {
	res, err := g.Impl.DetectAdHocScripts(projectRootPath)
	return res, err
}
