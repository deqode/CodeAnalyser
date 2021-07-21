package buildDirectory

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	"context"
)

//GRPCClient is an implementation of BuildDirectory that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client pb.BuildDirectoryClient
}

//Detect method call on client side over gRPC
func (g *GRPCClient) Detect(input *helpers.Input) (*languageSpecific.BuildDirectoryOutput, error) {
	res, err := g.Client.Detect(context.Background(), &helpers.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl interfaces.BuildDirectory
}

//Detect method will detect build directories
func (g *GRPCServer) Detect(ctx context.Context, input *helpers.Input) (*languageSpecific.BuildDirectoryOutput, error) {
	res, err := g.Impl.Detect(input)
	return &languageSpecific.BuildDirectoryOutput{
		Used:           res.Used,
		BuildDirectory: res.BuildDirectory,
	}, err
}
