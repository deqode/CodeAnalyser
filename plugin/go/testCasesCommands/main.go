package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/testCasesCommands"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

type TestCasesCommands struct{}

func (t *TestCasesCommands) Detect(input *pb.Input) (*pb.TestCommandOutput, error) {
	return &pb.TestCommandOutput{
		Commands: []*global.Command{
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
