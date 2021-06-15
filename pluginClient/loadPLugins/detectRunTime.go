package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/helpers"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type DetectRunTimePlugin struct {
	Methods interfaces.DetectRunTime
	Client  *plugin.Client
}

func (plugin *DetectRunTimePlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateDetectRuntimeClient(utils.CallPluginCommand(yamlFile.Command))
}

func (plugin *DetectRunTimePlugin) Run(projectRootPath string) (*helpers.StringOutput, error) {
	languageVersion, err := plugin.Methods.Detect(&helpers.StringInput{Value: projectRootPath})
	if err != nil {
		return nil, err
	}
	return languageVersion, nil
}
