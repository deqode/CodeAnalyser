package orm

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
	"context"
)

// GRPCClient is an implementation of Framework that talks over RPC.
type GRPCClient struct {
	Client plugin.OrmClient
}

//Detect will detect orm used and supported DB
func (G *GRPCClient) Detect(input *pb.Input) (*plugin.OrmOutput, error) {
	res, err := G.Client.Detect(context.Background(), input)
	return res, err
}

//IsUsed will  check if ORM used or not
func (G *GRPCClient) IsUsed(input *pb.Input) (*pb.BoolOutput, error) {
	res, err := G.Client.IsORMUsed(context.Background(), input)
	return res, err
}

//PercentOfUsed will check % of orm used
func (G *GRPCClient) PercentOfUsed(input *pb.Input) (*pb.FloatOutput, error) {
	res, err := G.Client.PercentOfORMUsed(context.Background(), input)
	return res, err
}

//GRPCServer  is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	Impl interfaces.ORMVersion
}

//Detect will detect orm used and supported DB
func (G *GRPCServer) Detect(ctx context.Context, input *pb.Input) (*plugin.OrmOutput, error) {
	res, err := G.Impl.Detect(input)
	return res, err
}

//IsORMUsed will  check if ORM used or not
func (G *GRPCServer) IsORMUsed(ctx context.Context, input *pb.Input) (*pb.BoolOutput, error) {
	res, err := G.Impl.IsUsed(input)
	return res, err
}

//PercentOfORMUsed will check % of orm used
func (G *GRPCServer) PercentOfORMUsed(ctx context.Context, input *pb.Input) (*pb.FloatOutput, error) {
	res, err := G.Impl.PercentOfUsed(input)
	return res, err
}
