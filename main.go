package main

import (
	"code-analyser/analyser"
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"code-analyser/runners"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	path := "./"
	Scrape(path)
	runtimeResponse, client := pluginClient.DependenciesPluginCall(exec.Command("sh", "-c", "go run plugin/go/dependencies/main.go"))
	defer client.Kill()
	log.Println(runtimeResponse.GetDependencies(&pb.ServiceInput{
		RuntimeVersion: "243",
		Root:           "./",
	}))
	helpers.SeverValidate(">=1,<=3", "1.2")
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
			//runners.GetParsedDependencis(nil, path, runtimeVersion, yamlLangObject)
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
