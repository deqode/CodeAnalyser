package loadPLugins

import (
	"code-analyser/GlobalFiles"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/helpers"
	global "code-analyser/protos/pb/output/globalFiles"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type ProcFilePlugin struct {
	Methods GlobalFiles.ProcFile
	Client  *plugin.Client
}

func (plugin *ProcFilePlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateProcFileClient(utils.CallPluginCommand(yamlFile.Command))
}

func (plugin *ProcFilePlugin) Run(projectRootPath string) (*global.ProcFile, error) {
	procFile, err := plugin.Methods.Detect(&helpers.StringInput{Value: projectRootPath})
	if err != nil {
		return nil, err
	}
	return procFile, err
}
