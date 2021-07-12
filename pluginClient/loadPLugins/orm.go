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

type OrmPlugin struct {
	Orms map[string]*OrmVersion
	Setting *utils.Setting
}

type OrmVersion struct {
	Version map[string]*OrmPluginDetails
}

type OrmPluginDetails struct {
	Semver  []string
	Methods interfaces.Orm
	Client  *plugin.Client
	Setting *utils.Setting
}

func (plugin *OrmPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateOrmClient(utils.CallPluginCommand(yamlFile.Command))
	if plugin.Orms == nil {
		plugin.Orms = map[string]*OrmVersion{}
	}
	if value, ok := plugin.Orms[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &OrmPluginDetails{
			Methods: methods,
			Client:  client,
			Semver:  yamlFile.Semver,
		}
	} else {
		plugin.Orms[yamlFile.Name] = &OrmVersion{
			Version: map[string]*OrmPluginDetails{
				yamlFile.Version: {
					Methods: methods,
					Client:  client,
					Semver:  yamlFile.Semver,
				},
			},
		}
	}
}

func (plugin *OrmPlugin) Extract(ctx context.Context, projectDependencies map[string]string) []*utils.Dependency {
	var orms []*utils.Dependency
	for name, details := range plugin.Orms {
		if usedOrmVersion, ok := projectDependencies[name]; ok {
			for version, versionDetails := range details.Version {
				if helpers.SemverValidateFromArray(versionDetails.Semver, usedOrmVersion) {
					orms = append(orms, &utils.Dependency{
						Name:    name,
						Version: version,
					})
				}
			}
		}
	}
	return orms
}

func (plugin *OrmPlugin) Run(ctx context.Context, orms []*utils.Dependency, runTimeVersion, projectRootPath string) ([]*languagePB.OrmOutput, error) {
	var output []*languagePB.OrmOutput
	for _, orm := range orms {
		name := orm.Name
		version := orm.Version
		response, err := plugin.Orms[name].Version[version].Run(ctx, name, version, runTimeVersion, projectRootPath)
		if err != nil {
			return nil, err
		}
		output = append(output, response)
	}
	return output, nil
}

func (plugin *OrmPluginDetails) Run(ctx context.Context, name, version, runTimeVersion, projectRootPath string) (*languagePB.OrmOutput, error) {
	pluginInput := &pbHelpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	}
	output := &languagePB.OrmOutput{
		Name:    name,
		Version: version,
	}

	detect, err := plugin.Methods.Detect(pluginInput)
	if err != nil {
		return nil, err
	}
	if detect.Error != nil || detect.Used == false {
		output.Error = detect.Error
		return output, err
	}
	output.DbVersion = detect.DbVersion
	output.DbName = detect.DbName

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
