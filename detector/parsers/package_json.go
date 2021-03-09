package parsers

import (
	"code-analyser/detector"
	"encoding/json"
)

type PackageJson struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Scripts         map[string]string `json:"scripts"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Engines         map[string]string `json:"engines"`
}

func ParsePackageJSON(pjson string) (*PackageJson, error) {
	file, err := detector.ReadFile(pjson)
	if err != nil {
		return nil, err
	}
	var res PackageJson
	json.Unmarshal([]byte(file), &res)
	return &res, err
}
