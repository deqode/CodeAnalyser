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

type MakeFilePlugin struct {
	Methods GlobalFiles.Makefile
	Client  *plugin.Client
}

func (plugin *MakeFilePlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateMakeFileClient(utils.CallPluginCommand(yamlFile.Command))
}

func (plugin *MakeFilePlugin) Run(ctx context.Context,projectRootPath string) (*gloabl.MakeFile, error) {
	makeFile, err := plugin.Methods.Detect(&helpers.StringInput{Value: projectRootPath})
	if err != nil {
		return nil, err
	}
	return makeFile, nil
}
