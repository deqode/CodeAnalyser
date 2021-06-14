package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
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

func (receiver *LibraryPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateLibraryClient(utils.CallPluginCommand(yamlFile.Command))
	if value, ok := receiver.Libraries[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &LibraryPluginDetails{
			Methods: &methods,
			Client:  client,
		}
	} else {
		receiver.Libraries = map[string]*LibraryVersion{
			yamlFile.Name: {
				Version: map[string]*LibraryPluginDetails{
					yamlFile.Version: {
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
