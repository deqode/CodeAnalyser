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
	Setting   *utils.Setting
}

type LibraryVersion struct {
	Version map[string]*LibraryPluginDetails
}

type LibraryPluginDetails struct {
	Semver  []string
	Methods interfaces.Library
	Client  *plugin.Client
	Setting *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugins *LibraryPlugin) Load(yamlFile *pbUtils.Details) {
	plugins.Setting.Logger.Debug(yamlFile.Name + " plugins client creation started")

	methods, client := pluginClient.CreateLibraryClient(utils.CallPluginCommand(yamlFile.Command))
	if plugins.Libraries == nil {
		plugins.Libraries = map[string]*LibraryVersion{}
	}
	if value, ok := plugins.Libraries[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &LibraryPluginDetails{
			Methods: methods,
			Client:  client,
			Semver:  yamlFile.Semver,
			Setting: plugins.Setting,
		}
	} else {
		plugins.Libraries[yamlFile.Name] = &LibraryVersion{
			Version: map[string]*LibraryPluginDetails{
				yamlFile.Version: {
					Methods: methods,
					Client:  client,
					Semver:  yamlFile.Semver,
					Setting: plugins.Setting,
				},
			},
		}
	}

	plugins.Setting.Logger.Debug(yamlFile.Name + " plugin client created successfully")
}

//Extract It takes dependency map as an input and filter out dependency supported by us(calculate intersection between dependency list and plugin list)
func (plugins *LibraryPlugin) Extract(ctx context.Context, projectDependencies map[string]string) []*utils.Dependency {
	plugins.Setting.Logger.Debug("filtration process of library's plugin supported by us started")

	var libraries []*utils.Dependency
	for name, details := range plugins.Libraries {
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

	plugins.Setting.Logger.Debug("filtration process of library's plugin completed")
	return libraries
}

//Run it takes dependency list, fetch out plugin client from receiver and run one by one
func (plugins *LibraryPlugin) Run(ctx context.Context, libraries []*utils.Dependency, runTimeVersion, projectRootPath string) ([]*languagePB.LibraryOutput, error) {
	plugins.Setting.Logger.Debug("library's plugin methods execution started")

	var output []*languagePB.LibraryOutput
	for _, framework := range libraries {
		name := framework.Name
		version := framework.Version

		plugins.Setting.Logger.Info(name + " plugin execution started")
		response, err := plugins.Libraries[name].Version[version].Run(ctx, name, version, runTimeVersion, projectRootPath)
		if err != nil {
			return nil, err
		}
		plugins.Setting.Logger.Info(name + " plugin execution completed")

		output = append(output, response)
	}

	plugins.Setting.Logger.Debug("library's plugin methods execution completed")
	return output, nil
}

//Run it runs plugin(execute methods of plugin)
func (plugins *LibraryPluginDetails) Run(ctx context.Context, name, version, runTimeVersion, projectRootPath string) (*languagePB.LibraryOutput, error) {
	plugins.Setting.Logger.Debug(name + " plugin execution started")

	pluginInput := &pbHelpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	}
	output := &languagePB.LibraryOutput{
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
	output.Type = detect.Type
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
