package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/db"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

type Postgres_V_1_X struct {
}

func (receiver *Postgres_V_1_X) IsDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Value: true,
		Error: nil,
	}, nil
}

func (receiver *Postgres_V_1_X) PercentOfDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	return &pb.ServiceOutputFloat{
		Error: nil,
		Value: 88,
	}, nil
}

// it returns port as well
func (receiver *Postgres_V_1_X) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBoolInt, error) {
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
				Impl: &Postgres_V_1_X{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
