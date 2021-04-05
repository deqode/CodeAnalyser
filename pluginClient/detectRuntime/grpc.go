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
//DetectRuntime will return runtime of language
func (m *GRPCClient) DetectRuntime(inputString *pb.ServiceInputString) (*pb.ServiceOutputString, error) {
	res, err := m.Client.DetectRuntime(context.Background(), inputString)
	if res != nil {
		return res, err
	}
	return res, err
}
//DetectRuntime will return runtime of language
func (m *GRPCServer) DetectRuntime(ctx context.Context, inputString *pb.ServiceInputString) (*pb.ServiceOutputString, error) {
	res, err := m.Impl.DetectRuntime(inputString)
	return res, err
}
