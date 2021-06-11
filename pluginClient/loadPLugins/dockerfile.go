package loadPLugins

import (
	"code-analyser/GlobalFiles"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type DockerFileClient struct {
	Methods *GlobalFiles.DockerFile
	Client  *plugin.Client
}

func (receiver *DockerFileClient) Load(pluginPath string) {
	methods, client := pluginClient.CreateDockerFileClient(utils.CallPluginCommand(pluginPath))
	receiver.Client = client
	receiver.Methods = &methods
}
