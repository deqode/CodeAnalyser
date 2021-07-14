package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/dependencies"
	pb "code-analyser/protos/pb/helpers"
	"github.com/hashicorp/go-plugin"
)

//GetGoDependencies is a plugin of type get language dependencies
type GetGoDependencies struct{}

//GetDependencies will return map of dependencies and pluginDetails
func (d *GetGoDependencies) GetDependencies(input *pb.Input) (*pb.StringMapOutput, error) {
	return &pb.StringMapOutput{
		Value: map[string]string{
			"beego":    "1.2",
			"gorm":     "1.1",
			"postgres": "1.1",
			"kafka":    "1.2",
		},
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.Dependencies: &dependencies.GRPCPlugin{Impl: &GetGoDependencies{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
