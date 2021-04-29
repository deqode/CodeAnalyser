package main

import (
	utilsPB "code-analyser/protos/pb/output/utils"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

//ReadPluginYamlFile It wil read yaml file of specific plugin
func ReadPluginYamlFile(filePath struct {
	path string
	dir  string
}) (*utilsPB.Plugin, error) {
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

func ParseGlobalPluginYaml(globalPath string) *versionsPB.GlobalPlugin {
	var pluginDetailsFileLst []struct {
		path string
		dir  string
	}
	err := filepath.Walk(globalPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == "pluginDetails.yaml" {
				dir := strings.Split(path, info.Name())[0]
				pluginDetailsFileLst = append(pluginDetailsFileLst, struct {
					path string
					dir  string
				}{path: path, dir: dir})
			}
			return nil
		})
	if err != nil {
		utils.Logger(err)
	}
	var wg sync.WaitGroup
	globalPlugin := &versionsPB.GlobalPlugin{}
	for _, pluginFile := range pluginDetailsFileLst {
		wg.Add(1)
		pluginFile := pluginFile
		go func() {
			defer wg.Done()

			parsedRawFile, _ := ReadPluginYamlFile(pluginFile)
			if parsedRawFile != nil {
				parsedFile := parsedRawFile.PluginDetails
				switch parsedFile.Type {
				case "dockerFile":
					globalPlugin.DockerFile = parsedFile.Command
				case "procFile":
					globalPlugin.ProcFile = parsedFile.Command
				case "makeFile":
					globalPlugin.MakeFile = parsedFile.Command
				case "commands":
					globalPlugin.Commands = parsedFile.Command
				}
			}
		}()

	}
	wg.Wait()
	return globalPlugin
}

/*ParsePluginYamlFile it will fetch paths of all plugins using walk function and
parse all dependencies to their categories for example postgres is a Db
*/
func ParsePluginYamlFile(rootPath string) *versionsPB.LanguageVersion {
	var pluginDetailsFileLst []struct {
		path string
		dir  string
	}
	err := filepath.Walk(rootPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == "pluginDetails.yaml" {
				dir := strings.Split(path, info.Name())[0]
				pluginDetailsFileLst = append(pluginDetailsFileLst, struct {
					path string
					dir  string
				}{path: path, dir: dir})
			}
			return nil
		})
	if err != nil {
		utils.Logger(err)
	}
	languageVersion := &versionsPB.LanguageVersion{}
	var wg sync.WaitGroup

	for _, pluginFile := range pluginDetailsFileLst {
		wg.Add(1)
		pluginFile := pluginFile
		go func() {
			defer wg.Done()

			parsedRawFile, _ := ReadPluginYamlFile(pluginFile)
			if parsedRawFile != nil {
				parsedFile := parsedRawFile.PluginDetails
				switch parsedFile.Type {
				case "framework":
					if val, ok := languageVersion.Framework[parsedFile.Name]; ok {
						val.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}
					} else {
						dependencyDetails := versionsPB.DependencyDetails{Version: map[string]*versionsPB.DependencyVersionDetails{}}
						dependencyDetails.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}
						if languageVersion.Framework == nil {
							languageVersion.Framework = map[string]*versionsPB.DependencyDetails{parsedFile.Name: &dependencyDetails}
						} else {
							languageVersion.Framework[parsedFile.Name] = &dependencyDetails
						}
					}
				case "detectRuntime":
					languageVersion.Detectruntimecommand = parsedFile.Command
				case "commands":
					languageVersion.Commands = parsedFile.Command
				case "env":
					languageVersion.DetectEnvCommand = parsedFile.Command
				case "preDetectCommand":
					languageVersion.PreDetectCommand = parsedFile.Command
				case "staticAssets":
					languageVersion.StaticAssetsCommand = parsedFile.Command
				case "buildDirectory":
					languageVersion.BuildDirectoryCommand = parsedFile.Command
				case "testCasesCommands":
					languageVersion.DetectTestCasesCommand = parsedFile.Command
				case "orm":
					if val, ok := languageVersion.Orms[parsedFile.Name]; ok {
						val.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}
					} else {
						dependencyDetails := versionsPB.DependencyDetails{Version: map[string]*versionsPB.DependencyVersionDetails{}}
						dependencyDetails.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}

						if languageVersion.Orms == nil {
							languageVersion.Orms = map[string]*versionsPB.DependencyDetails{parsedFile.Name: &dependencyDetails}
						} else {
							languageVersion.Orms[parsedFile.Name] = &dependencyDetails
						}

					}
				case "library":
					if val, ok := languageVersion.Libraries[parsedFile.Name]; ok {
						val.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}
					} else {
						dependencyDetails := versionsPB.DependencyDetails{Version: map[string]*versionsPB.DependencyVersionDetails{}}
						dependencyDetails.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}
						if languageVersion.Libraries == nil {
							languageVersion.Libraries = map[string]*versionsPB.DependencyDetails{parsedFile.Name: &dependencyDetails}
						} else {
							languageVersion.Libraries[parsedFile.Name] = &dependencyDetails
						}
					}
				case "database":
					if val, ok := languageVersion.Databases[parsedFile.Name]; ok {
						val.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}
					} else {
						dependencyDetails := versionsPB.DependencyDetails{Version: map[string]*versionsPB.DependencyVersionDetails{}}
						dependencyDetails.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}
						if languageVersion.Databases == nil {
							languageVersion.Databases = map[string]*versionsPB.DependencyDetails{parsedFile.Name: &dependencyDetails}
						} else {
							languageVersion.Databases[parsedFile.Name] = &dependencyDetails
						}
					}
				case "getDependencies":
					if languageVersion.Runtimeversions != nil {
						languageVersion.Runtimeversions[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
						}
					} else {
						languageVersion.Runtimeversions = map[string]*versionsPB.DependencyVersionDetails{parsedFile.Version: {
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
						}}
					}
				}
			}
		}()

	}
	wg.Wait()

	return languageVersion
}

//SupportedLanguagedParser it reads yaml file and fetch out supported languges by our system
func SupportedLanguagedParser() (*versionsPB.SupportedLanguages, error) {
	filename, _ := filepath.Abs("./static/supportedLanguages.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var lang versionsPB.SupportedLanguages

	err = yaml.Unmarshal(yamlFile, &lang)
	if err != nil {
		return nil, err
	}
	return &lang, nil
}
