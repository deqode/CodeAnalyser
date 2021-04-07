package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/dependencies"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

type GetLibraries struct{}

func (g GetLibraries) GetLibraries(inputString *pb.ServiceInput) (*pb.ServiceOutputStringMap, error) {
	return &pb.ServiceOutputStringMap{
		Value: map[string]string{"kafka": "1.2"},
		Error: nil,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserLibraries: &dependencies.DependenciesGRPCPlugin{Impl: &GetLibraries{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
