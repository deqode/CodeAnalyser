package main

import (
	"code-analyser/pluginClient"
	dockerFile "code-analyser/pluginClient/docker"
	pb "code-analyser/protos/pb/helpers"
	global "code-analyser/protos/pb/output/globalFiles"
	"github.com/hashicorp/go-plugin"
	"os"
)

//DockerFile is a plugin
type DockerFile struct {
}

//TODO add file path instead of file name  for docker and docker compose

func (d DockerFile) DetectDockerFile(path *pb.StringInput) (*global.DockerFile, error) {
	dockerFileOutput := &global.DockerFile{
		Error: nil,
	}
	if _, err := os.Stat(path.Value + "/Dockerfile"); !os.IsNotExist(err) {
		dockerFileOutput.Used = true
		dockerFileOutput.FilePath = path.Value + "/Dockerfile"
	}
	return dockerFileOutput, nil
}

func (d DockerFile) DetectDockerComposeFile(path *pb.StringInput) (*global.DockerCompose, error) {
	dockerCompose := &global.DockerCompose{
		Error: nil,
	}
	var err error
	if _, err = os.Stat(path.Value + "/docker-compose.yml"); !os.IsNotExist(err) {
		dockerCompose.FilePath = path.Value + "/docker-compose.yml"
		dockerCompose.Used = true
	}
	if _, err := os.Stat(path.Value + "/docker-compose.yaml"); !os.IsNotExist(err) {
		dockerCompose.FilePath = path.Value + "/docker-compose.yaml"
		dockerCompose.Used = true
	}
	return dockerCompose, nil
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
