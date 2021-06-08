package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/env"
	"code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

//EnvPlugin is a plugin to check env available
type EnvPlugin struct{}

//Detect function for envs keyvalues and keys implemented from Env interface
func (e EnvPlugin) Detect(input *pb.Input) (*pb.EnvsOutput, error) {
	return &pb.EnvsOutput{
		Value:&languageSpecific.Envs{
			EnvKeyValues: map[string]string{
				"username": "boss",
				"password": "aag_laga_denge_aag",
			},
			Keys:  []string{"345"},
			Used: true,
		} ,
		Error: nil,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.Env: &env.GRPCPlugin{Impl: &EnvPlugin{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
