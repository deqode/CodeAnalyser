package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/library"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

type KafkaV1x struct{}

func (k *KafkaV1x) Detect(input *pb.Input) (*pb.LibraryType, error) {
	return &pb.LibraryType{
		Value: true,
		Error: nil,
		Type:  pb.LibraryType_EXTERNAL,
	}, nil
}

func (k *KafkaV1x) IsUsed(input *pb.Input) (*pb.BoolOutput, error) {
	return &pb.BoolOutput{
		Value: true,
		Error: nil,
	}, nil
}

func (k *KafkaV1x) PercentOfUsed(input *pb.Input) (*pb.FloatOutput, error) {
	return &pb.FloatOutput{
		Value: 70,
		Error: nil,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.Library: &library.GRPCPlugin{Impl: &KafkaV1x{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
