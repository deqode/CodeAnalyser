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

type EnvPlugin struct {
	Methods interfaces.Env
	Client  *plugin.Client
	Setting *utils.Setting
}

func (plugin *EnvPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateEnvClient(utils.CallPluginCommand(yamlFile.Command))
}

func (plugin *EnvPlugin) Run(ctx context.Context, runTimeVersion, projectRootPath string) (*languageSpecific.Envs, error) {
	envVariables, err := plugin.Methods.Detect(&helpers.Input{RootPath: projectRootPath, RuntimeVersion: runTimeVersion})
	if err != nil {
		return nil, err
	}
	return envVariables, nil
}
