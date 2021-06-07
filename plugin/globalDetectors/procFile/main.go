package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/procFile"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	"github.com/chrismytton/procfile"
	"github.com/hashicorp/go-plugin"
	"io/ioutil"
	"os"
)

//ProcFile is a plugin
type ProcFile struct {
}

//TODO add path of procfile and detection logic

// Detect it returns all procfile
func (m ProcFile) Detect(path *pb.StringInput) (*pb.ProcFileOutput, error) {
	procFileOutput := &pb.ProcFileOutput{
		Error: nil,
	}
	if _, err := os.Stat(path.Value + "/Procfile"); !os.IsNotExist(err) {
		fileData, err := ioutil.ReadFile(path.Value + "/Procfile") // just pass the file name
		if err != nil {
			procFileOutput.Error = &pb.Error{
				Message: "not able to parse procfile",
				Cause:   path.Value + "/Procfile",
			}
		}
		procFileOutput.ProcFile = &global.ProcFile{
			Used:     true,
			FilePath: path.Value + "/Procfile",
			Commands: map[string]*global.Command{},
		}
		procList := procfile.Parse(string(fileData)) // convert content to a 'string'
		for name, command := range procList {
			procFileOutput.ProcFile.Commands[name] = &global.Command{
				Command: command.Command,
				Args:    command.Arguments,
			}
		}
	}
	return procFileOutput, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.ProcFile: &procFile.GRPCPlugin{
				Impl: &ProcFile{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
