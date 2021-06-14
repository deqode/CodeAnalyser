package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type TestCommandPlugin struct {
	Methods *interfaces.TestCasesRunCommands
	Client  *plugin.Client
}

func (receiver *TestCommandPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateTestCaseCommandClient(utils.CallPluginCommand(yamlFile.Command))
	receiver.Client = client
	receiver.Methods = &methods
}

