package main

import (
	pluginDetailsPB "code-analyser/protos/pb/pluginDetails"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

const SupportedLanguages = "./static/supportedLanguages.yaml"

//SupportedLanguagesParser it reads yaml file and fetch out supported languages by our system (like go or js )
func SupportedLanguagesParser() (*pluginDetailsPB.SupportedLanguages, error) {
	filename, _ := filepath.Abs(SupportedLanguages)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var lang pluginDetailsPB.SupportedLanguages

	err = yaml.Unmarshal(yamlFile, &lang)
	if err != nil {
		return nil, err
	}
	return &lang, nil
}
