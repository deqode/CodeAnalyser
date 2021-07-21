package dockerFile

import (
	"code-analyser/GlobalFiles"
	pb "code-analyser/protos/pb/helpers"
	global "code-analyser/protos/pb/output/globalFiles"
	"code-analyser/protos/pb/plugin"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client plugin.DockerFileClient
}

//DetectDockerFile method call on client side over gRPC
func (g *GRPCClient) DetectDockerFile(input *pb.StringInput) (*global.DockerFile, error) {
	res, err := g.Client.DetectDockerFiles(context.Background(), input)
	return res, err
}

//DetectDockerComposeFile method call on client side over gRPC
func (g *GRPCClient) DetectDockerComposeFile(input *pb.StringInput) (*global.DockerCompose, error) {
	res, err := g.Client.DetectDockerComposeFiles(context.Background(), input)
	return res, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	Impl GlobalFiles.DockerFile
}

//DetectDockerFiles will detect docker file and their commands
func (g *GRPCServer) DetectDockerFiles(ctx context.Context, input *pb.StringInput) (*global.DockerFile, error) {
	res, err := g.Impl.DetectDockerFile(input)
	return res, err
}

//DetectDockerComposeFiles will detect docker-compose file and their commands
func (g *GRPCServer) DetectDockerComposeFiles(ctx context.Context, input *pb.StringInput) (*global.DockerCompose, error) {
	res, err := g.Impl.DetectDockerComposeFile(input)
	return res, err
}
