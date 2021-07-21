package detectRuntime

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client plugin.DetectRuntimeClient
}

//Detect method call on client side over gRPC
func (m *GRPCClient) Detect(inputString *pb.StringInput) (*pb.StringOutput, error) {
	res, err := m.Client.Detect(context.Background(), inputString)
	if res != nil {
		return res, err
	}
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl interfaces.DetectRunTime
}

//Detect will return runtime of language
func (m *GRPCServer) Detect(ctx context.Context, inputString *pb.StringInput) (*pb.StringOutput, error) {
	res, err := m.Impl.Detect(inputString)
	return res, err
}
