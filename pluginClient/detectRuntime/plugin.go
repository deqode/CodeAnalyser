package detectRuntime

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient/pb"
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)
// GRPCPlugin is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type GRPCPlugin struct {
	plugin.Plugin
	Impl interfaces.DetectRunTime
}

//GRPCServer plugin.GRPCPlugin Implementation
func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterDetectRuntimeServiceServer(server, &GRPCServer{Impl: p.Impl})
	return nil
}

//GRPCClient plugin.GRPCPlugin Implementation
func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{Client: pb.NewDetectRuntimeServiceClient(conn)}, nil
}
