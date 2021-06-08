package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/framework"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

//BeegoV1x is a framework Plugin
type BeegoV1x struct{}

//Detect will detect framework Used
func (b BeegoV1x) Detect(input *pb.Input) (*pb.BoolOutput, error) {
	return &pb.BoolOutput{
		Error: nil,
		Value: true,
	}, nil
}

//IsUsed is framework Used
func (b BeegoV1x) IsUsed(input *pb.Input) (*pb.BoolOutput, error) {
	return &pb.BoolOutput{
		Error: nil,
		Value: true,
	}, nil
}

//PercentOfUsed %  framework Used
func (b BeegoV1x) PercentOfUsed(input *pb.Input) (*pb.FloatOutput, error) {
	return &pb.FloatOutput{
		Error: nil,
		Value: 88.8,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.Framework: &framework.GRPCPlugin{Impl: &BeegoV1x{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
