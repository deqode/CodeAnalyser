package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type FrameworkPlugin struct {
	Frameworks map[string]*FrameworkVersion
}

type FrameworkVersion struct {
	Version map[string]*FrameworkPluginDetails
}

type FrameworkPluginDetails struct {
	Methods *interfaces.Framework
	Client  *plugin.Client
}

func (receiver *FrameworkPlugin) Load(name, version, pluginPath string) {
	methods, client := pluginClient.CreateFrameworkClient(utils.CallPluginCommand(pluginPath))
	if value, ok := receiver.Frameworks[name]; ok {
		value.Version[version] = &FrameworkPluginDetails{
			Methods: &methods,
			Client:  client,
		}
	} else {
		receiver.Frameworks = map[string]*FrameworkVersion{
			name: {
				Version: map[string]*FrameworkPluginDetails{
					version: {
						Methods: &methods,
						Client:  client,
					},
				},
			},
		}
	}
}

func (receiver FrameworkPlugin) Run(name, version string) {

}
