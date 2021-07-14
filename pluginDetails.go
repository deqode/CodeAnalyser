package main

import (
	utilsPB "code-analyser/protos/pb/output/utils"
	pluginDetailsPB "code-analyser/protos/pb/pluginDetails"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)






//TODO error bhi return karni hai
//GetLanguagePluginsPath It will read pluginDetails.yaml files in plugin directory of given languagePath
//func GetLanguagePluginsPath(ctx context.Context, languagePluginsRootPath string) *pluginDetailsPB.LanguagePlugins {
//
//	pluginYamlFiles, err := SearchFileInDirectory("pluginDetails.yaml", languagePluginsRootPath)
//	if err != nil {
//		utils.Logger(err)
//		return nil
//	}
//
//	languagePlugins := &pluginDetailsPB.LanguagePlugins{
//		Frameworks:      map[string]*pluginDetailsPB.DependencyDetails{},
//		Databases:       map[string]*pluginDetailsPB.DependencyDetails{},
//		Orms:            map[string]*pluginDetailsPB.DependencyDetails{},
//		Libraries:       map[string]*pluginDetailsPB.DependencyDetails{},
//		RuntimeVersions: map[string]*pluginDetailsPB.DependencyVersionDetails{},
//	}
//
//	//TODO: discuss with Atul better ways to implement concurrency
//	for _, pluginFile := range pluginYamlFiles {
//
//		parsedRawFile, err := ReadPluginYamlFile(ctx, pluginFile)
//		if err != nil {
//			log.Println(err)
//			//TODO remove this continue , on error whole function should stop
//			continue
//		}
//
//		pluginYamlFile := parsedRawFile.PluginDetails
//
//		switch pluginYamlFile.Type {
//		case "detectRuntime":
//			languagePlugins.DetectRuntimePluginPath = pluginYamlFile.Command
//		case "commands":
//			languagePlugins.CommandsPluginPath = pluginYamlFile.Command
//		case "env":
//			languagePlugins.EnvPluginPath = pluginYamlFile.Command
//		case "preDetectCommand":
//			languagePlugins.PreDetectCommandPluginPath = pluginYamlFile.Command
//		case "staticAssets":
//			languagePlugins.StaticAssetsPluginPath = pluginYamlFile.Command
//		case "buildDirectory":
//			languagePlugins.BuildDirectoryPluginPath = pluginYamlFile.Command
//		case "testCasesCommands":
//			languagePlugins.TestCasesCommandPluginPath = pluginYamlFile.Command
//
//		case "framework":
//			if framework, ok := languagePlugins.Frameworks[pluginYamlFile.Name]; ok {
//				AddPluginVersionDetails(framework, pluginYamlFile)
//			} else {
//				languagePlugins.Frameworks[pluginYamlFile.Name] = CreatePluginVersionMap(pluginYamlFile)
//			}
//
//		case "orm":
//			if orm, ok := languagePlugins.Orms[pluginYamlFile.Name]; ok {
//				AddPluginVersionDetails(orm, pluginYamlFile)
//			} else {
//				languagePlugins.Orms[pluginYamlFile.Name] = CreatePluginVersionMap(pluginYamlFile)
//			}
//
//		case "library":
//			if library, ok := languagePlugins.Libraries[pluginYamlFile.Name]; ok {
//				AddPluginVersionDetails(library, pluginYamlFile)
//			} else {
//				languagePlugins.Libraries[pluginYamlFile.Name] = CreatePluginVersionMap(pluginYamlFile)
//			}
//
//		case "database":
//			if database, ok := languagePlugins.Databases[pluginYamlFile.Name]; ok {
//				AddPluginVersionDetails(database, pluginYamlFile)
//			} else {
//				languagePlugins.Databases[pluginYamlFile.Name] = CreatePluginVersionMap(pluginYamlFile)
//			}
//
//		case "getDependencies":
//			languagePlugins.RuntimeVersions[pluginYamlFile.Version] = &pluginDetailsPB.DependencyVersionDetails{
//				Semver:     pluginYamlFile.Semver,
//				PluginPath: pluginYamlFile.Command,
//			}
//		}
//
//	}
//	return languagePlugins
//}

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
