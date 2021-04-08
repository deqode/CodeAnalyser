package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/detectRuntime"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

//DetectGoRuntime to detect runtime version of language
type DetectGoRuntime struct{}

//DetectRuntime to detect runtime version of language
func (d DetectGoRuntime) DetectRuntime(inputString *pb.ServiceInputString) (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Value: "1.2",
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserDetectRuntime: &detectRuntime.GRPCPlugin{Impl: &DetectGoRuntime{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
