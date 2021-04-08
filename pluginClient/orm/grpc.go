package orm

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/plugin"
	"context"
)

// GRPCClient is an implementation of FrameworkVersions that talks over RPC.
type GRPCClient struct {
	Client pb.OrmServiceClient
}

//Detect will detect orm used and supported DB
func (G *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputDetectOrm, error) {
	res, err := G.Client.Detect(context.Background(), input)
	return res, err
}

//IsORMUsed will  check if ORM used or not
func (G *GRPCClient) IsORMUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := G.Client.IsORMUsed(context.Background(), input)
	return res, err
}

//PercentOfORMUsed will check % of orm used
func (G *GRPCClient) PercentOfORMUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	res, err := G.Client.PercentOfORMUsed(context.Background(), input)
	return res, err
}

//GRPCServer  is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	Impl interfaces.ORMVersion
}

//Detect will detect orm used and supported DB
func (G *GRPCServer) Detect(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputDetectOrm, error) {
	res, err := G.Impl.Detect(input)
	return res, err
}

//IsORMUsed will  check if ORM used or not
func (G *GRPCServer) IsORMUsed(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := G.Impl.IsORMUsed(input)
	return res, err
}

//PercentOfORMUsed will check % of orm used
func (G *GRPCServer) PercentOfORMUsed(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	res, err := G.Impl.PercentOfORMUsed(input)
	return res, err
}
