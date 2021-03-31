package runners

//todo: think good name

import (
	"code-analyser/pluginClient"
	"code-analyser/protos/protos"
	"code-analyser/utils"
	"context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Dependency struct {
	Name    string
	Version string
}

func DetectRuntime(ctx context.Context, path string, langYamlObject *protos.LanguageVersion) string {
	runtime,client:=pluginClient.DetectRuntimeCall()

	defer client.Kill()
	return ""
}

func ParseLangaugeYML(filePath string) *protos.LanguageVersion {
	filename, _ := filepath.Abs(filePath)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		utils.Logger(err)
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
