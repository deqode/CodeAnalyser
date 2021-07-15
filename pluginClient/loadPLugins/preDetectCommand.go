package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/helpers"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"errors"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

type PreDetectCommandsPlugin struct {
	Methods interfaces.PreDetectCommands
	Client  *plugin.Client
	Setting *utils.Setting
}

func (plugin *PreDetectCommandsPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("preDetectCommands plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreatePreDetectCommandClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("preDetectCommands plugin client created successfully")
}

func (plugin *PreDetectCommandsPlugin) Run(ctx context.Context,runTimeVersion, projectRootPath string) error {
	plugin.Setting.Logger.Debug("preDetectCommands plugin execution started")

	plugin.Setting.Logger.Debug("preDetect method execution started")
	response, err := plugin.Methods.RunPreDetect(&helpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	})
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Message + ", cause:- " + response.Error.Cause)
	}
	plugin.Setting.Logger.Info("pre detection command executed successfully")

	plugin.Setting.Logger.Debug("preDetectCommands plugin execution completed")
	return nil
}