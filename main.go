package main

import (
	"code-analyser/analyser"
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"code-analyser/runners"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
//	path := "./"
//	Scrape(path)
	beegoDetector, client := pluginClient.DetectRuntimePluginCall(exec.Command("sh", "-c", " go run plugin/go/detectRuntime/main.go"))
	log.Println(beegoDetector.DetectRuntime(&pb.ServiceInputString{}))
	client.Kill()
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
		log.Println(languagePath)
		yamlObject=runners.ParseLangaugeYML(languagePath)
		runtimeVersion := runners.DetectRuntime(nil, path, yamlObject)

		runners.GetParsedDependencis(runtimeVersion, yamlObject, path )
		runners.RunDetectors(languageDetector, runtimeVersion, path)

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
