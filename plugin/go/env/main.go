package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/env"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

type EnvPlugin struct{}

func (e EnvPlugin) Detect(input *pb.ServiceInput) (*pb.ServiceOutputEnv, error) {
	return &pb.ServiceOutputEnv{
		EnvKeyValues: map[string]string{
			"username": "boss",
			"password": "aag_laga_denge_aag",
		},
		Keys:  []string{"345"},
		Error: nil,
	}, nil
}

func (e EnvPlugin) IsUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Value: true,
		Error: nil,
	}, nil
}

func main()  {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserEnv: &env.EnvGRPCPlugin{Impl: &EnvPlugin{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}