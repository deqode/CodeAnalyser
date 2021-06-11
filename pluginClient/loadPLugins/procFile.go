package loadPLugins

import (
	"code-analyser/GlobalFiles"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type ProcFileClient struct {
	Methods *GlobalFiles.ProcFile
	Client  *plugin.Client
}

func (receiver *ProcFileClient) Load(pluginPath string) {
	methods, client := pluginClient.CreateProcFileClient(utils.CallPluginCommand(pluginPath))
	receiver.Client = client
	receiver.Methods = &methods
}
