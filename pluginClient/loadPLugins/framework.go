package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
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
	Methods interfaces.Framework
	Client  *plugin.Client
}

func (plugin *FrameworkPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateFrameworkClient(utils.CallPluginCommand(yamlFile.Command))
	if value, ok := plugin.Frameworks[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &FrameworkPluginDetails{
			Methods: methods,
			Client:  client,
		}
	} else {
		plugin.Frameworks = map[string]*FrameworkVersion{
			yamlFile.Name: {
				Version: map[string]*FrameworkPluginDetails{
					yamlFile.Version: {
						Methods: methods,
						Client:  client,
					},
				},
			},
		}
	}
}

func (plugin *FrameworkPluginDetails) Run(name, version, runTimeVersion, projectRootPath string) (*languageSpecific.FrameworkOutput, error) {
	pluginInput := &helpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	}
	output := &languageSpecific.FrameworkOutput{
		Name:    name,
		Version: version,
	}

	detect, err := plugin.Methods.Detect(pluginInput)
	if err != nil {
		return nil, err
	}
	if detect.Error != nil || detect.Value == false {
		output.Error = detect.Error
		return output, err
	}

	isUsed, err := plugin.Methods.IsUsed(pluginInput)
	if err != nil {
		return nil, err
	}
	if isUsed.Error != nil {
		output.Error = isUsed.Error
		return output, err
	}
	output.Used = isUsed.Value

	percentageUsed, err := plugin.Methods.PercentOfUsed(pluginInput)
	if err != nil {
		return nil, err
	}
	if percentageUsed.Error != nil {
		output.Error = percentageUsed.Error
		return output, err
	}
	output.PercentageUsed = percentageUsed.Value

	return output, nil
}
