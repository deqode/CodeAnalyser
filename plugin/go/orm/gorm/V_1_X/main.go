package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/orm"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

//GormV1X is a plugin of type ORM
type GormV1X struct{}

//Detect will return usage of ORM
func (g GormV1X) Detect(input *pb.ServiceInput) (*pb.ServiceOutputDetectOrm, error) {
	return &pb.ServiceOutputDetectOrm{
		Used:               true,
		SupportedDbName:    "postgres",
		SupportedDbVersion: "6.8",
		Error:              nil,
	}, nil
}

//IsORMUsed will return % usage of ORM
func (g GormV1X) IsORMUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Value: true,
		Error: nil,
	}, nil
}

//PercentOfORMUsed will return % usage of ORM
func (g GormV1X) PercentOfORMUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	return &pb.ServiceOutputFloat{
		Value: 89.5,
		Error: nil,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserOrm: &orm.GRPCPlugin{
				Impl: &GormV1X{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
