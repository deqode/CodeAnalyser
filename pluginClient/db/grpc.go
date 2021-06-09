package db

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/protos/pb/helpers"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of Framework that talks over RPC.
type GRPCClient struct {
	Client pb.DbClient
}

//TODO handle all errors

//Detect will detect if DB used
func (G *GRPCClient) Detect(input *helpers.Input) (*pb.BoolIntOutput, error) {
	res, err := G.Client.Detect(context.Background(), &helpers.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return res, err
}

//IsUsed will return true if DB used
func (G *GRPCClient) IsUsed(input *helpers.Input) (*helpers.BoolOutput, error) {
	res, err := G.Client.IsUsed(context.Background(), &helpers.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return res, err
}

//PercentOfUsed will return % of  db used
func (G *GRPCClient) PercentOfUsed(input *helpers.Input) (*helpers.FloatOutput, error) {
	res, err := G.Client.PercentOfUsed(context.Background(), &helpers.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return res, err
}

//GRPCServer  is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	Impl interfaces.Db
}

//Detect will Detect if DB used
func (m *GRPCServer) Detect(ctx context.Context, input *helpers.Input) (*pb.BoolIntOutput, error) {
	res, err := m.Impl.Detect(input)
	return &pb.BoolIntOutput{
		Value:    res.Value,
		IntValue: res.IntValue,
		Error:    res.Error,
	}, err
}

//IsUsed will return true if DB used
func (m *GRPCServer) IsUsed(ctx context.Context, input *helpers.Input) (*helpers.BoolOutput, error) {
	res, err := m.Impl.IsUsed(input)
	return &helpers.BoolOutput{
		Value: res.Value,
		Error: res.Error,
	}, err
}

//PercentOfUsed will return % of  db used
func (m *GRPCServer) PercentOfUsed(ctx context.Context, input *helpers.Input) (*helpers.FloatOutput, error) {
	res, err := m.Impl.PercentOfUsed(input)
	return &helpers.FloatOutput{
		Value: res.Value,
		Error: res.Error,
	}, err
}
