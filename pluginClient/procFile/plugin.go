package procFile

import (
	"code-analyser/GlobalFiles"
	pb "code-analyser/protos/pb/plugin"
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type GRPCPlugin struct {
	plugin.Plugin
	Impl GlobalFiles.ProcFile
}

func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterProcFileServer(server, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewProcFileClient(conn)}, nil
}
