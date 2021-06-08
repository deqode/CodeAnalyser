package main

import (
	"code-analyser/analyser"
	"code-analyser/pluginClient"
	decisionmakerPB "code-analyser/protos/pb"
	pb "code-analyser/protos/pb/plugin"
	"code-analyser/runners"
	"code-analyser/utils"
	"golang.org/x/net/context"
	"log"
	"math"
)

func main() {
	//path := os.Args[1]
	//log.Println("Initialized Scrapping ")
	//log.Println("Scrapping on "+path)
	//decisionMakerInput := Scrape("/home/deqode/Desktop/project/go/code-analyser/testingRepos/detectDockerFile/repo1")
	//log.Println(decisionMakerInput.LanguageSpecificDetection)
	////res, err := json.MarshalIndent(decisionMakerInput.LanguageSpecificDetection[0], "", "  ")
	////if err != nil {
	////	log.Println("error:", err)
	////}
	////fmt.Print(string(res))
	//log.Println(decisionMakerInput.GloabalDetections)

	dbResponse, client := pluginClient.CreateLibraryClient(utils.CallPluginCommand("python3 plugin/python/libraries/numpy/v1_x/server.py"))
	for client.Exited() {
		client.Kill()
	}
	log.Println(dbResponse.PercentOfUsed(&pb.Input{
		RuntimeVersion: "",
		RootPath:       "",
	}))
}

//Scrape it scrape language, framework, orm etc .....
func Scrape(ctx context.Context, path string) *decisionmakerPB.DecisionMakerInput {
	languages, _, _ := analyser.GetLanguagesWithPercent(path)
	supportedLanguages, _ := SupportedLanguagesParser()

	decisionMakerInput := &decisionmakerPB.DecisionMakerInput{
		LanguageSpecificDetection: []*decisionmakerPB.LanguageSpecificDetections{},
		GloabalDetections:         &decisionmakerPB.GlobalDetections{},
	}

	globalPlugins := GetGlobalPluginsPath(ctx, "./plugin/globalDetectors")
	globalDetections := decisionmakerPB.GlobalDetections{}
	LoadGlobalPlugins(ctx, &globalDetections, globalPlugins, path)
	decisionMakerInput.GloabalDetections = &globalDetections

	var mxLang = 0.0
	//TODO:formula to get main language used priority wise
	for _, lang := range languages {
		mxLang = math.Max(mxLang, lang.Percent)
	}
	for _, language := range languages {
		if language.Percent == mxLang {
			var languagePath string
			for _, supportedLanguage := range supportedLanguages.Languages {
				if supportedLanguage.Name == language.Name {
					languagePath = supportedLanguage.Path
					break
				}
			}

			//if no supported language found=>skip
			if languagePath == "" {
				continue
			}

			pluginDetails := GetLanguagePluginspath(ctx,languagePath)
			runtimeVersion := runners.ExecuteRuntimeDetectionPlugin(ctx, path, pluginDetails.DetectRuntimePluginPath)
			//TODO: change name after dicuss
			runners.ExecutePreDetectionPlugin(ctx, &pb.Input{
				RuntimeVersion: runtimeVersion,
				RootPath:       path,
			}, pluginDetails.PreDetectCommandPluginPath)

			//if no app runtime found =>skip
			if runtimeVersion == "" {
				continue
			}
			allDependencies := runners.GetDependenciesFromProject(ctx, runtimeVersion, path, pluginDetails)
			languageSpecificDetections := decisionmakerPB.LanguageSpecificDetections{
				Name:           language.Name,
				RuntimeVersion: runtimeVersion,
			}
			LoadLanguagePlugins(ctx, &languageSpecificDetections, *allDependencies, pluginDetails, runtimeVersion, path)
			decisionMakerInput.LanguageSpecificDetection = append(decisionMakerInput.LanguageSpecificDetection, &languageSpecificDetections)
		}

	}
	return decisionMakerInput
}
