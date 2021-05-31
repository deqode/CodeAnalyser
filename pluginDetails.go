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

// PluginDetailsFile stores path of pluginDetails.yaml and dir of pluginDetails.yaml
type PluginDetailsFile struct {
	path string
	dir  string
}

//ReadPluginYamlFile It wil read yaml file of specific plugin
func ReadPluginYamlFile(filePath PluginDetailsFile) (*utilsPB.Plugin, error) {
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

//LoadGlobalFilesPluginInfo It will read pluginDetails.yaml files in plugin/globalDetectors directory
func LoadGlobalFilesPluginInfo(globalPath string) *versionsPB.GlobalPlugin {
	var pluginDetailsFiles []PluginDetailsFile
	err := filepath.Walk(globalPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == "pluginDetails.yaml" {
				dir := strings.Split(path, info.Name())[0]
				pluginDetailsFiles = append(pluginDetailsFiles, PluginDetailsFile{path: path, dir: dir})
			}
			return nil
		})
	if err != nil {
		utils.Logger(err)
	}
	var wg sync.WaitGroup
	//TODO: discuss with Atul better ways to implement concurrency
	globalPlugin := &versionsPB.GlobalPlugin{}
	for _, pluginFile := range pluginDetailsFiles {
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

//LoadLanguageSpecificPluginInfo It will read pluginDetails.yaml files in plugin directory of given languagePath
func LoadLanguageSpecificPluginInfo(languagePath string) *versionsPB.LanguageVersion {
	var pluginDetailsFiles []PluginDetailsFile
	err := filepath.Walk(languagePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == "pluginDetails.yaml" {
				dir := strings.Split(path, info.Name())[0]
				pluginDetailsFiles = append(
					pluginDetailsFiles,
					PluginDetailsFile{
						path: path,
						dir:  dir,
					},
				)
			}
			return nil
		})
	if err != nil {
		utils.Logger(err)
	}
	languagePluginInfo := &versionsPB.LanguageVersion{
		Detectruntimecommand: "",
		Runtimeversions:      nil,
		Framework:            map[string]*versionsPB.DependencyDetails{},
		Databases:            map[string]*versionsPB.DependencyDetails{},
		Orms:                 map[string]*versionsPB.DependencyDetails{},
		Libraries:            map[string]*versionsPB.DependencyDetails{},
		//TODO:Not implemented yet
		Dependencies:           nil,
		DetectEnvCommand:       "",
		PreDetectCommand:       "",
		StaticAssetsCommand:    "",
		BuildDirectoryCommand:  "",
		DetectTestCasesCommand: "",
		Commands:               "",
	}
	var wg sync.WaitGroup
	//TODO: discuss with Atul better ways to implement concurrency
	for _, pluginFile := range pluginDetailsFiles {
		wg.Add(1)
		pluginFile := pluginFile
		go func() {
			defer wg.Done()
			parsedRawFile, _ := ReadPluginYamlFile(pluginFile)
			if parsedRawFile != nil {
				parsedFile := parsedRawFile.PluginDetails
				switch parsedFile.Type {
				case "framework":
					if val, ok := languagePluginInfo.Framework[parsedFile.Name]; ok {
						val.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}
					} else {
						languagePluginInfo.Framework[parsedFile.Name] = &versionsPB.DependencyDetails{
							Version: map[string]*versionsPB.DependencyVersionDetails{
								parsedFile.Version: {
									Semver:        parsedFile.Semver,
									Plugincommand: parsedFile.Command,
									Libraries:     parsedFile.Libraries,
								},
							},
						}
					}
				case "detectRuntime":
					languagePluginInfo.Detectruntimecommand = parsedFile.Command
				case "commands":
					languagePluginInfo.Commands = parsedFile.Command
				case "env":
					languagePluginInfo.DetectEnvCommand = parsedFile.Command
				case "preDetectCommand":
					languagePluginInfo.PreDetectCommand = parsedFile.Command
				case "staticAssets":
					languagePluginInfo.StaticAssetsCommand = parsedFile.Command
				case "buildDirectory":
					languagePluginInfo.BuildDirectoryCommand = parsedFile.Command
				case "testCasesCommands":
					languagePluginInfo.DetectTestCasesCommand = parsedFile.Command
				case "orm":
					if val, ok := languagePluginInfo.Orms[parsedFile.Name]; ok {
						val.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}
					} else {
						languagePluginInfo.Orms[parsedFile.Name] = &versionsPB.DependencyDetails{
							Version: map[string]*versionsPB.DependencyVersionDetails{
								parsedFile.Version: {Semver: parsedFile.Semver,
									Plugincommand: parsedFile.Command,
									Libraries:     parsedFile.Libraries,
								},
							},
						}
					}
				case "library":
					if val, ok := languagePluginInfo.Libraries[parsedFile.Name]; ok {
						val.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}
					} else {
						languagePluginInfo.Libraries[parsedFile.Name] = &versionsPB.DependencyDetails{
							Version: map[string]*versionsPB.DependencyVersionDetails{
								parsedFile.Version: {
									Semver:        parsedFile.Semver,
									Plugincommand: parsedFile.Command,
									Libraries:     parsedFile.Libraries,
								},
							},
						}
					}
				case "database":
					if val, ok := languagePluginInfo.Databases[parsedFile.Name]; ok {
						val.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
							Libraries:     parsedFile.Libraries,
						}
					} else {
						languagePluginInfo.Databases[parsedFile.Name] = &versionsPB.DependencyDetails{
							Version: map[string]*versionsPB.DependencyVersionDetails{
								parsedFile.Version: {Semver: parsedFile.Semver,
									Plugincommand: parsedFile.Command,
									Libraries:     parsedFile.Libraries,
								},
							},
						}
					}
				case "getDependencies":
					if languagePluginInfo.Runtimeversions != nil {
						languagePluginInfo.Runtimeversions[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
						}
					} else {
						languagePluginInfo.Runtimeversions = map[string]*versionsPB.DependencyVersionDetails{
							parsedFile.Version: {
								Semver:        parsedFile.Semver,
								Plugincommand: parsedFile.Command,
							},
						}
					}
				}
			}
		}()

	}
	wg.Wait()
	return languagePluginInfo
}

const SupportedLanguages = "./static/supportedLanguages.yaml"

//SupportedLanguagesParser it reads yaml file and fetch out supported languages by our system (like go or js )
func SupportedLanguagesParser() (*versionsPB.SupportedLanguages, error) {
	filename, _ := filepath.Abs(SupportedLanguages)
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
