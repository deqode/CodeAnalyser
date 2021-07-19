package loadPLugins

import (
	"code-analyser/GlobalFiles"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/helpers"
	gloabl "code-analyser/protos/pb/output/globalFiles"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

//MakeFilePlugin contains Methods, Client object of this plugin and Setting for logger related info
type MakeFilePlugin struct {
	Methods GlobalFiles.Makefile
	Client  *plugin.Client
	Setting *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugin *MakeFilePlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("makefile plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreateMakeFileClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("makefile plugin client created successfully")
}

//Run it runs plugin(execute methods of plugin)
func (plugin *MakeFilePlugin) Run(ctx context.Context, projectRootPath string) (*gloabl.MakeFile, error) {
	plugin.Setting.Logger.Debug("makefile plugin execution started")

	plugin.Setting.Logger.Debug("makefile detection started")
	makeFile, err := plugin.Methods.Detect(&helpers.StringInput{Value: projectRootPath})
	if err != nil {
		return nil, err
	}
	plugin.Setting.Logger.Debug("makefile detection completed")

	plugin.Setting.Logger.Debug("makefile plugin execution completed")
	return makeFile, nil
}
