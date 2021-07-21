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

//DetectBuildCommands is a Commands' plugin method, it will detect build commands
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

//DetectStartUpCommands is a Commands' plugin method, it will detect startup commands
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

//DetectSeedCommands is a Commands' plugin method, it will detect seed commands
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

//DetectMigrationCommands is a Commands' plugin method, it will detect migration commands
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

//DetectAdHocScripts is a Commands' plugin method, it will detect adHoc-scripts
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
