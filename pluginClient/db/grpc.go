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

func (G *GRPCClient) GetVersionName() (*pb.ServiceOutputString, error) {
	res, err := G.Client.GetVersionName(context.Background(), &pb.ServiceEmpty{})
	return res, err
}

func (G *GRPCClient) GetSemver() (*pb.ServiceOutputString, error) {
	res, err := G.Client.GetSemver(context.Background(), &pb.ServiceEmpty{})
	return res, err
}

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

func (G *GRPCClient) IsDbFound(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := G.Client.IsDbFound(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}

func (G *GRPCClient) GetDbName() (*pb.ServiceOutputString, error) {
	res, err := G.Client.GetDbName(context.Background(), &pb.ServiceEmpty{})
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

func (m *GRPCServer) GetVersionName(context.Context, *pb.ServiceEmpty) (*pb.ServiceOutputString, error) {
	res, err := m.Impl.GetVersionName()
	return &pb.ServiceOutputString{
		Value: res.Value,
		Error: res.Error,
	}, err
}

func (m *GRPCServer) GetSemver(ctx context.Context, empty *pb.ServiceEmpty) (*pb.ServiceOutputString, error) {
	res, err := m.Impl.GetSemver()
	return &pb.ServiceOutputString{
		Value: res.Value,
		Error: res.Error,
	}, err
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

func (m *GRPCServer) IsDbFound(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := m.Impl.IsDbFound(input)
	return &pb.ServiceOutputBool{
		Value: res.Value,
		Error: res.Error,
	}, err
}

func (m *GRPCServer) GetDbName(ctx context.Context, input *pb.ServiceEmpty) (*pb.ServiceOutputString, error) {
	res, err := m.Impl.GetDbName()
	return &pb.ServiceOutputString{
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
