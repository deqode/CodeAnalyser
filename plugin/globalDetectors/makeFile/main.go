package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/makeFile"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

//MakeFile is a plugin
type MakeFile struct {
}

// Detect it returns all detectors
func (m MakeFile) Detect(inputString *pb.ServiceInputString) (*pb.ServiceOutputMakeFile, error) {

	return &pb.ServiceOutputMakeFile{
		Error: nil,
		MakeFile: []*global.MakefileOutput{{
			Used: true,
			MakeFiles: []*global.MakeFile{
				{
					FileName: "MakeFile",
					MakeCommands: []*global.Command{
						{
							Command: "run",
							Args:    []string{"run", "make"},
						},
					},
				},
			},
		}},
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserMakeFile: &makeFile.GRPCPlugin{
				Impl: &MakeFile{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
