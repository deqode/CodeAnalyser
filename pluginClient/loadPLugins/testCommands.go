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

type TestCommandPlugin struct {
	Methods interfaces.TestCasesRunCommands
	Client  *plugin.Client
	Setting *utils.Setting
}

func (plugin *TestCommandPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateTestCaseCommandClient(utils.CallPluginCommand(yamlFile.Command))
}

func (plugin *TestCommandPlugin) Run(ctx context.Context,runTimeVersion, projectRootPath string) (*languagePB.TestCasesCommand, error) {
	commands, err := plugin.Methods.Detect(&helpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	})
	if err != nil {
		return nil, err
	}
	return commands, nil
}
