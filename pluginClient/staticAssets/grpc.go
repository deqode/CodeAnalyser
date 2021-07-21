package staticAssets

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	"code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client plugin.StaticAssetsClient
}

//Detect method call on client side over gRPC
func (g *GRPCClient) Detect(input *pb.Input) (*languageSpecific.StaticAssetsOutput, error) {
	res, err := g.Client.Detect(context.Background(), input)
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl interfaces.StaticAssets
}

//Detect will detect static assets
func (g *GRPCServer) Detect(ctx context.Context, input *pb.Input) (*languageSpecific.StaticAssetsOutput, error) {
	res, err := g.Impl.Detect(input)
	return res, err
}
