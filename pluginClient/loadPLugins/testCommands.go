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

//TestCommandPlugin contains Methods, Client object of this plugin and Setting for logger related info
type TestCommandPlugin struct {
	Methods interfaces.TestCasesRunCommands
	Client  *plugin.Client
	Setting *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugin *TestCommandPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("test commands plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreateTestCaseCommandClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("test commands plugin client creation completed")
}

//Run it runs plugin(execute methods of plugin)
func (plugin *TestCommandPlugin) Run(ctx context.Context, runTimeVersion, projectRootPath string) (*languagePB.TestCasesCommand, error) {
	plugin.Setting.Logger.Debug("test commands plugin methods execution started")

	plugin.Setting.Logger.Debug("test commands detection started")
	commands, err := plugin.Methods.Detect(&helpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	})
	if err != nil {
		return nil, err
	}
	plugin.Setting.Logger.Debug("test commands detection started")

	plugin.Setting.Logger.Debug("test commands plugin methods executed successfully")
	return commands, nil
}
