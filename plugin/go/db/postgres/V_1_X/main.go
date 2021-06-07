package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/db"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

//PostgresV1X is a plugin
type PostgresV1X struct {
}

//IsDbUsed if DB used in repo
func (receiver *PostgresV1X) IsUsed(input *pb.Input) (*pb.BoolOutput, error) {
	return &pb.BoolOutput{
		Value: true,
		Error: nil,
	}, nil
}

//PercentOfDbUsed % of DB used
func (receiver *PostgresV1X) PercentOfUsed(input *pb.Input) (*pb.FloatOutput, error) {
	return &pb.FloatOutput{
		Error: nil,
		Value: 88,
	}, nil
}

// Detect it returns port as well
func (receiver *PostgresV1X) Detect(input *pb.Input) (*pb.BoolIntOutput, error) {
	return &pb.BoolIntOutput{
		Value:    true,
		IntValue: 6588,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.DB: &db.GRPCPlugin{
				Impl: &PostgresV1X{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
