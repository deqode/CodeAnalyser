package detectRuntime

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of Framework that talks over RPC.
type GRPCClient struct {
	Client pb.DetectRuntimeServiceClient
}

//DetectRuntime will return runtime of language
func (m *GRPCClient) Detect(inputString *pb.ServiceInputString) (*pb.ServiceOutputString, error) {
	res, err := m.Client.DetectRuntime(context.Background(), inputString)
	if res != nil {
		return res, err
	}
	return res, err
}

//GRPCServer  is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	Impl interfaces.DetectRunTime
}

//DetectRuntime will return runtime of language
func (m *GRPCServer) DetectRuntime(ctx context.Context, inputString *pb.ServiceInputString) (*pb.ServiceOutputString, error) {
	res, err := m.Impl.Detect(inputString)
	return res, err
}
