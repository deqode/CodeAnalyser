package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/output/languageSpecific"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

type CommandsPlugin struct {
	Methods interfaces.Commands
	Client  *plugin.Client
}

func (plugin *CommandsPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateCommandsClient(utils.CallPluginCommand(yamlFile.Command))
}

func (plugin *CommandsPlugin) Run(ctx context.Context, projectRootPath string) (*languageSpecific.Commands, error) {

	return nil, nil
}
