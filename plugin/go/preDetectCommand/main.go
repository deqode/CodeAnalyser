package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/pluginClient/preDetectCommand"
	"github.com/hashicorp/go-plugin"
)

type PreDetectCommand struct{}

func (p *PreDetectCommand) RunPreDetect(input *pb.ServiceInput) (*pb.ServiceEmptyOutput, error) {
	return &pb.ServiceEmptyOutput{Error: nil}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserPreDetectCommand: &preDetectCommand.GRPCPlugin{Impl: &PreDetectCommand{}},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
