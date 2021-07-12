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

type GetDependenciesPlugin struct {
	Dependencies map[string]*DependenciesVersion
	Setting *utils.Setting
}

type DependenciesVersion struct {
	Semver  []string
	Methods interfaces.Dependencies
	Client  *plugin.Client
}

func (plugin *GetDependenciesPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateDependenciesClient(utils.CallPluginCommand(yamlFile.Command))
	if plugin.Dependencies == nil {
		plugin.Dependencies = map[string]*DependenciesVersion{}
	}
	plugin.Dependencies[yamlFile.Version] = &DependenciesVersion{
		Semver:  yamlFile.Semver,
		Methods: methods,
		Client:  client,
	}
}

func (plugin *GetDependenciesPlugin) Run(ctx context.Context, runtTimeVersion, rootPath string) (*pbHelpers.StringMapOutput, error) {
	for _, details := range plugin.Dependencies {
		if helpers.SemverValidateFromArray(details.Semver, runtTimeVersion) {
			detect, err := details.Methods.GetDependencies(&pbHelpers.Input{
				RuntimeVersion: runtTimeVersion,
				RootPath:       rootPath,
			})
			if err != nil {
				return nil, err
			}
			return detect, nil
		}
	}
	return &pbHelpers.StringMapOutput{
		Value: nil,
		Error: &pbHelpers.Error{
			Message: "no plugin found for current version " + runtTimeVersion,
			Cause:   "getdependency plugin",
		},
	}, nil

}
