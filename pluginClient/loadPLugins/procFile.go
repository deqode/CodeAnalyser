package loadPLugins

import (
	"code-analyser/GlobalFiles"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/helpers"
	global "code-analyser/protos/pb/output/globalFiles"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

type ProcFilePlugin struct {
	Methods GlobalFiles.ProcFile
	Client  *plugin.Client
	Setting *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugin *ProcFilePlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("procfile plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreateProcFileClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("procfile plugin client created successfully")
}

//Run it runs plugin(execute methods of plugin)
func (plugin *ProcFilePlugin) Run(ctx context.Context, projectRootPath string) (*global.ProcFile, error) {
	plugin.Setting.Logger.Debug("procfile plugin execution started")

	plugin.Setting.Logger.Debug("procfile detection started")
	procFile, err := plugin.Methods.Detect(&helpers.StringInput{Value: projectRootPath})
	if err != nil {
		return nil, err
	}
	plugin.Setting.Logger.Debug("procfile detection completed")

	plugin.Setting.Logger.Debug("procfile plugin execution completed")
	return procFile, err
}
