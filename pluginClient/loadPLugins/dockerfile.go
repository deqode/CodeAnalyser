package loadPLugins

import (
	"code-analyser/GlobalFiles"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type DockerFilePlugin struct {
	Methods *GlobalFiles.DockerFile
	Client  *plugin.Client
}

func (receiver *DockerFilePlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateDockerFileClient(utils.CallPluginCommand(yamlFile.Command))
	receiver.Client = client
	receiver.Methods = &methods
}
