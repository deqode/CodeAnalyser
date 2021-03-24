package versions

import (
	protos "code-analyser/protos/pb"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

const (
	GO = "./static/versionFile/go.yaml"
)

func ParserVersion(file string) (*protos.LanguageVersion, error) {
	filename, _ := filepath.Abs("/home/deqode/Documents/code-analyser/versions/" + file)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var lang protos.LanguageVersion

	err = yaml.Unmarshal(yamlFile, &lang)
	if err != nil {
		return nil, err
	}
	return &lang, nil

}
