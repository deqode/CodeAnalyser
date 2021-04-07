package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/library"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

type KafkaV1x struct{}

func (k *KafkaV1x) Detect(input *pb.ServiceInput) (*pb.ServiceOutputLibraryType, error) {
	return &pb.ServiceOutputLibraryType{
		Value: true,
		Error: nil,
		Type:  pb.ServiceOutputLibraryType_EXTERNAL,
	}, nil
}

func (k *KafkaV1x) IsUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Value: true,
		Error: nil,
	}, nil
}

func (k *KafkaV1x) PercentOfUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	return &pb.ServiceOutputFloat{
		Value: 70,
		Error: nil,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserLibrary: &library.GRPCPlugin{Impl: &KafkaV1x{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
