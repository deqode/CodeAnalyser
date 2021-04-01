package detectRuntime

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client pb.DetectRuntimeServiceClient
}

type GRPCServer struct {
	Impl interfaces.DetectRunTime
}

func (m *GRPCClient) DetectRuntime(inputString *pb.ServiceInputString) (*pb.ServiceOutputString, error) {
	res, err := m.Client.DetectRuntime(context.Background(), inputString)
	if res != nil {
		return res, err
	}
	return res, err
}

func (m *GRPCServer) DetectRuntime(ctx context.Context, inputString *pb.ServiceInputString) (*pb.ServiceOutputString, error) {
	res, err := m.Impl.DetectRuntime(inputString)
	return res, err
}
