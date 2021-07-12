package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/buildDirectory"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	"github.com/hashicorp/go-plugin"
)

type BuildDirectpryPlugin struct{}

func (b *BuildDirectpryPlugin) Detect(input *pb.Input) (*languageSpecific.BuildDirectoryOutput, error) {
	return &languageSpecific.BuildDirectoryOutput{
		Used: true,
		BuildDirectory: []*languageSpecific.BuildDirectory{{
			Name: "josha",
			Root: "dfgdf",
		}},
		Error: nil,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.BuildDirectory: &buildDirectory.GRPCPlugin{Impl: &BuildDirectpryPlugin{}},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
