package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/detectRuntime"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

//DetectGoRuntime to detect runtime version of language
type DetectGoRuntime struct{}

//Detect to detect runtime version of language
func (d DetectGoRuntime) Detect(path *pb.StringInput) (*pb.StringOutput, error) {
	return &pb.StringOutput{
		Value: "1.2",
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.DetectRuntime: &detectRuntime.GRPCPlugin{Impl: &DetectGoRuntime{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
