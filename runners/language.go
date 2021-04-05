package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"code-analyser/utils"
	"context"
	"log"
	"os/exec"
)

//todo: think good name

type DependencyDetail struct {
	Version    string
	Command string
}

const (
	Framework = "FRAMEWORK"
	DB        = "DB"
	ORM       = "ORM"
)

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

func GetParsedDependencis(ctx context.Context,languageVersion, path string, langYamlObject *protos.LanguageVersion) *protos.LanguageVersion {
	AllDependencies := map[string]map[string]DependencyDetail{}

	var dependenciesCommand *protos.DependencyVersionDetails
	var runtimeVersion string
	for rt, supportedRuntimeVersions := range langYamlObject.Runtimeversions {
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

		//log.Println(dependenciesList)

		AllDependencies[Framework] = ParseFrameworkFromDependencies(dependenciesList, langYamlObject)
		AllDependencies[DB] = ParseDbFromDependencies(dependenciesList, langYamlObject)
		AllDependencies[ORM] = ParseOrmFromDependencies(dependenciesList, langYamlObject)
		log.Println(AllDependencies)
		//log.Println(OrmRunner(AllDependencies[ORM], runtimeVersion, path).Orms)
		//log.Println(DbRunner(AllDependencies[DB], runtimeVersion, path).Databases)
		//log.Println(FrameworkRunner(AllDependencies[Framework], runtimeVersion, path))

	}

	return nil

}
