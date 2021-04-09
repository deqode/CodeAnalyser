package dockerFile

import (
	"code-analyser/GlobalFiles"
	pb "code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	Client pb.DockerFileServiceClient
}

func (g *GRPCClient) DetectDockerFiles(input *pb.ServiceInputString) (*pb.ServiceOutputDockerFile, error) {
	res, err := g.Client.DetectDockerFiles(context.Background(), input)
	return res, err
}
func (g *GRPCClient) DetectDockerComposeFiles(input *pb.ServiceInputString) (*pb.ServiceOutputDockerComposeFile, error) {
	res, err := g.Client.DetectDockerComposeFiles(context.Background(), input)
	return res, err
}


type GRPCServer struct {
	Impl GlobalFiles.DockerFile
}

func (g *GRPCServer) DetectDockerFiles(ctx context.Context, input *pb.ServiceInputString) (*pb.ServiceOutputDockerFile, error) {
	res, err := g.Impl.DetectDockerFiles(input)
	return res, err
}
func (g *GRPCServer) DetectDockerComposeFiles(ctx context.Context, input *pb.ServiceInputString) (*pb.ServiceOutputDockerComposeFile, error) {
	res, err := g.Impl.DetectDockerComposeFiles(input)
	return res, err
}

