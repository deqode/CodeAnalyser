package buildDirectory

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	"context"
)

type GRPCClient struct {
	Client pb.BuildDirectoryClient
}

func (g *GRPCClient) Detect(input *helpers.Input) (*languageSpecific.BuildDirectoryOutput, error) {
	res, err := g.Client.Detect(context.Background(), &helpers.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return res, err
}

type GRPCServer struct {
	Impl interfaces.BuildDirectory
}

func (g *GRPCServer) Detect(ctx context.Context, input *helpers.Input) (*languageSpecific.BuildDirectoryOutput, error) {
	res, err := g.Impl.Detect(input)
	return &languageSpecific.BuildDirectoryOutput{
		Used:           res.Used,
		BuildDirectory: res.BuildDirectory,
	}, err
}
