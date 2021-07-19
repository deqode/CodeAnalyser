package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/staticAssets"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	"github.com/hashicorp/go-plugin"
)

//StaticAssets is a go plugin
type StaticAssets struct{}

//Detect is a StaticAssets plugin method, it will detect StaticAsset
func (g *StaticAssets) Detect(input *pb.Input) (*languageSpecific.StaticAssetsOutput, error) {
	return &languageSpecific.StaticAssetsOutput{
		Error: nil,
		Used:  true,
		StaticAssets: []*languageSpecific.StaticAsset{
			{
				Name:   "static",
				Path:   "hghkjg",
				Hashed: true,
				Public: true,
			},
		},
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.StaticAssets: &staticAssets.GRPCPlugin{Impl: &StaticAssets{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
