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
	"log"
)

type PreDetectCommandsPlugin struct {
	Methods interfaces.PreDetectCommands
	Client  *plugin.Client
	Setting *utils.Setting
}

func (plugin *PreDetectCommandsPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreatePreDetectCommandClient(utils.CallPluginCommand(yamlFile.Command))
}

func (plugin *PreDetectCommandsPlugin) Run(ctx context.Context,runTimeVersion, projectRootPath string) error {
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
	log.Println("pre detection command executed successfully")
	return nil
}
