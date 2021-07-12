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

type DbPlugin struct {
	Dbs map[string]*DbVersion
	Setting *utils.Setting
}

type DbVersion struct {
	Version map[string]*DbPluginDetails
}

type DbPluginDetails struct {
	Libraries []*pbUtils.Library
	Methods   interfaces.Db
	Client    *plugin.Client
	Setting *utils.Setting
}

func (plugins *DbPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateDbClient(utils.CallPluginCommand(yamlFile.Command))
	if plugins.Dbs == nil {
		plugins.Dbs = map[string]*DbVersion{}
	}
	if value, ok := plugins.Dbs[yamlFile.Name]; ok {
		value.Version[yamlFile.Version] = &DbPluginDetails{
			Methods:   methods,
			Client:    client,
			Libraries: yamlFile.Libraries,
		}
	} else {
		plugins.Dbs[yamlFile.Name] = &DbVersion{
			Version: map[string]*DbPluginDetails{
				yamlFile.Version: {
					Methods:   methods,
					Client:    client,
					Libraries: yamlFile.Libraries,
				},
			},
		}
	}
}

func (plugins *DbPlugin) Extract(ctx context.Context, projectDependencies map[string]string) []*utils.Dependency {
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
	return dbs
}

func (plugins *DbPlugin) Run(ctx context.Context, dbs []*utils.Dependency, runTimeVersion, projectRootPath string) ([]*languageSpecific.DBOutput, error) {
	var output []*languageSpecific.DBOutput
	for _, db := range dbs {
		name := db.Name
		version := db.Version
		response, err := plugins.Dbs[name].Version[version].Run(ctx, name, version, runTimeVersion, projectRootPath)
		if err != nil {
			return nil, err
		}
		output = append(output, response)
	}
	return output, nil
}

func (plugins *DbPluginDetails) Run(ctx context.Context, name, version, runTimeVersion, projectRootPath string) (*languageSpecific.DBOutput, error) {
	pluginInput := &pbHelpers.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	}
	output := &languageSpecific.DBOutput{
		Name:    name,
		Version: version,
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

	detect, err := plugins.Methods.Detect(pluginInput)
	if err != nil {
		return nil, err
	}
	if detect.Error != nil || detect.Value == false {
		output.Error = detect.Error
		return output, err
	}
	output.Port = detect.IntValue

	return output, nil
}
