package main

import (
	"code-analyser/analyser"
	"code-analyser/pluginClient/loadPLugins"
	decisionmakerPB "code-analyser/protos/pb"
	"code-analyser/utils"
	"golang.org/x/net/context"
	"log"
	"math"
)

func main() {
	//path := os.Args[1]
	//log.Println("Initialized Scrapping ")
	//log.Println("Scrapping on "+path)
	var ctx context.Context

	globalPluginYamlsPath, err := utils.SearchFileInDirectory("pluginDetails.yaml", "./plugin/globalDetectors")
	if err != nil {
		log.Println("not able to get yaml file of plugin")
	}

	var globalPlugins loadPLugins.GlobalPlugin

	log.Println(globalPlugins)
	err = globalPlugins.Load(ctx, globalPluginYamlsPath)
	if err != nil {
		log.Println("not able start global plugins")
	}
	log.Println(globalPlugins)
	dockerfile, _, err := globalPlugins.DockerFile.Run(nil, "./abc")
	log.Println(dockerfile, err)
}

//Scrape it scrape language, framework, orm etc .....
func Scrape(ctx context.Context, path string) *decisionmakerPB.DecisionMakerInput {
	languages, _, _ := analyser.GetLanguagesWithPercent(path)
	supportedLanguages, _ := SupportedLanguagesParser()

	decisionMakerInput := &decisionmakerPB.DecisionMakerInput{
		LanguageSpecificDetections: []*decisionmakerPB.LanguageSpecificDetections{},
		GloabalDetections:          &decisionmakerPB.GlobalDetections{},
	}

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

			//pluginDetails := GetLanguagePluginsPath(ctx, languagePath)
			//runtimeVersion := runners.ExecuteRuntimeDetectionPlugin(ctx, path, pluginDetails.DetectRuntimePluginPath)
			////TODO: change name after dicuss
			//runners.ExecutePreDetectionPlugin(ctx, &pb.Input{
			//	RuntimeVersion: runtimeVersion,
			//	RootPath:       path,
			//}, pluginDetails.PreDetectCommandPluginPath)
			//
			////if no app runtime found =>skip
			//if runtimeVersion == "" {
			//	continue
			//}
			//allDependency := runners.GetDependenciesFromProject(ctx, runtimeVersion, path, pluginDetails)
			//languageSpecificDetections := decisionmakerPB.LanguageSpecificDetections{
			//	Name:           language.Name,
			//	RuntimeVersion: runtimeVersion,
			//}
			//LoadLanguagePlugins(ctx, &languageSpecificDetections, *allDependency, pluginDetails, runtimeVersion, path)
			//decisionMakerInput.LanguageSpecificDetections = append(decisionMakerInput.LanguageSpecificDetections, &languageSpecificDetections)
		}

	}
	return decisionMakerInput
}
