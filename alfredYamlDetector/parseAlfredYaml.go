package alfredYmalDetector

import (
	protos "code-analyser/protos/pb"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func ParserAlfredYmal(path string) (*protos.AlfredYamlOutput, error) {
	filename, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	var lang protos.AlfredYamlOutput
	err = yaml.Unmarshal(yamlFile, &lang)
	if err != nil {
		return nil, err
	}
	return &lang, nil
}
