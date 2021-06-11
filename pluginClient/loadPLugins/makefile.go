package loadPLugins

import (
	"code-analyser/GlobalFiles"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type MakeFileClient struct {
	Methods *GlobalFiles.Makefile
	Client  *plugin.Client
}

func (receiver *MakeFileClient) Load(pluginPath string) {
	methods, client := pluginClient.CreateMakeFileClient(utils.CallPluginCommand(pluginPath))
	receiver.Client = client
	receiver.Methods = &methods
}

