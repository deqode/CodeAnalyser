package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/dependencies"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

//GetGoDependencies is a plugin of type get language dependencies
type GetGoDependencies struct{}

//GetDependencies will return map of dependencies and versions
func (d *GetGoDependencies) GetDependencies(inputString *pb.ServiceInput) (*pb.ServiceOutputStringMap, error) {
	return &pb.ServiceOutputStringMap{
		Value: map[string]string{"beego": "1.2", "gorm": "1.1", "postgres": "1.1"},
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserDependencies: &dependencies.GRPCPlugin{Impl: &GetGoDependencies{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
