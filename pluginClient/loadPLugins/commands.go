package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	helpersPb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

type CommandsPlugin struct {
	Methods *interfaces.Commands
	Client  *plugin.Client
}

func (plugin *CommandsPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateCommandsClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Client = client
	plugin.Methods = &methods
}

func (plugin *CommandsPlugin) Run(ctx context.Context, projectRootPath string) (*languageSpecific.Commands, error) {
	commands := languageSpecific.Commands{}

	rootPathInput := &helpersPb.StringInput{
		Value: projectRootPath,
	}

	plugin.Methods

	adHocScripts, err := plugin.Methods
	if err != nil || adHocScripts.Error != nil {
		utils.Logger(err, adHocScripts)
		return &commands
	}
	commands.AdHocScripts = adHocScripts
}
