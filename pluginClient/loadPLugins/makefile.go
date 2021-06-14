package loadPLugins

import (
	"code-analyser/GlobalFiles"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type MakeFilePlugin struct {
	Methods *GlobalFiles.Makefile
	Client  *plugin.Client
}

func (receiver *MakeFilePlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateMakeFileClient(utils.CallPluginCommand(yamlFile.Command))
	receiver.Client = client
	receiver.Methods = &methods
}

