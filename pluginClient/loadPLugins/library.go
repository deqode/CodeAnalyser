package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type LibraryPlugin struct {
	Libraries map[string]*LibraryVersion
}

type LibraryVersion struct {
	Version map[string]*LibraryPluginDetails
}

type LibraryPluginDetails struct {
	Methods *interfaces.Library
	Client  *plugin.Client
}

func (receiver *LibraryPlugin) Load(name, version, pluginPath string) {
	methods, client := pluginClient.CreateLibraryClient(utils.CallPluginCommand(pluginPath))
	if value, ok := receiver.Libraries[name]; ok {
		value.Version[version] = &LibraryPluginDetails{
			Methods: &methods,
			Client:  client,
		}
	} else {
		receiver.Libraries = map[string]*LibraryVersion{
			name: {
				Version: map[string]*LibraryPluginDetails{
					version: {
						Methods: &methods,
						Client:  client,
					},
				},
			},
		}
	}
}

func (receiver LibraryPlugin) Run(name, version string) {

}
