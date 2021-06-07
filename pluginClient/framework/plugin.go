package framework

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// GRPCPlugin is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type GRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl interfaces.Framework
}

//GRPCServer plugin.GRPCPlugin Implementation
func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterFrameworkServiceServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

//GRPCClient plugin.GRPCPlugin Implementation
func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewFrameworkServiceClient(c)}, nil
}
