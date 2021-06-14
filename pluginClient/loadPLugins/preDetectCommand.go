package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type PreDetectCommandsPlugin struct {
	Methods *interfaces.PreDetectCommands
	Client  *plugin.Client
}

func (receiver *PreDetectCommandsPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreatePreDetectCommandClient(utils.CallPluginCommand(yamlFile.Command))
	receiver.Client = client
	receiver.Methods = &methods
}
