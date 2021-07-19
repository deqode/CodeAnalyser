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

//EnvPlugin contains Methods, Client object of this plugin,
//Setting for logger related info
type EnvPlugin struct {
	Methods interfaces.Env
	Client  *plugin.Client
	Setting *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugin *EnvPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("env plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreateEnvClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("env plugin client created successfully")
}

//Run it runs plugin(execute methods of plugin)
func (plugin *EnvPlugin) Run(ctx context.Context, runTimeVersion, projectRootPath string) (*languageSpecific.Envs, error) {
	plugin.Setting.Logger.Debug("env plugin methods execution started")

	plugin.Setting.Logger.Debug("environment variables detection started")
	envVariables, err := plugin.Methods.Detect(&helpers.Input{RootPath: projectRootPath, RuntimeVersion: runTimeVersion})
	if err != nil {
		return nil, err
	}
	plugin.Setting.Logger.Debug("environment variables detected")

	plugin.Setting.Logger.Debug("env plugin methods execution completed")
	return envVariables, nil
}
