package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type DetectRunTimePlugin struct {
	Methods interfaces.DetectRunTime
	Client  *plugin.Client
}

func (plugin *DetectRunTimePlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateDetectRuntimeClient(utils.CallPluginCommand(yamlFile.Command))
}
