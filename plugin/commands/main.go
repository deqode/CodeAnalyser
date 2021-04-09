package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/commands"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

//Commands is a plugin
type Commands struct {
}

func (d Commands) DetectBuildCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceOutputDetectBuildCommands, error) {
	return &pb.ServiceOutputDetectBuildCommands{
		Error:         nil,
		BuildCommands:[]*global.BuildCommandsOutput{
			{
				Used:          true,
				BuildCommands: nil,
			},
		} ,
	}, nil
}

func (d Commands) DetectStartUpCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceOutputStartUpCommands, error) {
	return nil, nil
}

func (d Commands) DetectSeedCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceOutputSeedCommands, error) {
	return nil, nil
}

func (d Commands) DetectMigrationCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceMigrationCommands, error) {
	return nil, nil
}

func (d Commands) DetectAdHocScripts(inputString *pb.ServiceCommandsInput) (*pb.ServiceAdHocScripts, error) {
	return nil,nil}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserCommand: &commands.GRPCPlugin{
				Impl: &Commands{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
