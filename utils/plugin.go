package utils

import (
	utilsPB "code-analyser/protos/pb/output/utils"
	"errors"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

//CallPluginCommand it takes command and make it executable for bash
func CallPluginCommand(command string) *exec.Cmd {
	return exec.Command("sh", "-c", command)
}

// FileDetails stores path of pluginDetails.yaml and dir of pluginDetails.yaml
type FileDetails struct {
	path string
	dir  string
}

//ReadPluginYamlFile It wil read yaml file of specific plugin
func ReadPluginYamlFile(ctx context.Context, filePath FileDetails) (*utilsPB.Plugin, error) {
	filename, _ := filepath.Abs(filePath.path)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var lang utilsPB.Plugin

	err = yaml.Unmarshal(yamlFile, &lang)
	if err != nil {
		return nil, err
	}
	if lang.PluginDetails.Command != "" && strings.Contains(lang.PluginDetails.Command, "alfredPlugin/") {
		lang.PluginDetails.Command = strings.Replace(lang.PluginDetails.Command, "alfredPlugin/", filePath.dir, 1)
	} else {
		return nil, errors.New("invalid command in plugin")
	}
	return &lang, nil
}

//Dependency any external dependency with exact version for e.g. docker 3.2
type Dependency struct {
	Name    string
	Version string
}

//Library any external library and its semver(from version x to version y) enry ^1.4.2
type Library struct {
	Name   string
	Semver []string
}
