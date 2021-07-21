package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

//BuildDirectoryPlugin contains Methods and Client object of this plugin,
//Setting for logger related info
type BuildDirectoryPlugin struct {
	Methods interfaces.BuildDirectory
	Client  *plugin.Client
	Setting *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugin *BuildDirectoryPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("buildDirectory plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreateBuildDirectoryClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("buildDirectory plugin client created successfully")
}

//Run it runs plugin(execute methods of plugin)
func (plugin BuildDirectoryPlugin) Run(ctx context.Context, runTimeVersion, projectRootPath string) (*languageSpecific.BuildDirectoryOutput, error) {
	plugin.Setting.Logger.Debug("buildDirectory plugin execution started")

	plugin.Setting.Logger.Debug("buildDirectory detection started")
	buildDirectory, err := plugin.Methods.Detect(&helpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	})
	if err != nil {
		return nil, err
	}
	plugin.Setting.Logger.Debug("buildDirectory detection completed")

	plugin.Setting.Logger.Debug("buildDirectory plugin execution completed")
	return buildDirectory, nil
}
