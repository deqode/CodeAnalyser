package db

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type DbGRPCPlugin struct {
	plugin.Plugin

	Impl interfaces.DbVersion
}

func (p *DbGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterDbServiceServer(server, &GRPCServer{Impl: p.Impl})
	return nil
}


func (p *DbGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewDbServiceClient(conn)},nil
}


