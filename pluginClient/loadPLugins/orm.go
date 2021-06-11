package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type OrmPlugin struct {
	Orms map[string]*OrmVersion
}

type OrmVersion struct {
	Version map[string]*OrmPluginDetails
}

type OrmPluginDetails struct {
	Methods *interfaces.Orm
	Client  *plugin.Client
}

func (receiver *OrmPlugin) Load(name, version, pluginPath string) {
	methods, client := pluginClient.CreateOrmClient(utils.CallPluginCommand(pluginPath))
	if value, ok := receiver.Orms[name]; ok {
		value.Version[version] = &OrmPluginDetails{
			Methods: &methods,
			Client:  client,
		}
	} else {
		receiver.Orms = map[string]*OrmVersion{
			name: {
				Version: map[string]*OrmPluginDetails{
					version: {
						Methods: &methods,
						Client:  client,
					},
				},
			},
		}
	}
}

func (receiver *OrmPlugin) Run(name, version string) {

}