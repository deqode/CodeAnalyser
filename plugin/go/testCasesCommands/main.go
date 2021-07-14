package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/testCasesCommands"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	"github.com/hashicorp/go-plugin"
)

type TestCasesCommands struct{}

func (t *TestCasesCommands) Detect(input *pb.Input) (*languageSpecific.TestCasesCommand, error) {
	return &languageSpecific.TestCasesCommand{
		Commands: []*pb.Command{
			{
				Command: "go",
				Args:    []string{"run", "test"},
			},
		},
		Error: nil,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.TestCaseCommand: &testCasesCommands.GRPCPlugin{Impl: &TestCasesCommands{}},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
