package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"code-analyser/utils"
	"context"
	"os/exec"
)

//todo: think good name

//DependencyDetail It stores version and execution command of dependency as a plugin for example
//beego = {
//           version : 1.5
//           command : "go plugin/beego/main.go"
//        }
type DependencyDetail struct {
	Version string
	Command string
}

const (
	//Framework is FRAMEWORK
	Framework = "FRAMEWORK"
	//DB is DB ( database)
	DB = "DB"
	//ORM is ORM
	ORM = "ORM"
)

//DetectRuntime It will detect language and its version
func DetectRuntime(ctx context.Context, path string, yamlLangObject *protos.LanguageVersion) string {
	runtimeResponse, client := pluginClient.DetectRuntimePluginCall(exec.Command("sh", "-c", yamlLangObject.Detectruntimecommand))
	defer client.Kill()
	runtimeVersion, err := runtimeResponse.DetectRuntime(&pb.ServiceInputString{Value: path})
	if err != nil {
		utils.Logger(err)
		return ""
	}
	if runtimeVersion.Error != nil {
		utils.Logger(runtimeVersion.Error)
		return ""
	}
	return runtimeVersion.Value
}

//GetParsedDependencis get map of parsed dependencies for example beego is a framework
func GetParsedDependencis(ctx context.Context, languageVersion, path string, langYamlObject *protos.LanguageVersion) map[string]map[string]DependencyDetail {
	AllDependencies := map[string]map[string]DependencyDetail{}

	var dependenciesCommand *protos.DependencyVersionDetails
	var runtimeVersion string
	for rt, supportedRuntimeVersions := range pluginDetails.Runtimeversions {
		if helpers.SeverValidateFromArray(supportedRuntimeVersions.Semver, languageVersion) {
			dependenciesCommand = supportedRuntimeVersions
			runtimeVersion = rt
			break
		}
	}
	if dependenciesCommand != nil {
		dependenciesResponse, client := pluginClient.DependenciesPluginCall(exec.Command("sh", "-c", dependenciesCommand.Plugincommand))
		defer client.Kill()
		getdependenciesFound, err := dependenciesResponse.GetDependencies(&pb.ServiceInput{
			RuntimeVersion: runtimeVersion,
			Root:           path,
		})
		if err != nil {
			utils.Logger(err)
			return nil
		}
		if getdependenciesFound.Error != nil {
			utils.Logger(getdependenciesFound.Error)
			return nil
		}
		dependenciesList := getdependenciesFound.Value
		AllDependencies[Framework] = ParseFrameworkFromDependencies(dependenciesList, pluginDetails)
		AllDependencies[DB] = ParseDbFromDependencies(dependenciesList, pluginDetails)
		AllDependencies[ORM] = ParseOrmFromDependencies(dependenciesList, pluginDetails)
		return AllDependencies
	}

	return nil

}
