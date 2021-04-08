package buildDirectory

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/plugin"
	"context"
)

type GRPCClient struct {
	Client pb.BuildDirectoryClient
}

func (g *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputStringMap, error) {
	res, err := g.Client.Detect(context.Background(), &pb.ServiceInput{
		RuntimeVersion: input.RuntimeVersion,
		Root:           input.Root,
	})
	return res, err
}

type GRPCServer struct {
	Impl interfaces.BuildDirectory
}

func (g *GRPCServer) Detect(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputStringMap, error) {
	res, err := g.Impl.Detect(input)
	return &pb.ServiceOutputStringMap{
		Value: res.Value,
		Error: res.Error,
	}, err
}
