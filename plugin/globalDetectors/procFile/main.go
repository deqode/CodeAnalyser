package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/procFile"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
)

//ProcFile is a plugin
type ProcFile struct {
}

// Detect it returns all procfile
func (m ProcFile) Detect(inputString *pb.ServiceInputString) (*pb.ServiceOutputProcFile, error) {

	return &pb.ServiceOutputProcFile{
		Error: nil,
		ProcFile:  []*global.ProcFileOutput{{
			Used: true,
			ProcFiles: []*global.ProcFile{
				{
					FileName: "MakeFile",
					ProcCommands: []*global.Command{
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
			pluginClient.PluginDispenserProcFile: &procFile.GRPCPlugin{
				Impl: &ProcFile{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
