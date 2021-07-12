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
}

func (plugin *BuildDirectoryPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateBuildDirectoryClient(utils.CallPluginCommand(yamlFile.Command))
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
