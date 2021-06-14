package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type StaticAssetsPlugin struct {
	Methods *interfaces.StaticAssets
	Client  *plugin.Client
}

func (receiver *StaticAssetsPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateStaticAssetsClient(utils.CallPluginCommand(yamlFile.Command))
	receiver.Client = client
	receiver.Methods = &methods
}

