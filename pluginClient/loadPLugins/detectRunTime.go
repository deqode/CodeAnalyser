package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type DetectRunTimePlugin struct {
	Methods *interfaces.DetectRunTime
	Client  *plugin.Client
}

func (receiver *DetectRunTimePlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateDetectRuntimeClient(utils.CallPluginCommand(yamlFile.Command))
	receiver.Client = client
	receiver.Methods = &methods
}
