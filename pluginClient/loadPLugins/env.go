package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type EnvPlugin struct {
	Methods *interfaces.Env
	Client *plugin.Client
}

func (receiver EnvPlugin) name() {
    methods,client:=pluginClient.CreateEnvClient(utils.CallPluginCommand(plugin))
}