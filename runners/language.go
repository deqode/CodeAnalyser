package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"context"
	"log"
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
	Framework = "framework"
	//DB is DB ( database)
	DB = "db"
	//ORM is ORM
	ORM     = "orm"
	Library = "library"
)

//DetectRuntime It will detect language and its version
func DetectRuntime(ctx context.Context, path string, yamlLangObject *versionsPB.LanguageVersion) string {
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

func RunStaticAssetsCommand(ctx context.Context, input *pb.ServiceInput, pluginDetails *versionsPB.LanguageVersion) *languageSpecific.StaticAssetsOutput {
	res, client := pluginClient.StaticAssetsPluginCall(exec.Command("sh", "-c", pluginDetails.StaticAssetsCommand))
	defer client.Kill()
	detection, err := res.Detect(input)
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if detection.Error != nil {
		utils.Logger(err)
		return nil
	}
	staticAssetsOutput := languageSpecific.StaticAssetsOutput{
		Assets: detection.StaticAsset,
	}
	if len(detection.StaticAsset) > 0 {
		staticAssetsOutput.Used = true
	}
	return &staticAssetsOutput
}

func RunPreDetectCommand(ctx context.Context, input *pb.ServiceInput, pluginDetails *versionsPB.LanguageVersion) {
	runtimeResponse, client := pluginClient.PreDetectCommandPluginCall(exec.Command("sh", "-c", pluginDetails.PreDetectCommand))
	defer client.Kill()
	runtimeVersion, err := runtimeResponse.RunPreDetect(input)
	if err != nil {
		utils.Logger(err)
	}
	if runtimeVersion.Error != nil {
		utils.Logger(runtimeVersion.Error)
	}
	log.Println("------------------------ pre detect commands found and executed successfully ")
}

//GetParsedDependencis get map of parsed dependencies for example beego is a framework
func GetParsedDependencis(ctx context.Context, languageVersion, path string, pluginDetails *versionsPB.LanguageVersion) map[string]map[string]DependencyDetail {
	AllDependencies := map[string]map[string]DependencyDetail{}
	var dependenciesCommand *versionsPB.DependencyVersionDetails
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
		AllDependencies[Library] = ParseLibraryFromDependencies(dependenciesList, pluginDetails)
		return AllDependencies
	}

	return nil

}
