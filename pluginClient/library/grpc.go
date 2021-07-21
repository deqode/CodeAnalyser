package library

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
	"code-analyser/utils"
	"errors"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client plugin.LibraryClient
}

//Detect method call on client side over gRPC
func (g *GRPCClient) Detect(input *pb.Input) (*plugin.LibraryType, error) {
	if input != nil {
		res, err := g.Client.Detect(context.Background(), &pb.Input{
			RuntimeVersion: input.RuntimeVersion,
			RootPath:       input.RootPath,
		})
		return res, err
	}
	utils.Logger("service input nil found in library grpc detect method")
	return nil, errors.New("service input nil found")
}

//IsUsed method call on client side over gRPC
func (g *GRPCClient) IsUsed(serviceInput *pb.Input) (*pb.BoolOutput, error) {
	res, err := g.Client.IsUsed(context.Background(), &pb.Input{
		RuntimeVersion: serviceInput.RuntimeVersion,
		RootPath:       serviceInput.RootPath,
	})
	return res, err
}

//PercentOfUsed method call on client side over gRPC
func (g *GRPCClient) PercentOfUsed(serviceInput *pb.Input) (*pb.FloatOutput, error) {
	res, err := g.Client.PercentOfUsed(context.Background(), &pb.Input{
		RuntimeVersion: serviceInput.RuntimeVersion,
		RootPath:       serviceInput.RootPath,
	})
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl interfaces.Library
}

//Detect will detect is given Library found and port number
func (g *GRPCServer) Detect(ctx context.Context, serviceInput *pb.Input) (*plugin.LibraryType, error) {
	res, err := g.Impl.Detect(serviceInput)
	return &plugin.LibraryType{
		Type:  res.Type,
		Value: res.Value,
		Error: res.Error,
	}, err
}

//IsUsed will check is Library used
func (g *GRPCServer) IsUsed(ctx context.Context, serviceInput *pb.Input) (*pb.BoolOutput, error) {
	res, err := g.Impl.IsUsed(serviceInput)
	return &pb.BoolOutput{
		Value: res.Value,
		Error: res.Error,
	}, err
}

//PercentOfUsed will check percent of given Db used
func (g *GRPCServer) PercentOfUsed(ctx context.Context, serviceInput *pb.Input) (*pb.FloatOutput, error) {
	res, err := g.Impl.PercentOfUsed(serviceInput)
	return &pb.FloatOutput{
		Value: res.Value,
		Error: res.Error,
	}, err
}
