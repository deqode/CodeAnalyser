package orm

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
	"context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client plugin.OrmClient
}

//Detect method call on client side over gRPC
func (G *GRPCClient) Detect(input *pb.Input) (*plugin.OrmOutput, error) {
	res, err := G.Client.Detect(context.Background(), input)
	return res, err
}

//IsUsed method call on client side over gRPC
func (G *GRPCClient) IsUsed(input *pb.Input) (*pb.BoolOutput, error) {
	res, err := G.Client.IsUsed(context.Background(), input)
	return res, err
}

//PercentOfUsed method call on client side over gRPC
func (G *GRPCClient) PercentOfUsed(input *pb.Input) (*pb.FloatOutput, error) {
	res, err := G.Client.PercentOfUsed(context.Background(), input)
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl interfaces.Orm
}

//Detect will detect is given Orm found and Db used
func (G *GRPCServer) Detect(ctx context.Context, input *pb.Input) (*plugin.OrmOutput, error) {
	res, err := G.Impl.Detect(input)
	return res, err
}

//IsUsed will check if given Orm used or not
func (G *GRPCServer) IsUsed(ctx context.Context, input *pb.Input) (*pb.BoolOutput, error) {
	res, err := G.Impl.IsUsed(input)
	return res, err
}

//PercentOfUsed will check percent of given Orm used
func (G *GRPCServer) PercentOfUsed(ctx context.Context, input *pb.Input) (*pb.FloatOutput, error) {
	res, err := G.Impl.PercentOfUsed(input)
	return res, err
}
