package main

import (
	"code-analyser/analyser"
	"code-analyser/pluginClient/loadPLugins"
	decisionmakerPB "code-analyser/protos/pb"
	"code-analyser/protos/pb/pluginDetails"
	"code-analyser/utils"
	"golang.org/x/net/context"
	"log"
)

func main() {
	//path := os.Args[1]
	//log.Println("Initialized Scrapping ")
	//log.Println("Scrapping on "+path)
	var ctx context.Context
	var path = "/home/deqode/Downloads/compare-vue-master"
	log.Println(Scrape(ctx, path))
}

func LoadGlobalPlugin(ctx context.Context, globalPluginPath string) *loadPLugins.GlobalPlugin {
	PluginYamlFiles, err := utils.SearchFileInDirectory("pluginDetails.yaml", globalPluginPath)
	if err != nil {
		log.Println("not able to get yaml file of plugin")
	}

	plugins := loadPLugins.GlobalPlugin{}
	err = plugins.Load(ctx, PluginYamlFiles)
	if err != nil {
		log.Println("not able start global plugins")
	}
	return &plugins
}

func LoadLanguagePlugin(ctx context.Context, languagePluginPath []*pluginDetails.Language) map[string]*loadPLugins.LanguagePlugin {
	allLanguagePlugins := map[string]*loadPLugins.LanguagePlugin{}
	for _, language := range languagePluginPath {
		name := language.Name
		path := language.Path

		pluginYamlFiles, err := utils.SearchFileInDirectory("pluginDetails.yaml", path)
		if err != nil {
			log.Println("not able to get yaml file of " + name + " plugin")
		}

		plugins := &loadPLugins.LanguagePlugin{}
		err = plugins.Load(ctx, pluginYamlFiles)
		if err != nil {
			log.Println("not able start plugins of " + name)
		}
		allLanguagePlugins[name] = plugins
	}
	return allLanguagePlugins
}

//Scrape it scrape language, framework, orm etc .....
func Scrape(ctx context.Context, path string) (*decisionmakerPB.DecisionMakerInput, error) {
	languages, _, _ := analyser.GetLanguagesWithPercent(path)
	supportedLanguages, _ := SupportedLanguagesParser()

	decisionMakerInput := &decisionmakerPB.DecisionMakerInput{
		LanguageSpecificDetections: []*decisionmakerPB.LanguageSpecificDetections{},
		GloabalDetections:          &decisionmakerPB.GlobalDetections{},
	}

	globalPLugins := LoadGlobalPlugin(ctx, "./plugin/globalDetectors")
	languagePlugins := LoadLanguagePlugin(ctx, supportedLanguages.Languages)

	globalDetections, err := globalPLugins.Run(ctx, path)
	if err != nil {
		return decisionMakerInput, err
	}
	decisionMakerInput.GloabalDetections = globalDetections

	for _, language := range languages {
		if languagePlugin, ok := languagePlugins[language.Name]; ok {
			detections, err := languagePlugin.Run(ctx, path)
			if err != nil {
				return decisionMakerInput, err
			}
			decisionMakerInput.LanguageSpecificDetections = append(decisionMakerInput.LanguageSpecificDetections, detections)
		}
	}
	return decisionMakerInput, nil
}
