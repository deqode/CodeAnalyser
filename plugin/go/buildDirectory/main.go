package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/buildDirectory"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

type BuildDirectpryPlugin struct{}

func (b *BuildDirectpryPlugin) Detect(input *pb.ServiceInput) (*pb.ServiceOutputStringMap, error) {
	return &pb.ServiceOutputStringMap{
		Value: map[string]string{
			"build":   "/project/name/bin",
			"publish": "npm run build:lib && publish",
		},
		Error: nil,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserBuildDirectory: &buildDirectory.GRPCPlugin{Impl: &BuildDirectpryPlugin{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
