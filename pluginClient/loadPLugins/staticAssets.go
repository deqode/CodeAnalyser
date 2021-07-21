package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/helpers"
	languagePB "code-analyser/protos/pb/output/languageSpecific"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

//StaticAssetsPlugin contains Methods, Client object of this plugin and Setting for logger related info
type StaticAssetsPlugin struct {
	Methods interfaces.StaticAssets
	Client  *plugin.Client
	Setting *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugin *StaticAssetsPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("static assets plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreateStaticAssetsClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("static assets plugin client created successfully")
}

//Run it runs plugin(execute methods of plugin)
func (plugin *StaticAssetsPlugin) Run(ctx context.Context, runTimeVersion, projectRootPath string) (*languagePB.StaticAssetsOutput, error) {
	plugin.Setting.Logger.Debug("static assets plugin execution started")

	plugin.Setting.Logger.Debug("static assets detection started")
	staticAssets, err := plugin.Methods.Detect(&helpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	})
	if err != nil {
		return nil, err
	}
	plugin.Setting.Logger.Debug("static assets detection completed")

	plugin.Setting.Logger.Debug("static assets plugin execution completed")
	return staticAssets, nil
}
