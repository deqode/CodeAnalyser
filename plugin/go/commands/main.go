package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/commands"
	pb "code-analyser/protos/pb/helpers"
	LanguagePB "code-analyser/protos/pb/output/languageSpecific"
	"github.com/hashicorp/go-plugin"
)

//Commands is a plugin
type Commands struct {
}

func (d Commands) DetectBuildCommands(path *pb.StringInput) (*LanguagePB.CommandOutput, error) {
	return &LanguagePB.CommandOutput{
		Error: nil,
		Used:  true,
		Commands: []*pb.Command{
			{
				Command: "go",
				Args:    []string{"build one"},
			},
		},
	}, nil
}

func (d Commands) DetectStartUpCommands(path *pb.StringInput) (*LanguagePB.CommandOutput, error) {
	return &LanguagePB.CommandOutput{
		Error: nil,
		Used:  true,
		Commands: []*pb.Command{
			{
				Command: "go",
				Args:    []string{"start up"},
			},
		},
	}, nil
}

func (d Commands) DetectSeedCommands(path *pb.StringInput) (*LanguagePB.CommandOutput, error) {
	return &LanguagePB.CommandOutput{
		Error: nil,
		Used:  true,
		Commands: []*pb.Command{
			{
				Command: "go",
				Args:    []string{"seed "},
			},
		},
	}, nil
}

func (d Commands) DetectMigrationCommands(path *pb.StringInput) (*LanguagePB.CommandOutput, error) {
	return &LanguagePB.CommandOutput{
		Error: nil,
		Used:  true,
		Commands: []*pb.Command{
			{
				Command: "go",
				Args:    []string{"migration "},
			},
		},
	}, nil
}

func (d Commands) DetectAdHocScripts(path *pb.StringInput) (*LanguagePB.CommandOutput, error) {
	return &LanguagePB.CommandOutput{
		Error: nil,
		Used:  true,
		Commands: []*pb.Command{
			{
				Command: "go",
				Args:    []string{"adhoc "},
			},
		},
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.Command: &commands.GRPCPlugin{
				Impl: &Commands{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
