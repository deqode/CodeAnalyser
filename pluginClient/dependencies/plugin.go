package dependencies

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type DependenciesGRPCPlugin struct {
	plugin.Plugin
	Impl interfaces.Dependencies
}

func (p *DependenciesGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterDependenciesServiceServer(server, &GRPCServer{
		Impl: p.Impl,
	})
	return nil
}

func (p *DependenciesGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewDependenciesServiceClient(conn)}, nil
}
