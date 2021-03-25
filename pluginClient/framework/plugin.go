package framework

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type GreeterGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl interfaces.FrameworkVersions
}

func (p *GreeterGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterFrameworkServiceServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *GreeterGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewFrameworkServiceClient(c)}, nil
}
