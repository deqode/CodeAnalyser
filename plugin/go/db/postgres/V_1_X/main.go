package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/db"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

type PostgresV1X struct {
}
//IsDbUsed need to add cmt
func (receiver *PostgresV1X) IsDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Value: true,
		Error: nil,
	}, nil
}
//PercentOfDbUsed need to add cmt
func (receiver *PostgresV1X) PercentOfDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	return &pb.ServiceOutputFloat{
		Error: nil,
		Value: 88,
	}, nil
}

// Detect it returns port as well
func (receiver *PostgresV1X) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBoolInt, error) {
	return &pb.ServiceOutputBoolInt{
		Value:    true,
		IntValue: 6588,
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserDB: &db.DbGRPCPlugin{
				Impl: &PostgresV1X{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
