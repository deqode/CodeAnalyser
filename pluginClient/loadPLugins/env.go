package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type EnvPlugin struct {
	Methods interfaces.Env
	Client  *plugin.Client
}

func (plugin *EnvPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateEnvClient(utils.CallPluginCommand(yamlFile.Command))

}
