package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/procFile"
	pb "code-analyser/protos/pb/helpers"
	globalPB "code-analyser/protos/pb/output/globalFiles"
	"github.com/chrismytton/procfile"
	"github.com/hashicorp/go-plugin"
	"io/ioutil"
	"os"
)

//ProcFile is a go plugin
type ProcFile struct {
}

//TODO add path of procfile and detection logic

// Detect it returns all procfile
func (m ProcFile) Detect(path *pb.StringInput) (*globalPB.ProcFile, error) {
	procFileOutput := &globalPB.ProcFile{
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

		procFileOutput.Used = true
		procFileOutput.FilePath = path.Value + "/Procfile"
		procFileOutput.Commands = map[string]*pb.Command{}

		procList := procfile.Parse(string(fileData)) // convert content to a 'string'
		for name, command := range procList {
			procFileOutput.Commands[name] = &pb.Command{
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
