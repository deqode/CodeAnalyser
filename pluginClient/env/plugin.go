package env

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type EnvGRPCPlugin struct {
	plugin.Plugin
	Impl interfaces.Env
}

func (p *EnvGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterEnvServiceServer(server, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *EnvGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewEnvServiceClient(conn)}, nil
}
