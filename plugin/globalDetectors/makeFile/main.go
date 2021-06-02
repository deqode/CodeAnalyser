package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/makeFile"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
	"os"
)

//MakeFile is a plugin
type MakeFile struct {
}

//TODO impolement logic for detection and path of makefile

// Detect it returns all detectors
func (m MakeFile) Detect(path *pb.ServiceInputString) (*pb.ServiceOutputMakeFile, error) {
	makeFileOutput := &pb.ServiceOutputMakeFile{
		Error: nil,
	}
	if _, err := os.Stat(path.Value + "/Makefile"); !os.IsNotExist(err) {
		makeFileOutput.MakeFile = &global.MakefileOutput{
			Used:     true,
			FilePath: path.Value + "/Makefile",
		}
	}
	return makeFileOutput, nil
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
