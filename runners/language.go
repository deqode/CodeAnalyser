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

type DependencyDetail struct {
	Version string
	Command string
}

const (
	Framework = "FRAMEWORK"
	DB        = "DB"
	ORM       = "ORM"
)

func DetectRuntime(ctx context.Context, path string, pluginDetails *protos.LanguageVersion) string {
	runtimeResponse, client := pluginClient.DetectRuntimePluginCall(exec.Command("sh", "-c", pluginDetails.Detectruntimecommand))
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

func GetParsedDependencis(ctx context.Context, languageVersion, path string, pluginDetails *protos.LanguageVersion) map[string]map[string]DependencyDetail {
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
