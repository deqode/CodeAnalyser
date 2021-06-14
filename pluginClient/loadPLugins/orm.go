package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
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

func (receiver *OrmPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateOrmClient(utils.CallPluginCommand(yamlFile.Command))
	if value, ok := receiver.Orms[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &OrmPluginDetails{
			Methods: &methods,
			Client:  client,
		}
	} else {
		receiver.Orms = map[string]*OrmVersion{
			yamlFile.Name: {
				Version: map[string]*OrmPluginDetails{
					yamlFile.Version: {
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