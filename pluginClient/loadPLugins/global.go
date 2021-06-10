package loadPLugins

import (
	"code-analyser/GlobalFiles"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/pluginDetails"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type GlobalPlugin struct {
	ProcFile   Plugin
	Makefile   Plugin
	DockerFile Plugin
}

type Plugin struct {
	Call interface{}
	Client *plugin.Client
}

type GlobalPluginClient struct {
	ProcFile   *plugin.Client
	Makefile   *plugin.Client
	DockerFile *plugin.Client
}



func (receiver *GlobalPlugin) Load(pluginPaths *pluginDetails.GlobalPlugins) {
	receiver.DockerFile, _ := pluginClient.CreateDockerFileClient(utils.CallPluginCommand(pluginPaths.DockerFilePluginPath))

}
