package main

import (
	"code-analyser/analyser"
	"code-analyser/helpers"
	"code-analyser/protos/protos"
	"code-analyser/runners"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	path := "./"
	Scrape(path)
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
