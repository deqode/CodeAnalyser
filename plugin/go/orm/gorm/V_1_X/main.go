package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/orm"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

type Gorm_V_1_X struct{}

func (g Gorm_V_1_X) Detect(input *pb.ServiceInput) (*pb.ServiceOutputDetectOrm, error) {
	return &pb.ServiceOutputDetectOrm{
		Used:               true,
		SupportedDbName:    "postgres",
		SupportedDbVersion: "6.8",
		Error:              nil,
	}, nil
}

func (g Gorm_V_1_X) IsORMUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Value: true,
		Error: nil,
	}, nil
}

func (g Gorm_V_1_X) PercentOfORMUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	return &pb.ServiceOutputFloat{
		Value: 89.5,
		Error: nil,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserOrm: &orm.OrmGRPCPlugin{
				Impl: &Gorm_V_1_X{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
