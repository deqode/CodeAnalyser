package detectRuntime

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type DetectRuntimeGRPCPlugin struct {
	plugin.Plugin
	Impl interfaces.DetectRunTime
}

func (p *DetectRuntimeGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterDetectRuntimeServiceServer(server,&GRPCServer{Impl: p.Impl})
	return nil
}

func (p *DetectRuntimeGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
return &GRPCClient{Client: pb.NewDetectRuntimeServiceClient(conn)},nil
}
