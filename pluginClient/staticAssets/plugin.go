package staticAssets

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/plugin"
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

//GRPCPlugin is an implementation of plugin.GRPCPlugin, so we can serve/consume this struct
type GRPCPlugin struct {
	plugin.Plugin
	Impl interfaces.StaticAssets
}

//GRPCServer is the gRPC server and implementation of GRPCServer to communicate with gRPC client
func (g *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterStaticAssetsServer(server, &GRPCServer{Impl: g.Impl})
	return nil
}

//GRPCClient is the gRPC server and implementation of GRPCClient to communicate with gRPC client
func (g *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewStaticAssetsClient(conn)}, nil
}
