package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/makeFile"
	pb "code-analyser/protos/pb/helpers"
	globalPB "code-analyser/protos/pb/output/globalFiles"
	"github.com/hashicorp/go-plugin"
	"os"
)

//MakeFile is a go plugin
type MakeFile struct {
}

//Detect is a MakeFile plugin method, it will detect Makefile and their commands
func (m MakeFile) Detect(path *pb.StringInput) (*globalPB.MakeFile, error) {

	makeFileOutput := &globalPB.MakeFile{
		Error: nil,
	}

	if _, err := os.Stat(path.Value + "/Makefile"); !os.IsNotExist(err) {
		makeFileOutput.Used = true
		makeFileOutput.FilePath = path.Value + "/Makefile"
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
