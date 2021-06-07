package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/preDetectCommand"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

type PreDetectCommand struct{}

func (p *PreDetectCommand) RunPreDetect(input *pb.Input) (*pb.EmptyOutput, error) {
	return &pb.EmptyOutput{Error: nil}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PreDetectCommand: &preDetectCommand.GRPCPlugin{Impl: &PreDetectCommand{}},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
