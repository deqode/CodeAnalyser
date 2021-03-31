package main

import (
	"code-analyser/analyser"
	"code-analyser/protos/protos"
	"code-analyser/runners"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	log.Println(runners.ParseLangaugeYML("./static/go.yaml"))
	//path := "./"
	//Scrape(path)
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
		langYamlObject := runners.ParseLangaugeYML(languagePath)
		runtimeVersion := runners.DetectRuntime(nil, path, langYamlObject)

		runners.GetParsedDependencis(nil, runtimeVersion, path, langYamlObject)
		runners.RunDetectors(langYamlObject, runtimeVersion, path)

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
