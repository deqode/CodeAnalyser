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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
)

func main() {
	path := "/home/deqode/Desktop/InstagramDemo"
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
				case "buildDirectory":
					languageVersion.BuildDirectoryCommand = parsedFile.Command
				case "testCasesCommands":
					languageVersion.DetectTestCasesCommand = parsedFile.Command
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
		GloabalDetections:         &decisionmakerPB.GlobalDetections{},
	}
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	var ctx context.Context = nil
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
				globalPlugins := ParseGlobalPluginYaml("./plugin/globalDetectors")
				runtimeVersion := runners.DetectRuntime(ctx, path, pluginDetails)
				runners.RunPreDetectCommand(ctx, &pb.ServiceInput{
					RuntimeVersion: runtimeVersion,
					Root:           path,
				}, pluginDetails)
				if runtimeVersion != "" {
					allDependencies := runners.GetParsedDependencis(ctx, runtimeVersion, path, pluginDetails)
					languageSpecificDetections := decisionmakerPB.LanguageSpecificDetections{
						Name:           language.Name,
						RuntimeVersion: runtimeVersion,
					}
					gloabalDetections := decisionmakerPB.GlobalDetections{}
					commands := decisionmakerPB.Commands{}
					RunAllDetectors(ctx, &languageSpecificDetections, allDependencies, pluginDetails, runtimeVersion, path, &gloabalDetections, globalPlugins, &commands)
					log.Println(allDependencies)
					mutex.Lock()
					decisionMakerInput.LanguageSpecificDetection = append(decisionMakerInput.LanguageSpecificDetection, &languageSpecificDetections)
					decisionMakerInput.GloabalDetections = &gloabalDetections
					decisionMakerInput.Commands = &commands
					mutex.Unlock()
					log.Println(decisionMakerInput.LanguageSpecificDetection)
				}

			}
		}()
	}
	wg.Wait()

}

//RunAllDetectors it runs all detectors of dependencies ex. orm,framework etc ....
func RunAllDetectors(ctx context.Context, languageSpecificDetections *decisionmakerPB.LanguageSpecificDetections, allDependencies map[string]map[string]runners.DependencyDetail, pluginDetails *versionsPB.LanguageVersion, runtimeVersion string, path string, globalDetection *decisionmakerPB.GlobalDetections, globalPlugin *versionsPB.GlobalPlugin, commands *decisionmakerPB.Commands) {
	var wg sync.WaitGroup
	wg.Add(11)
	var mutex = &sync.Mutex{}
	go func() {
		defer wg.Done()
		mutex.Lock()
		globalDetection.DockerFile, globalDetection.DockerComposeFile = runners.DetectDockerAndComposeFile(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Orm = runners.OrmRunner(allDependencies[runners.ORM], runtimeVersion, path)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		commands.SeedCommands, commands.BuildCommands, commands.MigrationCommands, commands.StartUpCommands, commands.AdHocScriptsOutput = runners.DetectAndRunCommands(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		globalDetection.ProcFile = runners.DetectAndRunProcFile(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		globalDetection.Makefile = runners.DetectAndRunMakeFile(ctx, path, globalPlugin)
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
		if pluginDetails.DetectEnvCommand != "" {
			mutex.Lock()
			languageSpecificDetections.Env = runners.EnvDetectAndRunner(pluginDetails, runtimeVersion, path)
			mutex.Unlock()
		}

	}()
	go func() {
		defer wg.Done()
		if pluginDetails.StaticAssetsCommand != "" {
			mutex.Lock()
			languageSpecificDetections.StaticAssets = runners.RunStaticAssetsCommand(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Libraries = runners.LibraryRunner(allDependencies[runners.Library], runtimeVersion, path)
		mutex.Unlock()
	}()
	go func() {
		if pluginDetails.BuildDirectoryCommand != "" {
			mutex.Lock()
			languageSpecificDetections.BuildDirectory = runners.DetectAndRunBuildDirectory(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		if pluginDetails.DetectTestCasesCommand != "" {
			mutex.Lock()
			languageSpecificDetections.TestCases = runners.DetectTestCasesCommand(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
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
