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
func (m MakeFile) Detect(path *pb.StringInput) (*pb.MakeFileOutput, error) {
	makeFileOutput := &pb.MakeFileOutput{
		Error: nil,
	}
	if _, err := os.Stat(path.Value + "/Makefile"); !os.IsNotExist(err) {
		makeFileOutput.Makefile = &global.MakeFile{
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
			pluginClient.MakeFile: &makeFile.GRPCPlugin{
				Impl: &MakeFile{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
