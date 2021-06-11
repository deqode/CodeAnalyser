package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type CommandsClient struct {
	Methods *interfaces.Commands
	Client  *plugin.Client
}

func (receiver *CommandsClient) Load(pluginPath string) {
	methods, client := pluginClient.CreateCommandsClient(utils.CallPluginCommand(pluginPath))
	receiver.Client = client
	receiver.Methods = &methods
}
