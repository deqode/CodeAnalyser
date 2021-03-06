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

type BuildDirectoryPlugin struct {
	Methods interfaces.BuildDirectory
	Client  *plugin.Client
	Setting *utils.Setting
}

func (plugin *BuildDirectoryPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("buildDirectory plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreateBuildDirectoryClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("buildDirectory plugin client created successfully")
}

func (plugin BuildDirectoryPlugin) Run(ctx context.Context,runTimeVersion, projectRootPath string) (*languageSpecific.BuildDirectoryOutput, error) {
	buildDirectory, err := plugin.Methods.Detect(&helpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	})
	if err != nil {
		return nil, err
	}
	return buildDirectory, nil
}
