package main

import (
	"code-analyser/analyser"
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"code-analyser/runners"
	"code-analyser/utils"
	"io/ioutil"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	path := "./"
	Scrape(path)
	dbResponse, client := pluginClient.DbPluginCall(exec.Command("sh", "-c", "go run plugin/go/db/postgres/V_1_X/main.go"))
	defer client.Kill()
	utils.Logger(dbResponse.PercentOfDbUsed(&pb.ServiceInput{
		RuntimeVersion: "fdsfds",
		Root:           "dfsdf",
	}))
	utils.Logger(dbResponse.Detect(&pb.ServiceInput{
		RuntimeVersion: "fdsfds",
		Root:           "dfsdf",
	}))
}

func Scrape(path string) {
	languages, _, _ := analyser.Analyse(path)
	supportedLanguages, _ := supportedLanguagedParser()
	for _, language := range languages {
		var languagePath string
		for _, Supportedlanguage := range supportedLanguages.Languages {
			if Supportedlanguage.Name == language.Name {
				languagePath = Supportedlanguage.Path
				break
			}
		}
		if languagePath != "" {
			yamlLangObject := runners.ParseLangaugeYML(languagePath)
			runtimeVersion := (runners.DetectRuntime(nil, path, yamlLangObject))
			if runtimeVersion == "" {
				break
			}
			runners.GetParsedDependencis(nil, runtimeVersion, path, yamlLangObject)
		}

	}
}

func supportedLanguagedParser() (*protos.SupportedLanguages, error) {
	filename, _ := filepath.Abs("./static/supportedLanguages.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var lang protos.SupportedLanguages

	err = yaml.Unmarshal(yamlFile, &lang)
	if err != nil {
		return nil, err
	}
	return &lang, nil
}
