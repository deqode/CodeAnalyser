package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type FrameworkPlugin struct {
	Methods *interfaces.Framework
	Client  *plugin.Client
}

func (receiver *FrameworkPlugin) Load(pluginPath string) {
	methods, client := pluginClient.CreateFrameworkClient(utils.CallPluginCommand(pluginPath))
	receiver.Client = client
	receiver.Methods = &methods
}

func (receiver FrameworkPlugin) Run()  {

}
