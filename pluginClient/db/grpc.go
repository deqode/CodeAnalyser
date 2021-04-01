package db

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client pb.DbServiceClient
}

//TODO handle all errors

func (G *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBoolInt, error) {
	res, err := G.Client.Detect(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}

func (G *GRPCClient) IsDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := G.Client.IsDbUsed(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}

func (G *GRPCClient) PercentOfDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	res, err := G.Client.PercentOfDbUsed(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}

type GRPCServer struct {
	Impl interfaces.DbVersion
}

func (m *GRPCServer) Detect(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputBoolInt, error) {
	res, err := m.Impl.Detect(input)
	return &pb.ServiceOutputBoolInt{
		Value:    res.Value,
		IntValue: res.IntValue,
		Error:    res.Error,
	}, err
}

func (m *GRPCServer) IsDbUsed(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := m.Impl.IsDbUsed(input)
	return &pb.ServiceOutputBool{
		Value: res.Value,
		Error: res.Error,
	}, err
}


func (m *GRPCServer) PercentOfDbUsed(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	res, err := m.Impl.PercentOfDbUsed(input)
	return &pb.ServiceOutputFloat{
		Value: res.Value,
		Error: res.Error,
	}, err
}
