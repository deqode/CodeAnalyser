package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"github.com/hashicorp/go-plugin"
)

type GetDependenciesPlugin struct {
	Dependencies map[string]*DependenciesVersion
}

type DependenciesVersion struct {
	Methods *interfaces.Dependencies
	Client  *plugin.Client
}

//func (receiver *GetDependenciesPlugin) Load(yamlFile *pbUtils.Details) {
//	methods, client := pluginClient.CreateDependenciesClient(utils.CallPluginCommand(yamlFile.Command))
//	receiver.Client = client
//	receiver.Methods = &methods
//}
