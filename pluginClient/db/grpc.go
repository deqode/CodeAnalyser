package db

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/protos/pb/helpers"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client pb.DbClient
}

//TODO handle all errors

//Detect method call on client side over gRPC
func (G *GRPCClient) Detect(input *helpers.Input) (*pb.BoolIntOutput, error) {
	res, err := G.Client.Detect(context.Background(), &helpers.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return res, err
}

//IsUsed method call on client side over gRPC
func (G *GRPCClient) IsUsed(input *helpers.Input) (*helpers.BoolOutput, error) {
	res, err := G.Client.IsUsed(context.Background(), &helpers.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return res, err
}

//PercentOfUsed method call on client side over gRPC
func (G *GRPCClient) PercentOfUsed(input *helpers.Input) (*helpers.FloatOutput, error) {
	res, err := G.Client.PercentOfUsed(context.Background(), &helpers.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl interfaces.Db
}

//Detect will detect is given Db found and port number
func (m *GRPCServer) Detect(ctx context.Context, input *helpers.Input) (*pb.BoolIntOutput, error) {
	res, err := m.Impl.Detect(input)
	return &pb.BoolIntOutput{
		Value:    res.Value,
		IntValue: res.IntValue,
		Error:    res.Error,
	}, err
}

//IsUsed will check is db used
func (m *GRPCServer) IsUsed(ctx context.Context, input *helpers.Input) (*helpers.BoolOutput, error) {
	res, err := m.Impl.IsUsed(input)
	return &helpers.BoolOutput{
		Value: res.Value,
		Error: res.Error,
	}, err
}

//PercentOfUsed will check percent of given db used
func (m *GRPCServer) PercentOfUsed(ctx context.Context, input *helpers.Input) (*helpers.FloatOutput, error) {
	res, err := m.Impl.PercentOfUsed(input)
	return &helpers.FloatOutput{
		Value: res.Value,
		Error: res.Error,
	}, err
}
