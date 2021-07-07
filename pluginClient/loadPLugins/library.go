package loadPLugins

import (
	 "code-analyser/helpers"
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbHelpers "code-analyser/protos/pb/helpers"
	languagePB "code-analyser/protos/pb/output/languageSpecific"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

type LibraryPlugin struct {
	Libraries map[string]*LibraryVersion
}

type LibraryVersion struct {
	Version map[string]*LibraryPluginDetails
}

type LibraryPluginDetails struct {
	Semver  []string
	Methods interfaces.Library
	Client  *plugin.Client
}

func (plugin *LibraryPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateLibraryClient(utils.CallPluginCommand(yamlFile.Command))
	if plugin.Libraries == nil {
		plugin.Libraries = map[string]*LibraryVersion{}
	}
	if value, ok := plugin.Libraries[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &LibraryPluginDetails{
			Methods: methods,
			Client:  client,
			Semver:  yamlFile.Semver,
		}
	} else {
		plugin.Libraries[yamlFile.Name] = &LibraryVersion{
			Version: map[string]*LibraryPluginDetails{
				yamlFile.Version: {
					Methods: methods,
					Client:  client,
					Semver:  yamlFile.Semver,
				},
			},
		}
	}
}

func (plugin *LibraryPlugin) Extract(ctx context.Context, projectDependencies map[string]string) []*utils.Dependency {
	var libraries []*utils.Dependency
	for name, details := range plugin.Libraries {
		if usedLibraryVersion, ok := projectDependencies[name]; ok {
			for version, versionDetails := range details.Version {
				if helpers.SemverValidateFromArray(versionDetails.Semver, usedLibraryVersion) {
					libraries = append(libraries, &utils.Dependency{
						Name:    name,
						Version: version,
					})
				}
			}
		}
	}
	return libraries
}

func (plugin *LibraryPlugin) Run(ctx context.Context, libraries []*utils.Dependency, runTimeVersion, projectRootPath string) ([]*languagePB.LibraryOutput, error) {
	var output []*languagePB.LibraryOutput
	for _, framework := range libraries {
		name := framework.Name
		version := framework.Version
		response, err := plugin.Libraries[name].Version[version].Run(ctx, name, version, runTimeVersion, projectRootPath)
		if err != nil {
			return nil, err
		}
		output = append(output, response)
	}
	return output, nil
}

func (plugin *LibraryPluginDetails) Run(ctx context.Context, name, version, runTimeVersion, projectRootPath string) (*languagePB.LibraryOutput, error) {
	pluginInput := &pbHelpers.Input{
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
