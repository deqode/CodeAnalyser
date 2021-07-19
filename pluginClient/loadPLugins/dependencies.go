package loadPLugins

import (
	"code-analyser/helpers"
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbHelpers "code-analyser/protos/pb/helpers"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

//GetDependenciesPlugin contains 1. map of Dependencies where key is dependency name and value is DependenciesVersion
// 2. Setting
type GetDependenciesPlugin struct {
	Dependencies map[string]*DependenciesVersion
	Setting      *utils.Setting
}

//DependenciesVersion contains Methods, Client object of this plugin,
type DependenciesVersion struct {
	Semver  []string
	Methods interfaces.Dependencies
	Client  *plugin.Client
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugin *GetDependenciesPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Setting.Logger.Debug("getDependencies plugin client creation started")

	methods, client := pluginClient.CreateDependenciesClient(utils.CallPluginCommand(yamlFile.Command))
	if plugin.Dependencies == nil {
		plugin.Dependencies = map[string]*DependenciesVersion{}
	}
	plugin.Dependencies[yamlFile.Version] = &DependenciesVersion{
		Semver:  yamlFile.Semver,
		Methods: methods,
		Client:  client,
	}

	plugin.Setting.Logger.Debug("getDependencies plugin client created successfully")
}

//Run it runs plugin(execute methods of plugin)
func (plugin *GetDependenciesPlugin) Run(ctx context.Context, runtTimeVersion, rootPath string) (*pbHelpers.StringMapOutput, error) {
	plugin.Setting.Logger.Debug("detectDependencies plugin methods execution started")

	for _, details := range plugin.Dependencies {
		if helpers.SemverValidateFromArray(details.Semver, runtTimeVersion) {

			plugin.Setting.Logger.Debug("dependency fetching started")
			detect, err := details.Methods.GetDependencies(&pbHelpers.Input{
				RuntimeVersion: runtTimeVersion,
				RootPath:       rootPath,
			})
			if err != nil {
				return nil, err
			}
			plugin.Setting.Logger.Debug("dependency fetched successfully")

			return detect, nil
		}
	}

	plugin.Setting.Logger.Debug("detectDependencies plugin methods executed successfully")
	return &pbHelpers.StringMapOutput{
		Value: nil,
		Error: &pbHelpers.Error{
			Message: "no plugin found for current version " + runtTimeVersion,
			Cause:   "getdependency plugin",
		},
	}, nil

}
