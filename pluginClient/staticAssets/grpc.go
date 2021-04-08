package staticAssets

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client pb.StaticAssetsServiceClient
}

func (g *GRPCClient) Detect(input *pb.ServiceInput) (*pb.ServiceOutputStaticAssets, error) {
	res, err := g.Client.Detect(context.Background(), input)
	return res, err
}

type GRPCServer struct {
	Impl interfaces.StaticAssets
}

func (g *GRPCServer) Detect(ctx context.Context, input *pb.ServiceInput) (*pb.ServiceOutputStaticAssets, error) {
	res, err := g.Impl.Detect(input)
	return res, err
}
