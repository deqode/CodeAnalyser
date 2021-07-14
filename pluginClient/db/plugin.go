package db

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// GRPCPlugin is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type GRPCPlugin struct {
	plugin.Plugin

	Impl interfaces.Db
}

//GRPCServer plugin.GRPCPlugin Implementation
func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterDbServer(server, &GRPCServer{Impl: p.Impl})
	return nil
}

//GRPCClient plugin.GRPCPlugin Implementation
func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewDbClient(conn)}, nil
}
