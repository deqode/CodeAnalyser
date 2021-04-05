package orm

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"context"
)

type GRPCClient struct {
	Client pb.OrmServiceClient
}

func (G *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputDetectOrm, error) {
	res, err := G.Client.Detect(context.Background(), input)
	return res, err
}

func (G *GRPCClient) IsORMUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := G.Client.IsORMUsed(context.Background(), input)
	return res, err
}

func (G *GRPCClient) PercentOfORMUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	res, err := G.Client.PercentOfORMUsed(context.Background(), input)
	return res, err
}

type GRPCServer struct {
	Impl interfaces.ORMVersion
}

func (G *GRPCServer) Detect(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputDetectOrm, error) {
	res, err := G.Impl.Detect(input)
	return res, err
}

func (G *GRPCServer) IsORMUsed(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := G.Impl.IsORMUsed(input)
	return res, err
}

func (G *GRPCServer) PercentOfORMUsed(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	res, err := G.Impl.PercentOfORMUsed(input)
	return res, err
}
