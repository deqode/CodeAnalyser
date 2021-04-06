package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/framework"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

//BeegoV1x is a framework Plugin
type BeegoV1x struct{}

//Detect will detect framework Used
func (b BeegoV1x) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Error: nil,
		Value: true,
	}, nil
}

//IsFrameworkUsed is framework Used
func (b BeegoV1x) IsFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Error: nil,
		Value: true,
	}, nil
}

//PercentOfFrameworkUsed %  framework Used
func (b BeegoV1x) PercentOfFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	return &pb.ServiceOutputFloat{
		Error: nil,
		Value: 88.8,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserFramework: &framework.GRPCPlugin{Impl: &BeegoV1x{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
