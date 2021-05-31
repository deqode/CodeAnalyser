package main

import (
	"code-analyser/analyser"
	decisionmakerPB "code-analyser/protos/pb"
	pb "code-analyser/protos/pb/plugin"
	"code-analyser/runners"
	"golang.org/x/net/context"
	"log"
	"math"
)

func main() {
	//path := os.Args[1]
	//log.Println("Initialized Scrapping ")
	//log.Println("Scrapping on "+path)
	decisionMakerInput := Scrape("/home/deqode/Desktop/basicRepo")
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
	languages, _, _ := analyser.GetLanguagesWithPercent(path)
	supportedLanguages, _ := SupportedLanguagesParser()
	decisionMakerInput := &decisionmakerPB.DecisionMakerInput{
		LanguageSpecificDetection: []*decisionmakerPB.LanguageSpecificDetections{},
		GloabalDetections:         &decisionmakerPB.GlobalDetections{},
	}
	var ctx context.Context = nil

	globalPlugins := LoadGlobalFilesPluginInfo("./plugin/globalDetectors")
	globalDetections := decisionmakerPB.GlobalDetections{}
	RunAllGlobalPlugins(&globalDetections, globalPlugins, ctx, path)
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

			pluginDetails := LoadLanguageSpecificPluginInfo(languagePath)
			runtimeVersion := runners.DetectRuntime(ctx, path, pluginDetails)
			//TODO: change name after dicuss
			runners.RunPreDetectCommand(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)

			//if no app runtime found =>skip
			if runtimeVersion == "" {
				continue
			}
			allDependencies := runners.GetParsedDependencies(ctx, runtimeVersion, path, pluginDetails)
			languageSpecificDetections := decisionmakerPB.LanguageSpecificDetections{
				Name:           language.Name,
				RuntimeVersion: runtimeVersion,
			}
			commands := decisionmakerPB.Commands{}
			RunAllLanguageSpecificPlugins(ctx, &languageSpecificDetections, allDependencies, pluginDetails, runtimeVersion, path, globalPlugins, &commands)
			decisionMakerInput.LanguageSpecificDetection = append(decisionMakerInput.LanguageSpecificDetection, &languageSpecificDetections)
			decisionMakerInput.Commands = &commands
		}

	}
	return decisionMakerInput
}
