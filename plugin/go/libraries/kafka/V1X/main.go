package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/library"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	pluginPB "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

//KafkaV1x is a go plugin
type KafkaV1x struct{}

//Detect is KafkaV1x plugin method, it will detect kafka presence
func (k *KafkaV1x) Detect(input *pb.Input) (*pluginPB.LibraryType, error) {
	return &pluginPB.LibraryType{
		Value: true,
		Error: nil,
		Type:  languageSpecific.Type_EXTERNAL,
	}, nil
}

//IsUsed is a KafkaV1x plugin method, it will check is kafka used or not
func (k *KafkaV1x) IsUsed(input *pb.Input) (*pb.BoolOutput, error) {
	return &pb.BoolOutput{
		Value: true,
		Error: nil,
	}, nil
}

//PercentOfUsed is a KafkaV1x plugin method, it will check percent of used
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
