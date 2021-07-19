package db

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//GRPCPlugin is an implementation of plugin.GRPCPlugin, so we can serve/consume this struct
type GRPCPlugin struct {
	plugin.Plugin

	Impl interfaces.Db
}

//GRPCServer is the gRPC server and implementation of GRPCServer to communicate with gRPC client
func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterDbServer(server, &GRPCServer{Impl: p.Impl})
	return nil
}

//GRPCClient is the gRPC server and implementation of GRPCClient to communicate with gRPC client
func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewDbClient(conn)}, nil
}
