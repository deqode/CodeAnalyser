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

type DockerFilePlugin struct {
	Methods GlobalFiles.DockerFile
	Client  *plugin.Client
}

func (plugin *DockerFilePlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateDockerFileClient(utils.CallPluginCommand(yamlFile.Command))
}

func (plugin *DockerFilePlugin) Run(ctx context.Context, projectRootPath string) (*global.DockerFile, *global.DockerCompose, error) {
	dockerFile, err := plugin.Methods.DetectDockerFile(&helpers.StringInput{Value: projectRootPath})

	if err != nil {
		return nil, nil, err
	}

	dockerCompose, err := plugin.Methods.DetectDockerComposeFile(&helpers.StringInput{Value: projectRootPath})
	if err != nil {
		return dockerFile, nil, err
	}

	return dockerFile, dockerCompose, err
}

func (plugin *DockerFilePlugin) Kill() bool {
	plugin.Client.Kill()
	return plugin.Client.Exited()
}
