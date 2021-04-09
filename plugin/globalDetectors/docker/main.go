package main

import (
	"code-analyser/pluginClient"
	dockerFile "code-analyser/pluginClient/docker"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

//DockerFile is a plugin
type DockerFile struct {
}

func (d DockerFile) DetectDockerFiles(inputString *pb.ServiceInputString) (*pb.ServiceOutputDockerFile, error) {
	return &pb.ServiceOutputDockerFile{
		Error: nil,
		DockerFile: []*global.DockerFileOutput{
			{
				Used: false,
				DockerFiles: []*global.DockerFile{
					{
						FileName: "Docker_Compose",
						DockerCommands: []*global.Command{
							{
								Command: "wow",
								Args:    []string{"nice", "good"},
							},
						},
					},
				},
			},
		},
	}, nil
}


func (d DockerFile) DetectDockerComposeFiles(inputString *pb.ServiceInputString) (*pb.ServiceOutputDockerComposeFile, error) {
	return &pb.ServiceOutputDockerComposeFile{
		Error: nil,
		DockerComposeFile: []*global.DockerComposeFileOutput{
			{
				Used: false,
				DockerComposeFiles:  []*global.DockerComposeFile{
					{
						FileName: "Docker_Compose",
						DockerComposeCommands: []*global.Command{
							{
								Command: "compose",
								Args:    []string{"nice", "good"},
							},
						},
					},
				},
			},
		},
	}, nil}


func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserDockerFile: &dockerFile.GRPCPlugin{
				Impl: &DockerFile{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
