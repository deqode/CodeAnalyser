package main

import (
	"code-analyser/analyser"
	decisionmakerPB "code-analyser/protos/pb"
	utilsPB "code-analyser/protos/pb/output/utils"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/runners"
	"code-analyser/utils"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
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
						}
					} else {
						dependencyDetails := versionsPB.DependencyDetails{Version: map[string]*versionsPB.DependencyVersionDetails{}}
						dependencyDetails.Version[parsedFile.Version] = &versionsPB.DependencyVersionDetails{
							Semver:        parsedFile.Semver,
							Plugincommand: parsedFile.Command,
						}
						languageVersion.Framework = map[string]*versionsPB.DependencyDetails{parsedFile.Name: &dependencyDetails}
					}
				case "detectRuntime":
					languageVersion.Detectruntimecommand = parsedFile.Command
				case "env":
					languageVersion.DetectEnvCommand = parsedFile.Command
				case "preDetectCommand":
					languageVersion.PreDetectCommand = parsedFile.Command
				case "staticAssets":
					languageVersion.StaticAssetsCommand = parsedFile.Command
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
				case "library":
					if val, ok := languageVersion.Libraries[parsedFile.Name]; ok {
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
						languageVersion.Libraries = map[string]*versionsPB.DependencyDetails{parsedFile.Name: &dependencyDetails}
					}
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

//Scrape it scrape language, framework, orm etc .....
func Scrape(path string) {
	languages, _, _ := analyser.Analyse(path)
	supportedLanguages, _ := SupportedLanguagedParser()
	decisionMakerInput := &decisionmakerPB.DecisionMakerInput{
		LanguageSpecificDetection: []*decisionmakerPB.LanguageSpecificDetections{},
		GloabalDetections:         nil,
	}
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}


	for _, language := range languages {
		wg.Add(1)
		language := language
		go func() {
			defer wg.Done()
			var languagePath string
			for _, supportedLanguage := range supportedLanguages.Languages {
				if supportedLanguage.Name == language.Name {
					languagePath = supportedLanguage.Path
					break
				}
			}
			if languagePath != "" {
				pluginDetails := ParsePluginYamlFile(languagePath)
				runtimeVersion := runners.DetectRuntime(nil, path, pluginDetails)
				runners.RunPreDetectCommand(nil, &pb.ServiceInput{
					RuntimeVersion: runtimeVersion,
					Root:           path,
				}, pluginDetails)
				if runtimeVersion != "" {
					allDependencies := runners.GetParsedDependencis(nil, runtimeVersion, path, pluginDetails)
					languageSpecificDetections := decisionmakerPB.LanguageSpecificDetections{
						Name:           language.Name,
						RuntimeVersion: runtimeVersion,
					}
					RunAllDetectors(&languageSpecificDetections, allDependencies, pluginDetails, runtimeVersion, path)
					mutex.Lock()
					decisionMakerInput.LanguageSpecificDetection = append(decisionMakerInput.LanguageSpecificDetection, &languageSpecificDetections)
					mutex.Unlock()

					log.Println(decisionMakerInput)
				}

			}
		}()
	}
	wg.Wait()

}

//RunAllDetectors it runs all detectors of dependencies ex. orm,framework etc ....
func RunAllDetectors(languageSpecificDetections *decisionmakerPB.LanguageSpecificDetections, allDependencies map[string]map[string]runners.DependencyDetail, pluginDetails *versionsPB.LanguageVersion, runtimeVersion string, path string) {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	wg.Add(6)
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Orm = runners.OrmRunner(allDependencies[runners.ORM], runtimeVersion, path)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Db = runners.DbRunner(allDependencies[runners.DB], runtimeVersion, path)
		mutex.Unlock()

	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Framework = runners.FrameworkRunner(allDependencies[runners.Framework], runtimeVersion, path)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Env = runners.EnvDetectAndRunner(pluginDetails, runtimeVersion, path)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.StaticAssets = runners.RunStaticAssetsCommand(nil, &pb.ServiceInput{
			RuntimeVersion: runtimeVersion,
			Root:           path,
		}, pluginDetails)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Libraries = runners.LibraryRunner(allDependencies[runners.Library], runtimeVersion, path)
		mutex.Unlock()
	}()

	wg.Wait()

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
