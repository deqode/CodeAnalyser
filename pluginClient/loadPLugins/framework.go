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

//FrameworkPlugin contains 1. map of Frameworks where key is framework name and value is FrameworkVersion
// 2. contains Setting
type FrameworkPlugin struct {
	Frameworks map[string]*FrameworkVersion
	Setting    *utils.Setting
}

//FrameworkVersion It contains FrameworkPlugin.Frameworks's version map(FrameworkVersion) where key is version of library and value is FrameworkPluginDetails
type FrameworkVersion struct {
	Version map[string]*FrameworkPluginDetails
}

//FrameworkPluginDetails contains Methods, Client object and Libraries of this plugin,
//Setting for logger related info
type FrameworkPluginDetails struct {
	Libraries []*pbUtils.Library
	Methods   interfaces.Framework
	Client    *plugin.Client
	Setting   *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugins *FrameworkPlugin) Load(yamlFile *pbUtils.Details) {
	plugins.Setting.Logger.Debug(yamlFile.Name + " plugin client creation started")

	methods, client := pluginClient.CreateFrameworkClient(utils.CallPluginCommand(yamlFile.Command))
	if plugins.Frameworks == nil {
		plugins.Frameworks = map[string]*FrameworkVersion{}
	}
	if value, ok := plugins.Frameworks[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &FrameworkPluginDetails{
			Methods:   methods,
			Client:    client,
			Libraries: yamlFile.Libraries,
			Setting:   plugins.Setting,
		}
	} else {
		plugins.Frameworks[yamlFile.Name] = &FrameworkVersion{
			Version: map[string]*FrameworkPluginDetails{
				yamlFile.Version: {
					Methods:   methods,
					Client:    client,
					Libraries: yamlFile.Libraries,
					Setting:   plugins.Setting,
				},
			},
		}
	}

	plugins.Setting.Logger.Debug(yamlFile.Name + " plugin client created successfully")
}

//Extract It takes dependency map as an input and filter out dependency supported by us(calculate intersection between dependency list and plugin list)
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

//Run it takes dependency list, fetch out plugin client from receiver and run one by one
func (plugins *FrameworkPlugin) Run(ctx context.Context, frameworks []*utils.Dependency, runTimeVersion, projectRootPath string) ([]*languageSpecific.FrameworkOutput, error) {
	plugins.Setting.Logger.Debug("framework's plugin methods execution started")

	var output []*languageSpecific.FrameworkOutput
	for _, framework := range frameworks {
		name := framework.Name
		version := framework.Version

		plugins.Setting.Logger.Info(name + " plugin execution started")
		response, err := plugins.Frameworks[name].Version[version].Run(ctx, name, version, runTimeVersion, projectRootPath)
		if err != nil {
			return nil, err
		}
		plugins.Setting.Logger.Info(name + " plugin execution completed")

		output = append(output, response)
	}

	plugins.Setting.Logger.Debug("framework's plugin methods execution completed")
	return output, nil
}

//Run it runs plugin(execute methods of plugin)
func (plugins *FrameworkPluginDetails) Run(ctx context.Context, name, version, runTimeVersion, projectRootPath string) (*languageSpecific.FrameworkOutput, error) {
	plugins.Setting.Logger.Debug(name + " plugin execution started")

	pluginInput := &pbHelpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	}
	output := &languageSpecific.FrameworkOutput{
		Name:    name,
		Version: version,
	}

	plugins.Setting.Logger.Debug(name + "'s detection started")
	detect, err := plugins.Methods.Detect(pluginInput)
	if err != nil {
		return nil, err
	}
	if detect.Error != nil || detect.Value == false {
		output.Error = detect.Error
		return output, err
	}
	plugins.Setting.Logger.Debug(name + "'s detection completed")

	plugins.Setting.Logger.Debug(name + "'s usage checking started")
	isUsed, err := plugins.Methods.IsUsed(pluginInput)
	if err != nil {
		return nil, err
	}
	if isUsed.Error != nil {
		output.Error = isUsed.Error
		return output, err
	}
	output.Used = isUsed.Value
	plugins.Setting.Logger.Debug(name + "'s usage checking completed")

	plugins.Setting.Logger.Debug(name + "'s percentage usage checking started")
	percentageUsed, err := plugins.Methods.PercentOfUsed(pluginInput)
	if err != nil {
		return nil, err
	}
	if percentageUsed.Error != nil {
		output.Error = percentageUsed.Error
		return output, err
	}
	output.PercentageUsed = percentageUsed.Value
	plugins.Setting.Logger.Debug(name + "'s percentage usage checking completed")

	plugins.Setting.Logger.Debug(name + " plugin execution completed")
	return output, nil
}
