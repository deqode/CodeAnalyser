package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
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

func (receiver *FrameworkPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateFrameworkClient(utils.CallPluginCommand(yamlFile.Command))
	if value, ok := receiver.Frameworks[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &FrameworkPluginDetails{
			Methods: &methods,
			Client:  client,
		}
	} else {
		receiver.Frameworks = map[string]*FrameworkVersion{
			yamlFile.Name: {
				Version: map[string]*FrameworkPluginDetails{
					yamlFile.Version: {
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
