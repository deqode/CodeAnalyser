package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/env"
	pb "code-analyser/protos/pb/helpers"
	languagePB "code-analyser/protos/pb/output/languageSpecific"
	"github.com/hashicorp/go-plugin"
)

//EnvPlugin is a plugin to check env available
type EnvPlugin struct{}

//Detect function for envs keyvalues and keys implemented from Env interface
func (e EnvPlugin) Detect(input *pb.Input) (*languagePB.Envs, error) {
	return &languagePB.Envs{
		EnvKeyValues: map[string]string{
			"username": "boss",
			"password": "aag_laga_denge_aag",
		},
		Keys:  []string{"345"},
		Used:  true,
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
