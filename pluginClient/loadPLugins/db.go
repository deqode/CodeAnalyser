package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type DbClient struct {
	Methods *interfaces.Db
	Client  *plugin.Client
}

func (receiver *DbClient) Load(pluginPath string) {
	methods, client := pluginClient.CreateDbClient(utils.CallPluginCommand(pluginPath))
	receiver.Client = client
	receiver.Methods = &methods
}
