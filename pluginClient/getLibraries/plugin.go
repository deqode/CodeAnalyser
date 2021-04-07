package getLibraries

import (
	"code-analyser/language_detectors/interfaces"
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type GRPCPlugin struct {
	plugin.Plugin
	Impl interfaces.GetLibraries
}

func (g *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {

}

func (g *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	panic("implement me")
}
