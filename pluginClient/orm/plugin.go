package orm

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type OrmGRPCPlugin struct {
	plugin.Plugin

	Impl interfaces.ORMVersion
}

func (p *OrmGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterOrmServiceServer(server,&GRPCServer{Impl: p.Impl})
   return nil
}

func (p *OrmGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
   return &GRPCClient{Client: pb.NewOrmServiceClient(conn)},nil
}
