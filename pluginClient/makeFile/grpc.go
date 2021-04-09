package makeFile

import (
	"code-analyser/GlobalFiles"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client pb.MakeFileServiceClient
}

func (g *GRPCClient) Detect(input *pb.ServiceInputString) (*pb.ServiceOutputMakeFile, error) {
	res, err := g.Client.Detect(context.Background(), input)
	return res, err
}

type GRPCServer struct {
	Impl GlobalFiles.Makefiles
}

func (g *GRPCServer) Detect(ctx context.Context, input *pb.ServiceInputString) (*pb.ServiceOutputMakeFile, error) {
	res, err := g.Impl.Detect(input)
	return res, err
}
