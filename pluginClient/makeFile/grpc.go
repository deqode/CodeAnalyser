package makeFile

import (
	"code-analyser/GlobalFiles"
	pb "code-analyser/protos/pb/helpers"
	global "code-analyser/protos/pb/output/globalFiles"
	"code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client plugin.MakeFileClient
}

func (g *GRPCClient) Detect(input *pb.StringInput) (*global.MakeFile, error) {
	res, err := g.Client.Detect(context.Background(), input)
	return res, err
}

type GRPCServer struct {
	Impl GlobalFiles.Makefile
}

func (g *GRPCServer) Detect(ctx context.Context, input *pb.StringInput) (*global.MakeFile, error) {
	res, err := g.Impl.Detect(input)
	return res, err
}
