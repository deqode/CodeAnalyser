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

type CommandsPlugin struct {
	Methods interfaces.Commands
	Client  *plugin.Client
	Setting *utils.Setting
}

func (plugin *CommandsPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("commands plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreateCommandsClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("commands plugin client created successfully")
}

func (plugin *CommandsPlugin) Run(ctx context.Context, projectRootPath string) (*languageSpecific.Commands, error) {

	rootPathInput := &helpers.StringInput{Value: projectRootPath}
	commands := &languageSpecific.Commands{}

	adHocScripts, err := plugin.Methods.DetectAdHocScripts(rootPathInput)
	if err != nil {
		return commands, err
	}
	commands.AdHocScripts = adHocScripts

	seedCommands, err := plugin.Methods.DetectSeedCommands(rootPathInput)
	if err != nil {
		return commands, err
	}
	commands.SeedCommands = seedCommands

	buildCommands, err := plugin.Methods.DetectBuildCommands(rootPathInput)
	if err != nil {
		return commands, err
	}
	commands.BuildCommands = buildCommands

	migrationCommands, err := plugin.Methods.DetectMigrationCommands(rootPathInput)
	if err != nil {
		return commands, err
	}
	commands.MigrationCommands = migrationCommands

	startUpCommands, err := plugin.Methods.DetectStartUpCommands(rootPathInput)
	if err != nil {
		return commands, err
	}
	commands.StartUpCommands = startUpCommands

	return commands, nil
}
