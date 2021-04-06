package env

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient/pb"
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type GRPCPlugin struct {
	plugin.Plugin
	Impl interfaces.Env
}

func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterEnvServiceServer(server, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewEnvServiceClient(conn)}, nil
}
