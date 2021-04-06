package main

import (
	"code-analyser/analyser"
	"code-analyser/protos/protos"
	utilsPB "code-analyser/protos/protos/outputs/utils"
	versionsPB "code-analyser/protos/protos/versions"
	"code-analyser/runners"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := "./"
	Scrape(path)
}
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
		log.Println(err)
	}
	var languageVersion versionsPB.LanguageVersion
	for _, pluginFile := range pluginDetailsFileLst {
		parsedRawFile, _ := ReadPluginYamlFile(pluginFile)
		if parsedRawFile == nil {
			continue
		}
		parsedFile := parsedRawFile.PluginDetails
		switch parsedFile.Type {
		case "framework":
			if val, ok := languageVersion.Framework[parsedFile.Name]; ok {
				val.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
			} else {
				dependencyDetails := versionsPB.DependencyDetails{Version: map[string]*versionsPB.DependencyVersionDetails{}}
				dependencyDetails.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
				languageVersion.Framework = map[string]*versionsPB.DependencyDetails{parsedFile.Name: &dependencyDetails}
			}
			break
		case "detectRuntime":
			languageVersion.Detectruntimecommand = parsedFile.Command
			break
		case "env":
			languageVersion.DetectenvCommand = parsedFile.Command
			break
		case "orm":
			if val, ok := languageVersion.Orms[parsedFile.Name]; ok {
				val.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
			} else {
				dependencyDetails := versionsPB.DependencyDetails{Version: map[string]*versionsPB.DependencyVersionDetails{}}
				dependencyDetails.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
				languageVersion.Orms = map[string]*versionsPB.DependencyDetails{parsedFile.Name: &dependencyDetails}
			}
			break
		case "database":
			if val, ok := languageVersion.Databases[parsedFile.Name]; ok {
				val.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
			} else {
				dependencyDetails := versionsPB.DependencyDetails{Version: map[string]*versionsPB.DependencyVersionDetails{}}
				dependencyDetails.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
					Semver:        parsedFile.Semver,
					Plugincommand: parsedFile.Command,
				}
				languageVersion.Databases = map[string]*versionsPB.DependencyDetails{parsedFile.Name: &dependencyDetails}
			}
			break
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
	return &languageVersion
}

//Scrape it scrape language, framework, orm etc .....
func Scrape(path string) {
	languages, _, _ := analyser.Analyse(path)
	supportedLanguages, _ := SupportedLanguagedParser()
	decisionMakerInput := &protos.DecisionMakerInput{
		LanguageSpecificDetection: []*protos.LanguageSpecificDetections{},
		GloabalDetections:         nil,
	}
	for _, language := range languages {
		var languagePath string
		for _, supportedlanguage := range supportedLanguages.Languages {
			if supportedlanguage.Name == language.Name {
				languagePath = supportedlanguage.Path
				break
			}
		}
		if languagePath != "" {
			yamlLangObject := ParsePluginYamlFile(languagePath)
			runtimeVersion := runners.DetectRuntime(nil, path, yamlLangObject)
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

//RunAllDetectors it runs all detectors of dependencies ex. orm,framework etc ....
func RunAllDetectors(languageSpecificDetections *protos.LanguageSpecificDetections, allDependencies map[string]map[string]runners.DependencyDetail, runtimeVersion string, path string) {
	languageSpecificDetections.Orm = runners.OrmRunner(allDependencies[runners.ORM], runtimeVersion, path)
	languageSpecificDetections.Db = runners.DbRunner(allDependencies[runners.DB], runtimeVersion, path)
	languageSpecificDetections.Framework = runners.FrameworkRunner(allDependencies[runners.Framework], runtimeVersion, path)
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
