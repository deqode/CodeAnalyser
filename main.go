package main

import (
	"code-analyser/analyser"
	"code-analyser/protos/protos"
	"code-analyser/runners"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	path := "./"
	Scrape(path)
}

func ReadPluginYamlFile(path string) (*protos.Plugin, error) {
	filename, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var lang protos.Plugin

	err = yaml.Unmarshal(yamlFile, &lang)
	if err != nil {
		return nil, err
	}
	return &lang, nil
}

func ParsePluginYamlFile(rootPath string) *protos.LanguageVersion {
	var pluginDetailsFileLst []string
	//TODO make path dynamic from supported language
	err := filepath.Walk(rootPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == "pluginDetails.yaml" {
				pluginDetailsFileLst = append(pluginDetailsFileLst, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	var languageVersion protos.LanguageVersion
	for _, pluginFile := range pluginDetailsFileLst {
		parsedRawFile, _ := ReadPluginYamlFile(pluginFile)
		if parsedRawFile == nil {
			continue
		}
		parsedFile := parsedRawFile.PluginDetails
		switch parsedFile.Type {
		case "framework":
			if val, ok := languageVersion.Framework[parsedFile.Name]; ok {
				val.Version[parsedFile.Version] = &protos.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
			} else {
				dependencyDetails := protos.DependencyDetails{Version: map[string]*protos.DependencyVersionDetails{}}
				dependencyDetails.Version[parsedFile.Version] = &protos.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
				languageVersion.Framework = map[string]*protos.DependencyDetails{parsedFile.Name: &dependencyDetails}
			}
			break
		case "detectRuntime":
			languageVersion.Detectruntimecommand = parsedFile.Command
			break
		case "orm":
			if val, ok := languageVersion.Orms[parsedFile.Name]; ok {
				val.Version[parsedFile.Version] = &protos.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
			} else {
				dependencyDetails := protos.DependencyDetails{Version: map[string]*protos.DependencyVersionDetails{}}
				dependencyDetails.Version[parsedFile.Version] = &protos.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
				languageVersion.Orms = map[string]*protos.DependencyDetails{parsedFile.Name: &dependencyDetails}
			}
			break
		case "database":
			if val, ok := languageVersion.Databases[parsedFile.Name]; ok {
				val.Version[parsedFile.Version] = &protos.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
			} else {
				dependencyDetails := protos.DependencyDetails{Version: map[string]*protos.DependencyVersionDetails{}}
				dependencyDetails.Version[parsedFile.Version] = &protos.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
				languageVersion.Databases = map[string]*protos.DependencyDetails{parsedFile.Name: &dependencyDetails}
			}
			break
		case "getDependencies":
			if languageVersion.Runtimeversions != nil {
				languageVersion.Runtimeversions[parsedFile.Version] = &protos.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
			} else {
				languageVersion.Runtimeversions = map[string]*protos.DependencyVersionDetails{parsedFile.Version: {
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}}
			}
		}
	}
	return &languageVersion
}

func Scrape(path string) {
	languages, _, _ := analyser.Analyse(path)
	supportedLanguages, _ := supportedLanguagedParser()
	decisionMakerInput := &protos.DecisionMakerInput{
		LanguageSpecificDetection: []*protos.LanguageSpecificDetections{},
		GloabalDetections:         nil,
	}
	for _, language := range languages {
		var languagePath string
		for _, Supportedlanguage := range supportedLanguages.Languages {
			if Supportedlanguage.Name == language.Name {
				languagePath = Supportedlanguage.Path
				break
			}
		}
		if languagePath != "" {
			yamlLangObject := ParsePluginYamlFile(languagePath)
			runtimeVersion := (runners.DetectRuntime(nil, path, yamlLangObject))
			if runtimeVersion == "" {
				break
			}
			allDependencies := runners.GetParsedDependencis(nil, runtimeVersion, path, yamlLangObject)
			languageSpecificDetections := protos.LanguageSpecificDetections{
				Name:           language.Name,
				RuntimeVersion: runtimeVersion,
			}
			RunAllDetectors(&languageSpecificDetections, allDependencies, runtimeVersion, path)
			decisionMakerInput.LanguageSpecificDetection = append(decisionMakerInput.LanguageSpecificDetection, &languageSpecificDetections)
			log.Println(decisionMakerInput)
		}

	}
}

func RunAllDetectors(languageSpecificDetections *protos.LanguageSpecificDetections, allDependencies map[string]map[string]runners.DependencyDetail, runtimeVersion, path string) {
	languageSpecificDetections.Orm = runners.OrmRunner(allDependencies[runners.ORM], runtimeVersion, path)
	languageSpecificDetections.Db = runners.DbRunner(allDependencies[runners.DB], runtimeVersion, path)
	languageSpecificDetections.Framework = runners.FrameworkRunner(allDependencies[runners.Framework], runtimeVersion, path)
}

func supportedLanguagedParser() (*protos.SupportedLanguages, error) {
	filename, _ := filepath.Abs("./static/supportedLanguages.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var lang protos.SupportedLanguages

	err = yaml.Unmarshal(yamlFile, &lang)
	if err != nil {
		return nil, err
	}
	return &lang, nil
}
