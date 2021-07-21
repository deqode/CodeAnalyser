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

//OrmPlugin contains 1. map of Orms where key is orm name and value is OrmVersion
// 2. contains Setting
type OrmPlugin struct {
	Orms    map[string]*OrmVersion
	Setting *utils.Setting
}

//OrmVersion It contains OrmPlugin.Orms version map(OrmVersion) where key is version of orm and value is OrmPluginDetails
type OrmVersion struct {
	Version map[string]*OrmPluginDetails
}

//OrmPluginDetails contains Methods and Client object of this plugin,
//Setting for logger related info
type OrmPluginDetails struct {
	Semver  []string
	Methods interfaces.Orm
	Client  *plugin.Client
	Setting *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugins *OrmPlugin) Load(yamlFile *pbUtils.Details) {
	plugins.Setting.Logger.Debug(yamlFile.Name + " plugins client creation started")

	methods, client := pluginClient.CreateOrmClient(utils.CallPluginCommand(yamlFile.Command))
	if plugins.Orms == nil {
		plugins.Orms = map[string]*OrmVersion{}
	}
	if value, ok := plugins.Orms[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &OrmPluginDetails{
			Methods: methods,
			Client:  client,
			Semver:  yamlFile.Semver,
			Setting: plugins.Setting,
		}
	} else {
		plugins.Orms[yamlFile.Name] = &OrmVersion{
			Version: map[string]*OrmPluginDetails{
				yamlFile.Version: {
					Methods: methods,
					Client:  client,
					Semver:  yamlFile.Semver,
					Setting: plugins.Setting,
				},
			},
		}
	}

	plugins.Setting.Logger.Debug(yamlFile.Name + " plugins client created successfully")
}

//Extract It takes dependency map as an input and filter out dependency supported by us(calculate intersection between dependency list and plugin list)
func (plugins *OrmPlugin) Extract(ctx context.Context, projectDependencies map[string]string) []*utils.Dependency {
	plugins.Setting.Logger.Debug("filtration process of orm's plugins supported by us started")

	var orms []*utils.Dependency
	for name, details := range plugins.Orms {
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

	plugins.Setting.Logger.Debug("filtration process of orm's plugins completed")
	return orms
}

//Run it takes dependency list, fetch out plugin client from receiver and run one by one
func (plugins *OrmPlugin) Run(ctx context.Context, orms []*utils.Dependency, runTimeVersion, projectRootPath string) ([]*languagePB.OrmOutput, error) {
	plugins.Setting.Logger.Debug("orm's plugins methods execution started")

	var output []*languagePB.OrmOutput
	for _, orm := range orms {
		name := orm.Name
		version := orm.Version

		plugins.Setting.Logger.Info(name + " plugins execution started")
		response, err := plugins.Orms[name].Version[version].Run(ctx, name, version, runTimeVersion, projectRootPath)
		if err != nil {
			return nil, err
		}
		plugins.Setting.Logger.Info(name + " plugins execution completed")

		output = append(output, response)
	}

	plugins.Setting.Logger.Debug("orm's plugins methods execution completed")
	return output, nil
}

//Run it runs plugin(execute methods of plugin)
func (plugins *OrmPluginDetails) Run(ctx context.Context, name, version, runTimeVersion, projectRootPath string) (*languagePB.OrmOutput, error) {
	plugins.Setting.Logger.Debug(name + " plugins execution started")

	pluginInput := &pbHelpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	}
	output := &languagePB.OrmOutput{
		Name:    name,
		Version: version,
	}

	plugins.Setting.Logger.Debug(name + "'s detection started")
	detect, err := plugins.Methods.Detect(pluginInput)
	if err != nil {
		return nil, err
	}
	if detect.Error != nil || detect.Used == false {
		output.Error = detect.Error
		return output, err
	}
	output.DbVersion = detect.DbVersion
	output.DbName = detect.DbName
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
