package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type DbPlugin struct {
	Dbs map[string]*DbVersion
}

type DbVersion struct {
	Version map[string]*DbPluginDetails
}

type DbPluginDetails struct {
	Methods *interfaces.Db
	Client  *plugin.Client
}

func (receiver *DbPlugin) Load(name, version, pluginPath string) {
	methods, client := pluginClient.CreateDbClient(utils.CallPluginCommand(pluginPath))
	if value, ok := receiver.Dbs[name]; ok {
		value.Version[version] = &DbPluginDetails{
			Methods: &methods,
			Client:  client,
		}
	} else {
		receiver.Dbs = map[string]*DbVersion{
			name: {
				Version: map[string]*DbPluginDetails{
					version: {
						Methods: &methods,
						Client:  client,
					},
				},
			},
		}
	}
}

func (receiver DbPlugin) Run(name, version string) {

}