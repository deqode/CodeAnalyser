package loadPLugins

import (
	"code-analyser/helpers"
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbHelpers "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

//DbPlugin contains 1. map of Dbs where key is db name and value is DbVersion
// 2. contains Setting
type DbPlugin struct {
	Dbs     map[string]*DbVersion
	Setting *utils.Setting
}

//DbVersion It contains DbPlugin.Dbs version map(DbVersion) where key is version of library and value is DbPluginDetails
type DbVersion struct {
	Version map[string]*DbPluginDetails
}

//DbPluginDetails contains Methods, Client object and Libraries of this plugin,
//Setting for logger related info
type DbPluginDetails struct {
	Libraries []*pbUtils.Library
	Methods   interfaces.Db
	Client    *plugin.Client
	Setting   *utils.Setting
}

//Load It takes plugin command and creates client(hashicorp plugin structure client)
func (plugins *DbPlugin) Load(yamlFile *pbUtils.Details) {
	plugins.Setting.Logger.Debug(yamlFile.Name + "db plugin client creation started")

	methods, client := pluginClient.CreateDbClient(utils.CallPluginCommand(yamlFile.Command))
	if plugins.Dbs == nil {
		plugins.Dbs = map[string]*DbVersion{}
	}
	if value, ok := plugins.Dbs[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &DbPluginDetails{
			Methods:   methods,
			Client:    client,
			Libraries: yamlFile.Libraries,
			Setting:   plugins.Setting,
		}
	} else {
		plugins.Dbs[yamlFile.Name] = &DbVersion{
			Version: map[string]*DbPluginDetails{
				yamlFile.Version: {
					Methods:   methods,
					Client:    client,
					Libraries: yamlFile.Libraries,
					Setting:   plugins.Setting,
				},
			},
		}
	}

	plugins.Setting.Logger.Debug(yamlFile.Name + "db plugin client created successfully")
}

//Extract It takes dependency map as an input and filter out dependency supported by us(calculate intersection between dependency list and plugin list)
func (plugins *DbPlugin) Extract(ctx context.Context, projectDependencies map[string]string) []*utils.Dependency {
	plugins.Setting.Logger.Debug("filtration process of db's plugin supported by us started")

	var dbs []*utils.Dependency
	for name, details := range plugins.Dbs {
		for dbVersion, versionDetails := range details.Version {
			for _, library := range versionDetails.Libraries {
				if usedLibraryVersion, ok := projectDependencies[library.Name]; ok {
					if helpers.SemverValidateFromArray(library.Semver, usedLibraryVersion) {
						dbs = append(dbs, &utils.Dependency{
							Version: dbVersion,
							Name:    name,
						})
					}
				}
			}
		}
	}

	plugins.Setting.Logger.Debug("filtration process of db's plugin completed")
	return dbs
}

//Run it takes dependency list, fetch out plugin client from receiver and run one by one
func (plugins *DbPlugin) Run(ctx context.Context, dbs []*utils.Dependency, runTimeVersion, projectRootPath string) ([]*languageSpecific.DBOutput, error) {
	plugins.Setting.Logger.Debug("db's plugin methods execution started")

	var output []*languageSpecific.DBOutput
	for _, db := range dbs {
		name := db.Name
		version := db.Version

		plugins.Setting.Logger.Info(name + " plugin execution started")
		response, err := plugins.Dbs[name].Version[version].Run(ctx, name, version, runTimeVersion, projectRootPath)
		if err != nil {
			return nil, err
		}
		plugins.Setting.Logger.Info(name + " plugin execution completed")

		output = append(output, response)
	}

	plugins.Setting.Logger.Debug("db's plugin methods execution completed")
	return output, nil
}

//Run it runs plugin(execute methods of plugin)
func (plugins *DbPluginDetails) Run(ctx context.Context, name, version, runTimeVersion, projectRootPath string) (*languageSpecific.DBOutput, error) {
	plugins.Setting.Logger.Debug(name + " plugin execution started")

	pluginInput := &pbHelpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	}
	output := &languageSpecific.DBOutput{
		Name:    name,
		Version: version,
	}

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

	plugins.Setting.Logger.Debug(name + "'s detection started")
	detect, err := plugins.Methods.Detect(pluginInput)
	if err != nil {
		return nil, err
	}
	if detect.Error != nil || detect.Value == false {
		output.Error = detect.Error
		return output, err
	}
	output.Port = detect.IntValue
	plugins.Setting.Logger.Debug(name + "'s detection completed")

	plugins.Setting.Logger.Debug(name + " plugin execution completed")
	return output, nil
}
