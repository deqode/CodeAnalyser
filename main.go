package main

import (
	"code-analyser/analyser"
	decisionmakerPB "code-analyser/protos/pb"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/runners"
	"log"
	"math"
	"sync"

	"golang.org/x/net/context"
)

func main() {
	//path := os.Args[1]
	//log.Println("Initialized Scrapping ")
	//log.Println("Scrapping on "+path)
	decisionMakerInput := Scrape("/home/deqode/Desktop/project/go/code-analyser")
	log.Println(decisionMakerInput.LanguageSpecificDetection)
	//res, err := json.MarshalIndent(decisionMakerInput.LanguageSpecificDetection[0], "", "  ")
	//if err != nil {
	//	log.Println("error:", err)
	//}
	//fmt.Print(string(res))
	log.Println(decisionMakerInput.GloabalDetections)
}

//Scrape it scrape language, framework, orm etc .....
func Scrape(path string) *decisionmakerPB.DecisionMakerInput {
	languages, _, _ := analyser.Analyse(path)
	supportedLanguages, _ := SupportedLanguagedParser()
	decisionMakerInput := &decisionmakerPB.DecisionMakerInput{
		LanguageSpecificDetection: []*decisionmakerPB.LanguageSpecificDetections{},
		GloabalDetections:         &decisionmakerPB.GlobalDetections{},
	}
	var wg sync.WaitGroup
	var ctx context.Context = nil
	globalPlugins := ParseGlobalPluginYaml("./plugin/globalDetectors")
	gloabalDetections := decisionmakerPB.GlobalDetections{}
	RunGlobalDetectors(&gloabalDetections, globalPlugins, ctx, path)
	decisionMakerInput.GloabalDetections = &gloabalDetections
	var mxLang = 0.0
	for _, lang := range languages {
		mxLang = math.Max(mxLang, lang.Percent)
	}
	for _, language := range languages {
		if language.Percent == mxLang {
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
					runtimeVersion := runners.DetectRuntime(ctx, path, pluginDetails)
					runners.RunPreDetectCommand(ctx, &pb.ServiceInput{
						RuntimeVersion: runtimeVersion,
						Root:           path,
					}, pluginDetails)
					if runtimeVersion != "" {
						allDependencies := runners.GetParsedDependencies(ctx, runtimeVersion, path, pluginDetails)
						languageSpecificDetections := decisionmakerPB.LanguageSpecificDetections{
							Name:           language.Name,
							RuntimeVersion: runtimeVersion,
						}
						commands := decisionmakerPB.Commands{}
						RunAllDetectors(ctx, &languageSpecificDetections, allDependencies, pluginDetails, runtimeVersion, path, globalPlugins, &commands)
						decisionMakerInput.LanguageSpecificDetection = append(decisionMakerInput.LanguageSpecificDetection, &languageSpecificDetections)
						decisionMakerInput.Commands = &commands
					}

				}
			}()
		}
	}
	wg.Wait()
	return decisionMakerInput
}

func RunGlobalDetectors(globalDetection *decisionmakerPB.GlobalDetections, globalPlugin *versionsPB.GlobalPlugin, ctx context.Context, path string) {
	var wg sync.WaitGroup
	wg.Add(3)
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
		globalDetection.ProcFile = runners.DetectProcFile(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		globalDetection.Makefile = runners.DetectMakeFile(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
}

//RunAllDetectors it runs all detectors of dependencies ex. orm,framework etc ....
func RunAllDetectors(ctx context.Context, languageSpecificDetections *decisionmakerPB.LanguageSpecificDetections, allDependencies map[string]map[string]runners.DependencyDetail, pluginDetails *versionsPB.LanguageVersion, runtimeVersion string, path string, globalPlugin *versionsPB.GlobalPlugin, commands *decisionmakerPB.Commands) {
	var wg sync.WaitGroup
	wg.Add(9)
	var mutex = &sync.Mutex{}
	go func() {
		defer wg.Done()
		if pluginDetails.Commands != "" {
			mutex.Lock()
			languageSpecificDetections.Commands = runners.GetCommands(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
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
		commands = runners.DetectAndRunCommands(ctx, path, globalPlugin)
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
