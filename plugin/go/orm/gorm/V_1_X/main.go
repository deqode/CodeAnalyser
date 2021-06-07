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
func (g GormV1X) Detect(input *pb.Input) (*pb.OrmOutput, error) {
	return &pb.OrmOutput{
		Used:               true,
		DbName:    "postgres",
		DbVersion: "6.8",
		Error:              nil,
	}, nil
}

//IsORMUsed will return % usage of ORM
func (g GormV1X) IsUsed(input *pb.Input) (*pb.BoolOutput, error) {
	return &pb.BoolOutput{
		Value: true,
		Error: nil,
	}, nil
}

//PercentOfORMUsed will return % usage of ORM
func (g GormV1X) PercentOfUsed(input *pb.Input) (*pb.FloatOutput, error) {
	return &pb.FloatOutput{
		Value: 89.5,
		Error: nil,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.Orm: &orm.GRPCPlugin{
				Impl: &GormV1X{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
