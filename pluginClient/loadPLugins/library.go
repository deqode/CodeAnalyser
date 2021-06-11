package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type LibraryPlugin struct {
	libraries map[string]*LibraryVersion
}

type LibraryVersion struct {
	version map[string]*LibraryPluginDetails
}

type LibraryPluginDetails struct {
	Methods *interfaces.Library
	Client  *plugin.Client
}

func (receiver *LibraryPlugin) Load(name, version, pluginPath string) {
	methods, client := pluginClient.CreateLibraryClient(utils.CallPluginCommand(pluginPath))
	if value, ok := receiver.libraries[name]; ok {
		value.version[version] = &LibraryPluginDetails{
			Methods: &methods,
			Client:  client,
		}
	} else {
		receiver.libraries = map[string]*LibraryVersion{
			name: {
				version: map[string]*LibraryPluginDetails{
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
