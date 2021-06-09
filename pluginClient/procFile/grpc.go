package procFile

import (
	"code-analyser/GlobalFiles"
	pb "code-analyser/protos/pb/helpers"
	gloabl "code-analyser/protos/pb/output/globalFiles"
	"code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client plugin.ProcFileClient
}

func (g *GRPCClient) Detect(projectRootPath *pb.StringInput) (*gloabl.ProcFile, error) {
	res, err := g.Client.Detect(context.Background(), projectRootPath)
	return res, err
}

type GRPCServer struct {
	Impl GlobalFiles.ProcFile
}

func (g *GRPCServer) Detect(ctx context.Context, projectRootPath *pb.StringInput) (*gloabl.ProcFile, error) {
	res, err := g.Impl.Detect(projectRootPath)
	return res, err
}
