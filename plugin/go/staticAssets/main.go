package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/staticAssets"
	"code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

type GetStaticAssets struct{}

func (g *GetStaticAssets) Detect(input *pb.ServiceInput) (*pb.ServiceOutputStaticAssets, error) {
	return &pb.ServiceOutputStaticAssets{
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
			pluginClient.PluginDispenserStaticAssets: &staticAssets.GRPCPlugin{Impl: &GetStaticAssets{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
