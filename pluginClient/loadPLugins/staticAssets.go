package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/helpers"
	languagePB "code-analyser/protos/pb/output/languageSpecific"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type StaticAssetsPlugin struct {
	Methods interfaces.StaticAssets
	Client  *plugin.Client
}

func (plugin *StaticAssetsPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateStaticAssetsClient(utils.CallPluginCommand(yamlFile.Command))
}

func (plugin *StaticAssetsPlugin) Run(runTimeVersion, projectRootPath string) (*languagePB.StaticAssetsOutput, error) {
	staticAssets, err := plugin.Methods.Detect(&helpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	})
	if err != nil {
		return nil, err
	}
	return staticAssets, nil
}
