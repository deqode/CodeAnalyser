package main

import (
	utilsPB "code-analyser/protos/pb/output/utils"
	pluginDetailsPB "code-analyser/protos/pb/pluginDetails"
	"code-analyser/utils"
	"errors"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// PluginDetailsFile stores path of pluginDetails.yaml and dir of pluginDetails.yaml
type PluginDetailsFile struct {
	path string
	dir  string
}

//ReadPluginYamlFile It wil read yaml file of specific plugin
func ReadPluginYamlFile(ctx context.Context, filePath PluginDetailsFile) (*utilsPB.Plugin, error) {
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

//SearchFileInDirectory search specific file in directory , it will return path of files
func SearchFileInDirectory(fileName, dirPath string) ([]PluginDetailsFile, error) {
	var pluginYamlFiles []PluginDetailsFile
	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == fileName {
				dir := strings.Split(path, info.Name())[0]
				pluginYamlFiles = append(pluginYamlFiles, PluginDetailsFile{path: path, dir: dir})
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return pluginYamlFiles, err
}

//TODO error return karni hai sath me
//GetGlobalPluginsPath It will read pluginDetails.yaml files in plugin/globalDetectors directory
func GetGlobalPluginsPath(ctx context.Context, globalPluginRootPath string) *pluginDetailsPB.GlobalPlugins {
	pluginYamlFiles, err := SearchFileInDirectory("pluginDetails.yaml", globalPluginRootPath)
	if err != nil {
		utils.Logger(err)
		return nil
	}

	//TODO: discuss with Atul better ways to implement concurrency
	globalPlugin := &pluginDetailsPB.GlobalPlugins{}

	for _, pluginFile := range pluginYamlFiles {

		parsedRawFile, err := ReadPluginYamlFile(ctx, pluginFile)
		if err != nil {
			log.Println(err)
			//TODO remove this continue and return error
			continue
		}

		parsedFile := parsedRawFile.PluginDetails
		switch parsedFile.Type {
		case "dockerFile":
			globalPlugin.DockerFilePluginPath = parsedFile.Command
		case "procFile":
			globalPlugin.ProcFilePluginPath = parsedFile.Command
		case "makeFile":
			globalPlugin.MakeFilePluginPath = parsedFile.Command
		}
	}
	return globalPlugin
}

//TODO error bhi return karni hai
//GetLanguagePluginsPath It will read pluginDetails.yaml files in plugin directory of given languagePath
func GetLanguagePluginsPath(ctx context.Context, languagePluginsRootPath string) *pluginDetailsPB.LanguagePlugins {

	pluginYamlFiles, err := SearchFileInDirectory("pluginDetails.yaml", languagePluginsRootPath)
	if err != nil {
		utils.Logger(err)
		return nil
	}

	languagePlugins := &pluginDetailsPB.LanguagePlugins{
		Frameworks:      map[string]*pluginDetailsPB.DependencyDetails{},
		Databases:       map[string]*pluginDetailsPB.DependencyDetails{},
		Orms:            map[string]*pluginDetailsPB.DependencyDetails{},
		Libraries:       map[string]*pluginDetailsPB.DependencyDetails{},
		RuntimeVersions: map[string]*pluginDetailsPB.DependencyVersionDetails{},
	}

	//TODO: discuss with Atul better ways to implement concurrency
	for _, pluginFile := range pluginYamlFiles {

		parsedRawFile, err := ReadPluginYamlFile(ctx, pluginFile)
		if err != nil {
			log.Println(err)
			//TODO remove this continue , on error whole function should stop
			continue
		}

		pluginYamlFile := parsedRawFile.PluginDetails

		switch pluginYamlFile.Type {
		case "detectRuntime":
			languagePlugins.DetectRuntimePluginPath = pluginYamlFile.Command
		case "commands":
			languagePlugins.CommandsPluginPath = pluginYamlFile.Command
		case "env":
			languagePlugins.EnvPluginPath = pluginYamlFile.Command
		case "preDetectCommand":
			languagePlugins.PreDetectCommandPluginPath = pluginYamlFile.Command
		case "staticAssets":
			languagePlugins.StaticAssetsPluginPath = pluginYamlFile.Command
		case "buildDirectory":
			languagePlugins.BuildDirectoryPluginPath = pluginYamlFile.Command
		case "testCasesCommands":
			languagePlugins.TestCasesCommandPluginPath = pluginYamlFile.Command

		case "framework":
			if framework, ok := languagePlugins.Frameworks[pluginYamlFile.Name]; ok {
				AddPluginVersionDetails(framework, pluginYamlFile)
			} else {
				languagePlugins.Frameworks[pluginYamlFile.Name] = CreatePluginVersionMap(pluginYamlFile)
			}

		case "orm":
			if orm, ok := languagePlugins.Orms[pluginYamlFile.Name]; ok {
				AddPluginVersionDetails(orm, pluginYamlFile)
			} else {
				languagePlugins.Orms[pluginYamlFile.Name] = CreatePluginVersionMap(pluginYamlFile)
			}

		case "library":
			if library, ok := languagePlugins.Libraries[pluginYamlFile.Name]; ok {
				AddPluginVersionDetails(library, pluginYamlFile)
			} else {
				languagePlugins.Libraries[pluginYamlFile.Name] = CreatePluginVersionMap(pluginYamlFile)
			}

		case "database":
			if database, ok := languagePlugins.Databases[pluginYamlFile.Name]; ok {
				AddPluginVersionDetails(database, pluginYamlFile)
			} else {
				languagePlugins.Databases[pluginYamlFile.Name] = CreatePluginVersionMap(pluginYamlFile)
			}

		case "getDependencies":
			languagePlugins.RuntimeVersions[pluginYamlFile.Version] = &pluginDetailsPB.DependencyVersionDetails{
				Semver:     pluginYamlFile.Semver,
				PluginPath: pluginYamlFile.Command,
			}
		}

	}
	return languagePlugins
}

func AddPluginVersionDetails(dependencyMap *pluginDetailsPB.DependencyDetails, pluginYamlFile *utilsPB.Details) {
	dependencyMap.Version[pluginYamlFile.Version] = &pluginDetailsPB.DependencyVersionDetails{
		Semver:     pluginYamlFile.Semver,
		PluginPath: pluginYamlFile.Command,
		Libraries:  pluginYamlFile.Libraries,
	}
}

func CreatePluginVersionMap(pluginYamlFile *utilsPB.Details) *pluginDetailsPB.DependencyDetails {
	return &pluginDetailsPB.DependencyDetails{
		Version: map[string]*pluginDetailsPB.DependencyVersionDetails{
			pluginYamlFile.Version: {
				Semver:     pluginYamlFile.Semver,
				PluginPath: pluginYamlFile.Command,
				Libraries:  pluginYamlFile.Libraries,
			},
		},
	}
}

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
