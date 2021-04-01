package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/dependencies"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

type GetGoDependencies struct{}

func (d *GetGoDependencies) GetDependencies(inputString *pb.ServiceInput) (*pb.ServiceOutputStringMap, error) {
	return &pb.ServiceOutputStringMap{
		Value: map[string]string{"beego":"1.2","gorm":"1.1","postgres":"1.1"},
	},nil
}

func main(){
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig:  pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserDependencies: &dependencies.DependenciesGRPCPlugin{Impl: &GetGoDependencies{}},
		},
		GRPCServer:       plugin.DefaultGRPCServer,
	})
}

