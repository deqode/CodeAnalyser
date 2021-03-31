package runners

//todo: think good name

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"code-analyser/utils"
	"context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

type Dependency struct {
	Name    string
	Version string
}

const (Framework ="FRAMEWORK" )



func DetectRuntime(ctx context.Context, path string, yamlLangObject *protos.LanguageVersion) string {
	runtimeResponse, client := pluginClient.DetectRuntimePluginCall(exec.Command("sh", "-c", yamlLangObject.Runtimeversions.Detector))
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

func ParseLangaugeYML(filePath string) *protos.LanguageVersion {
	path, _ := filepath.Abs(filePath)

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		utils.Logger(err, "ERROR")
		return nil
	}
	var lang protos.LanguageVersion
	err = yaml.Unmarshal(yamlFile, &lang)
	if err != nil {
		utils.Logger(err)
		return nil
	}
	return &lang
}

func GetParsedDependencis(ctx context.Context, runtimeVersion, path string, langYamlObject *protos.LanguageVersion) *protos.LanguageVersion {
	AllDependencies:= map[string]map[string]*protos.PluginSemver{}

	var dependenciesCommand *protos.PluginSemver
	for _, supportedRuntimeVersions := range langYamlObject.Runtimeversions.Versions {
		if helpers.SeverValidate(supportedRuntimeVersions.Semver, runtimeVersion) {
			dependenciesCommand = supportedRuntimeVersions
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
		log.Println(dependenciesList)

		AllDependencies[Framework]=ParseFrameworkFromDependencies(dependenciesList,langYamlObject)
		log.Println(AllDependencies[Framework]["beego"])

	}

	return nil

}


