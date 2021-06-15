package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/helpers"
	languagePB "code-analyser/protos/pb/output/languageSpecific"
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
	Methods interfaces.Library
	Client  *plugin.Client
}

func (plugin *LibraryPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateLibraryClient(utils.CallPluginCommand(yamlFile.Command))
	if value, ok := plugin.Libraries[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &LibraryPluginDetails{
			Methods: methods,
			Client:  client,
		}
	} else {
		plugin.Libraries = map[string]*LibraryVersion{
			yamlFile.Name: {
				Version: map[string]*LibraryPluginDetails{
					yamlFile.Version: {
						Methods: methods,
						Client:  client,
					},
				},
			},
		}
	}
}

func (plugin *LibraryPluginDetails) Run(name, version, runTimeVersion, projectRootPath string) (*languagePB.LibraryOutput, error) {
	pluginInput := &helpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	}
	output := &languagePB.LibraryOutput{
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
	output.Type = detect.Type

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

