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

//CommandsPlugin contains Methods and Client object of this plugin,
//Setting for logger related info
type CommandsPlugin struct {
	Methods interfaces.Commands
	Client  *plugin.Client
	Setting *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugin *CommandsPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("commands plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreateCommandsClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("commands plugin client created successfully")
}

//Run it runs plugin(execute methods of plugin)
func (plugin *CommandsPlugin) Run(ctx context.Context, projectRootPath string) (*languageSpecific.Commands, error) {
	plugin.Setting.Logger.Debug("commands plugin execution started")

	rootPathInput := &helpers.StringInput{Value: projectRootPath}
	commands := &languageSpecific.Commands{}

	plugin.Setting.Logger.Debug("makefile detection started")
	adHocScripts, err := plugin.Methods.DetectAdHocScripts(rootPathInput)
	if err != nil {
		return commands, err
	}
	commands.AdHocScripts = adHocScripts
	plugin.Setting.Logger.Debug("makefile detection completed")

	plugin.Setting.Logger.Debug("seedCommands detection started")
	seedCommands, err := plugin.Methods.DetectSeedCommands(rootPathInput)
	if err != nil {
		return commands, err
	}
	commands.SeedCommands = seedCommands
	plugin.Setting.Logger.Debug("seedCommands detection completed")

	plugin.Setting.Logger.Debug("buildCommands detection started")
	buildCommands, err := plugin.Methods.DetectBuildCommands(rootPathInput)
	if err != nil {
		return commands, err
	}
	commands.BuildCommands = buildCommands
	plugin.Setting.Logger.Debug("buildCommands detection completed")

	plugin.Setting.Logger.Debug("migrationCommands detection started")
	migrationCommands, err := plugin.Methods.DetectMigrationCommands(rootPathInput)
	if err != nil {
		return commands, err
	}
	commands.MigrationCommands = migrationCommands
	plugin.Setting.Logger.Debug("migrationCommands detection completed")

	plugin.Setting.Logger.Debug("startupCommands detection started")
	startUpCommands, err := plugin.Methods.DetectStartUpCommands(rootPathInput)
	if err != nil {
		return commands, err
	}
	commands.StartUpCommands = startUpCommands
	plugin.Setting.Logger.Debug("startupCommands detection completed")

	plugin.Setting.Logger.Debug("commands plugin execution completed")
	return commands, nil
}
