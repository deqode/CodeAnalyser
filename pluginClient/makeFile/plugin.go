package makeFile

import (
	"code-analyser/GlobalFiles"
	pb "code-analyser/protos/pb/plugin"
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type GRPCPlugin struct {
	plugin.Plugin
	Impl GlobalFiles.Makefile
}

func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterMakeFileServiceServer(server, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewMakeFileServiceClient(conn)}, nil
}
