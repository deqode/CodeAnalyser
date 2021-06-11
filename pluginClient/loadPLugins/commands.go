package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type CommandsPlugin struct {
	Methods *interfaces.Commands
	Client  *plugin.Client
}

func (receiver *CommandsPlugin) Load(pluginPath string) {
	methods, client := pluginClient.CreateCommandsClient(utils.CallPluginCommand(pluginPath))
	receiver.Client = client
	receiver.Methods = &methods
}
