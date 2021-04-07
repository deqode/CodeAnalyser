package library

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient/pb"
	"code-analyser/utils"
	"errors"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client pb.LibraryServiceClient
}

func (g *GRPCClient) Detect(serviceInput *pb.ServiceInput) (*pb.ServiceOutputLibraryType, error) {
	if serviceInput != nil {
		res, err := g.Client.Detect(context.Background(), &pb.ServiceInput{
			RuntimeVersion: serviceInput.RuntimeVersion,
			Root:           serviceInput.Root,
		})
		return res, err
	}
	utils.Logger("service input nil found in library grpc detect method")
	return nil, errors.New("service input nil found")
}

func (g *GRPCClient) IsUsed(serviceInput *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := g.Client.IsUsed(context.Background(), &pb.ServiceInput{
		RuntimeVersion: serviceInput.RuntimeVersion,
		Root:           serviceInput.Root,
	})
	return res, err
}

func (g *GRPCClient) PercentOfUsed(serviceInput *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	res, err := g.Client.PercentOfUsed(context.Background(), &pb.ServiceInput{
		RuntimeVersion: serviceInput.RuntimeVersion,
		Root:           serviceInput.Root,
	})
	return res, err
}

type GRPCServer struct {
	Impl interfaces.Library
}

func (g *GRPCServer) Detect(ctx context.Context, serviceInput *pb.ServiceInput) (*pb.ServiceOutputLibraryType, error) {
	res, err := g.Impl.Detect(serviceInput)
	return &pb.ServiceOutputLibraryType{
		Type:  res.Type,
		Value: res.Value,
		Error: res.Error,
	}, err
}

func (g *GRPCServer) IsUsed(ctx context.Context, serviceInput *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	res, err := g.Impl.IsUsed(serviceInput)
	return &pb.ServiceOutputBool{
		Value: res.Value,
		Error: res.Error,
	}, err
}

func (g *GRPCServer) PercentOfUsed(ctx context.Context, serviceInput *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	res, err := g.Impl.PercentOfUsed(serviceInput)
	return &pb.ServiceOutputFloat{
		Value: res.Value,
		Error: res.Error,
	}, err
}
