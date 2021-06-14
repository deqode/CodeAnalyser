package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
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

func (receiver *DbPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateDbClient(utils.CallPluginCommand(yamlFile.Command))
	if value, ok := receiver.Dbs[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &DbPluginDetails{
			Methods: &methods,
			Client:  client,
		}
	} else {
		receiver.Dbs = map[string]*DbVersion{
			yamlFile.Name: {
				Version: map[string]*DbPluginDetails{
					yamlFile.Version: {
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
