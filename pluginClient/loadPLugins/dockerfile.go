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

//DockerFilePlugin contains Methods and Client object of this plugin,
//Setting for logger related info
type DockerFilePlugin struct {
	Methods GlobalFiles.DockerFile
	Client  *plugin.Client
	Setting *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugin *DockerFilePlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("docker plugin client creation started")
	plugin.Methods, plugin.Client = pluginClient.CreateDockerFileClient(utils.CallPluginCommand(yamlFile.Command))
	plugin.Setting.Logger.Debug("docker plugin client created successfully")
}

//Run it runs plugin(execute methods of plugin)
func (plugin *DockerFilePlugin) Run(ctx context.Context, projectRootPath string) (*global.DockerFile, *global.DockerCompose, error) {
	plugin.Setting.Logger.Debug("docker plugin execution started")

	plugin.Setting.Logger.Debug("dockerfile detection started")
	dockerFile, err := plugin.Methods.DetectDockerFile(&helpers.StringInput{Value: projectRootPath})
	if err != nil {
		return nil, nil, err
	}
	plugin.Setting.Logger.Debug("dockerfile detection completed")

	plugin.Setting.Logger.Debug("docker-compose file detection started")
	dockerCompose, err := plugin.Methods.DetectDockerComposeFile(&helpers.StringInput{Value: projectRootPath})
	if err != nil {
		return dockerFile, nil, err
	}
	plugin.Setting.Logger.Debug("docker-compose file detection completed")

	plugin.Setting.Logger.Debug("docker plugin execution completed")
	return dockerFile, dockerCompose, err
}

//Kill will kill client
func (plugin *DockerFilePlugin) Kill() bool {
	plugin.Client.Kill()
	return plugin.Client.Exited()
}
