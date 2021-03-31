package runners

//todo: think good name

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"code-analyser/utils"
	"context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
)

type Dependency struct {
	Name    string
	Version string
}

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
	path := "/home/deqode/Documents/code-analyser" + filePath
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		utils.Logger(err,"ERROR")
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

func GetParsedDependencis(ctx context.Context, runtimeVersion, path string, langYamlObject *protos.LanguageVersion) {
	//this will firstly fetch version of repo used
	//	and ParseLangaugeYML
	//	check and segregate dependencies accordingly

}
