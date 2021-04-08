package db

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)
// GRPCClient is an implementation of FrameworkVersions that talks over RPC.
type GRPCClient struct {
	Client pb.DbServiceClient
}
//TODO handle all errors

//Detect will detect if DB used
func (G *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBoolInt, error) {
	res, err := G.Client.Detect(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}

//IsDbUsed will return true if DB used
func (G *GRPCClient) IsDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := G.Client.IsDbUsed(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}

//PercentOfDbUsed will return % of  db used
func (G *GRPCClient) PercentOfDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	res, err := G.Client.PercentOfDbUsed(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}
//GRPCServer  is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	Impl interfaces.DbVersion
}

//Detect will Detect if DB used
func (m *GRPCServer) Detect(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputBoolInt, error) {
	res, err := m.Impl.Detect(input)
	return &pb.ServiceOutputBoolInt{
		Value:    res.Value,
		IntValue: res.IntValue,
		Error:    res.Error,
	}, err
}

//IsDbUsed will return true if DB used
func (m *GRPCServer) IsDbUsed(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := m.Impl.IsDbUsed(input)
	return &pb.ServiceOutputBool{
		Value: res.Value,
		Error: res.Error,
	}, err
}

//PercentOfDbUsed will return % of  db used
func (m *GRPCServer) PercentOfDbUsed(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	res, err := m.Impl.PercentOfDbUsed(input)
	return &pb.ServiceOutputFloat{
		Value: res.Value,
		Error: res.Error,
	}, err
}
