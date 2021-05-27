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

//TODO add logic for detection of same
func (d Commands) DetectBuildCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceOutputDetectBuildCommands, error) {
	return &pb.ServiceOutputDetectBuildCommands{
		Error: nil,
		BuildCommands: &global.BuildCommandsOutput{
			Used: true,
			BuildCommands: []*global.Command{
				{
					Command: "go",
					Args:    []string{"build one"},
				},
			},
		},
	}, nil
}

func (d Commands) DetectStartUpCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceOutputStartUpCommands, error) {
	return &pb.ServiceOutputStartUpCommands{
		Error: nil,
		StartUpCommands: &global.StartUpCommandsOutput{
			Used: true,
			StartUpCommands: []*global.Command{
				{
					Command: "go",
					Args:    []string{"build one"},
				},
			},
		},
	}, nil
}

func (d Commands) DetectSeedCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceOutputSeedCommands, error) {
	return &pb.ServiceOutputSeedCommands{
		Error: nil,
		SeedCommands: &global.SeedCommandsOutput{
			Used: true,
			SeedCommands: []*global.Command{
				{
					Command: "go",
					Args:    []string{"build one"},
				},
			},
		},
	}, nil
}

func (d Commands) DetectMigrationCommands(inputString *pb.ServiceCommandsInput) (*pb.ServiceMigrationCommands, error) {
	return &pb.ServiceMigrationCommands{
		Error: nil,
		MigrationCommands: &global.MigrationCommandsOutput{
			Used: true,
			MigrationCommands: []*global.Command{
				{
					Command: "go",
					Args:    []string{"build one"},
				},
			},
		},
	}, nil
}

func (d Commands) DetectAdHocScripts(inputString *pb.ServiceCommandsInput) (*pb.ServiceAdHocScripts, error) {
	return &pb.ServiceAdHocScripts{
		Error: nil,
		AdHocScripts: &global.AdHocScriptsOutput{
			Used: true,
			AdHocScripts: []*global.Command{
				{
					Command: "go",
					Args:    []string{"build one"},
				},
			},
		},
	}, nil
}

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
