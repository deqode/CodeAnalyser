package loadPLugins

import (
	"code-analyser/GlobalFiles"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type ProcFileClient struct {
	Methods *GlobalFiles.ProcFile
	Client  *plugin.Client
}

func (receiver *ProcFileClient) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateProcFileClient(utils.CallPluginCommand(yamlFile.Command))
	receiver.Client = client
	receiver.Methods = &methods
}
