package makeFile

import (
	"code-analyser/GlobalFiles"
	pb "code-analyser/protos/pb/helpers"
	global "code-analyser/protos/pb/output/globalFiles"
	"code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client plugin.MakeFileClient
}

//Detect method call on client side over gRPC
func (g *GRPCClient) Detect(input *pb.StringInput) (*global.MakeFile, error) {
	res, err := g.Client.Detect(context.Background(), input)
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl GlobalFiles.Makefile
}

//Detect will detect makefile, commands of makefile and usage
func (g *GRPCServer) Detect(ctx context.Context, input *pb.StringInput) (*global.MakeFile, error) {
	res, err := g.Impl.Detect(input)
	return res, err
}
