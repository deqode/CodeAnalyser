package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/staticAssets"
	"code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

type GetStaticAssets struct{}

func (g *GetStaticAssets) Detect(input *pb.Input) (*pb.StaticAssetsOutput, error) {
	return &pb.StaticAssetsOutput{
		Error: nil,
		StaticAsset: []*languageSpecific.StaticAsset{
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
			pluginClient.StaticAssets: &staticAssets.GRPCPlugin{Impl: &GetStaticAssets{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
