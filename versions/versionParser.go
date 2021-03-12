package versions

import (
	protos "code-analyser/protos/pb"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func ParserVersion(path string) (*protos.LanguageVersion, error) {
	filename, _ := filepath.Abs(path)
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
