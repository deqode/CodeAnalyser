package procFile

import (
	"code-analyser/GlobalFiles"
	pb "code-analyser/protos/pb/plugin"
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

//GRPCPlugin is an implementation of plugin.GRPCPlugin, so we can serve/consume this struct
type GRPCPlugin struct {
	plugin.Plugin
	Impl GlobalFiles.ProcFile
}

//GRPCServer is the gRPC server and implementation of GRPCServer to communicate with gRPC client
func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterProcFileServer(server, &GRPCServer{Impl: p.Impl})
	return nil
}

//GRPCClient is the gRPC server and implementation of GRPCClient to communicate with gRPC client
func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewProcFileClient(conn)}, nil
}
