package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/detectRuntime"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

type DetectGoRuntime struct{}

func (d DetectGoRuntime) DetectRuntime(inputString *pb.ServiceInputString) (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Value: "1.2",
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserDetectRuntime: &detectRuntime.DetectRuntimeGRPCPlugin{Impl: &DetectGoRuntime{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
