package library

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
	"code-analyser/utils"
	"errors"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client plugin.LibraryClient
}

func (g *GRPCClient) Detect(input *pb.Input) (*plugin.LibraryType, error) {
	if input != nil {
		res, err := g.Client.Detect(context.Background(), &pb.Input{
			RuntimeVersion: input.RuntimeVersion,
			RootPath:           input.RootPath,
		})
		return res, err
	}
	utils.Logger("service input nil found in library grpc detect method")
	return nil, errors.New("service input nil found")
}

func (g *GRPCClient) IsUsed(serviceInput *pb.Input) (*pb.BoolOutput, error) {
	res, err := g.Client.IsUsed(context.Background(), &pb.Input{
		RuntimeVersion: serviceInput.RuntimeVersion,
		RootPath:           serviceInput.RootPath,
	})
	return res, err
}

func (g *GRPCClient) PercentOfUsed(serviceInput *pb.Input) (*pb.FloatOutput, error) {
	res, err := g.Client.PercentOfUsed(context.Background(), &pb.Input{
		RuntimeVersion: serviceInput.RuntimeVersion,
		RootPath:           serviceInput.RootPath,
	})
	return res, err
}

type GRPCServer struct {
	Impl interfaces.Library
}

func (g *GRPCServer) Detect(ctx context.Context, serviceInput *pb.Input) (*plugin.LibraryType, error) {
	res, err := g.Impl.Detect(serviceInput)
	return &plugin.LibraryType{
		Type:  res.Type,
		Value: res.Value,
		Error: res.Error,
	}, err
}

func (g *GRPCServer) IsUsed(ctx context.Context, serviceInput *pb.Input) (*pb.BoolOutput, error) {
	res, err := g.Impl.IsUsed(serviceInput)
	return &pb.BoolOutput{
		Value: res.Value,
		Error: res.Error,
	}, err
}

func (g *GRPCServer) PercentOfUsed(ctx context.Context, serviceInput *pb.Input) (*pb.FloatOutput, error) {
	res, err := g.Impl.PercentOfUsed(serviceInput)
	return &pb.FloatOutput{
		Value: res.Value,
		Error: res.Error,
	}, err
}
