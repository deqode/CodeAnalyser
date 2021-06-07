package main

import (
	"code-analyser/pluginClient"
	dockerFile "code-analyser/pluginClient/docker"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
	"os"
)

//DockerFile is a plugin
type DockerFile struct {
}

//TODO add file path instead of file name  for docker and docker compose

func (d DockerFile) DetectDockerFile(path *pb.StringInput) (*pb.DockerFileOutput, error) {
	dockerFileOutput := &pb.DockerFileOutput{
		Error: nil,
	}
	if _, err := os.Stat(path.Value + "/Dockerfile"); !os.IsNotExist(err) {
		dockerFileOutput.DockerFile = &global.DockerFile{
			Used:     true,
			FilePath: path.Value + "/Dockerfile",
		}
	}
	return dockerFileOutput, nil
}

func (d DockerFile) DetectDockerComposeFile(path *pb.StringInput) (*pb.DockerComposeFileOutput, error) {
	dockerComposeOutput := &pb.DockerComposeFileOutput{
		Error: nil,
	}
	var err error
	if _, err = os.Stat(path.Value + "/docker-compose.yml"); !os.IsNotExist(err) {
		dockerComposeOutput.DockerComposeFile = &global.DockerCompose{
			FilePath: path.Value + "/docker-compose.yml",
			Used:     true,
		}
	}
	if _, err := os.Stat(path.Value + "/docker-compose.yaml"); !os.IsNotExist(err) {
		dockerComposeOutput.DockerComposeFile = &global.DockerCompose{
			FilePath: path.Value + "/docker-compose.yaml",
			Used:     true,
		}
	}
	return dockerComposeOutput, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.DockerFile: &dockerFile.GRPCPlugin{
				Impl: &DockerFile{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
