package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/helpers"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

type DetectRunTimePlugin struct {
	Methods interfaces.DetectRunTime
	Client  *plugin.Client
	Setting *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugin *DetectRunTimePlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("detectRuntime plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreateDetectRuntimeClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("detectRuntime plugin client created successfully")
}

//Run it runs plugin(execute methods of plugin)
func (plugin *DetectRunTimePlugin) Run(ctx context.Context, projectRootPath string) (*helpers.StringOutput, error) {
	plugin.Setting.Logger.Debug("detectRuntime plugin methods execution started")

	plugin.Setting.Logger.Debug("runtime detection started")
	languageVersion, err := plugin.Methods.Detect(&helpers.StringInput{Value: projectRootPath})
	if err != nil {
		return nil, err
	}
	plugin.Setting.Logger.Debug("runtime detected")

	plugin.Setting.Logger.Debug("detectRuntime plugin methods execution completed")
	return languageVersion, nil
}
