package loadPLugins

import (
	helpers "code-analyser/helpers"
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbHelpers "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

type FrameworkPlugin struct {
	Frameworks map[string]*FrameworkVersion
	Setting *utils.Setting
}

type FrameworkVersion struct {
	Version map[string]*FrameworkPluginDetails
}

type FrameworkPluginDetails struct {
	Libraries []*pbUtils.Library
	Methods   interfaces.Framework
	Client    *plugin.Client
	Setting *utils.Setting
}

func (plugins *FrameworkPlugin) Load(yamlFile *pbUtils.Details) {
	plugins.Setting.Logger.Debug(yamlFile.Name + " plugin client creation started")

	methods, client := pluginClient.CreateFrameworkClient(utils.CallPluginCommand(yamlFile.Command))
	if plugins.Frameworks == nil {
		plugins.Frameworks = map[string]*FrameworkVersion{}
	}
	if value, ok := plugins.Frameworks[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &FrameworkPluginDetails{
			Methods: methods,
			Client:  client,
			Libraries: yamlFile.Libraries,
		}
	} else {
		plugins.Frameworks[yamlFile.Name] = &FrameworkVersion{
			Version: map[string]*FrameworkPluginDetails{
				yamlFile.Version: {
					Methods: methods,
					Client:  client,
					Libraries: yamlFile.Libraries,
				},
			},
		}
	}

	plugins.Setting.Logger.Debug(yamlFile.Name + " plugin client created successfully")
}


func (plugins *FrameworkPlugin) Extract(ctx context.Context, projectDependencies map[string]string) []*utils.Dependency {
	plugins.Setting.Logger.Debug("filtration process of framework's plugin supported by us started")

	var frameworks []*utils.Dependency
	for name, details := range plugins.Frameworks {
		for frameworkVersion, versionDetails := range details.Version {
			for _, library := range versionDetails.Libraries {
				if usedLibraryVersion, ok := projectDependencies[library.Name]; ok {
					if helpers.SemverValidateFromArray(library.Semver, usedLibraryVersion) {
						frameworks = append(frameworks, &utils.Dependency{
							Version: frameworkVersion,
							Name:    name,
						})
					}
				}
			}
		}
	}

	plugins.Setting.Logger.Debug("filtration process of framework's plugin completed")
	return frameworks
}

func (plugins *FrameworkPlugin) Run(ctx context.Context, frameworks []*utils.Dependency, runTimeVersion, projectRootPath string) ([]*languageSpecific.FrameworkOutput, error) {
	var output []*languageSpecific.FrameworkOutput
	for _, framework := range frameworks {
		name := framework.Name
		version := framework.Version
		response, err := plugins.Frameworks[name].Version[version].Run(ctx, name, version, runTimeVersion, projectRootPath)
		if err != nil {
			return nil, err
		}
		output = append(output, response)
	}
	return output, nil
}

func (plugins *FrameworkPluginDetails) Run(ctx context.Context, name, version, runTimeVersion, projectRootPath string) (*languageSpecific.FrameworkOutput, error) {
	pluginInput := &pbHelpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	}
	output := &languageSpecific.FrameworkOutput{
		Name:    name,
		Version: version,
	}

	detect, err := plugins.Methods.Detect(pluginInput)
	if err != nil {
		return nil, err
	}
	if detect.Error != nil || detect.Value == false {
		output.Error = detect.Error
		return output, err
	}

	isUsed, err := plugins.Methods.IsUsed(pluginInput)
	if err != nil {
		return nil, err
	}
	if isUsed.Error != nil {
		output.Error = isUsed.Error
		return output, err
	}
	output.Used = isUsed.Value

	percentageUsed, err := plugins.Methods.PercentOfUsed(pluginInput)
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
