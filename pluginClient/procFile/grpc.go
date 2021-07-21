package procFile

import (
	"code-analyser/GlobalFiles"
	pb "code-analyser/protos/pb/helpers"
	gloabl "code-analyser/protos/pb/output/globalFiles"
	"code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client plugin.ProcFileClient
}

//Detect method call on client side over gRPC
func (g *GRPCClient) Detect(projectRootPath *pb.StringInput) (*gloabl.ProcFile, error) {
	res, err := g.Client.Detect(context.Background(), projectRootPath)
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl GlobalFiles.ProcFile
}

//Detect will detect is ProcFile, path and commands of ProcFile
func (g *GRPCServer) Detect(ctx context.Context, projectRootPath *pb.StringInput) (*gloabl.ProcFile, error) {
	res, err := g.Impl.Detect(projectRootPath)
	return res, err
}
