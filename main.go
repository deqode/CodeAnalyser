package main

import (
	detectlanguage "code-analyser/detectLanguage"
	"code-analyser/pluginClient/loadPLugins"
	decisionmakerPB "code-analyser/protos/pb"
	"code-analyser/protos/pb/pluginDetails"
	"code-analyser/utils"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"log"
)

type Analyser struct {
	Setting *utils.Setting
}

func main() {
	var ctx context.Context
	var path = "./"
	logger, _ := zap.NewProduction()
	set := utils.Setting{
		Logger: logger,
	}
	an := Analyser{Setting: &set}
	an.Setting.Logger.Info("scrapping started")
	log.Println(an.Scrape(ctx, path))
	an.Setting.Logger.Info("scrapping completed")
}

//Scrape it scrape language, framework, orm etc .....
func (analyser *Analyser) Scrape(ctx context.Context, path string) (*decisionmakerPB.DecisionMakerInput, error) {
	analyser.Setting.Logger.Debug("scrapping started")

	languages, _, _ := detectlanguage.GetLanguagesWithPercent(path)
	supportedLanguages, _ := SupportedLanguagesParser()

	decisionMakerInput := &decisionmakerPB.DecisionMakerInput{
		LanguageSpecificDetections: []*decisionmakerPB.LanguageSpecificDetections{},
		GloabalDetections:          &decisionmakerPB.GlobalDetections{},
	}

	analyser.Setting.Logger.Info("global plugin loading started")
	globalPlugins := analyser.LoadGlobalPlugin(ctx, "./plugin/globalDetectors")
	analyser.Setting.Logger.Info("global plugin loaded successfully")

	analyser.Setting.Logger.Info("language plugins loading started")
	languagePlugins := analyser.LoadLanguagePlugin(ctx, supportedLanguages.Languages)
	analyser.Setting.Logger.Info("language plugins loaded")

	analyser.Setting.Logger.Info("global plugins execution started")
	globalDetections, err := globalPlugins.Run(ctx, path)
	if err != nil {
		return decisionMakerInput, err
	}
	analyser.Setting.Logger.Info("global plugins execution completed")
	decisionMakerInput.GloabalDetections = globalDetections

	analyser.Setting.Logger.Info("language plugin execution started")
	for _, language := range languages {
		if languagePlugin, ok := languagePlugins[language.Name]; ok {
			analyser.Setting.Logger.Info(language.Name + " plugins execution started")
			detections, err := languagePlugin.Run(ctx, path)
			if err != nil {
				return decisionMakerInput, err
			}
			decisionMakerInput.LanguageSpecificDetections = append(decisionMakerInput.LanguageSpecificDetections, detections)
			analyser.Setting.Logger.Info(language.Name + " plugins execution completed")
		}
	}
	analyser.Setting.Logger.Info("language plugin execution completed")

	analyser.Setting.Logger.Debug("scrapped completed")
	return decisionMakerInput, nil
}

func (analyser *Analyser) LoadGlobalPlugin(ctx context.Context, globalPluginPath string) *loadPLugins.GlobalPlugin {
	analyser.Setting.Logger.Debug("global plugin loading started")

	analyser.Setting.Logger.Info("global plugin's yaml file path reading completed")
	PluginYamlFiles, err := utils.SearchFileInDirectory("pluginDetails.yaml", globalPluginPath)
	if err != nil {
		analyser.Setting.Logger.Error("not able to get path of plugin yaml file, " + err.Error())
	}
	analyser.Setting.Logger.Info("global plugin's yaml file path reading completed")

	plugins := loadPLugins.GlobalPlugin{
		Setting: analyser.Setting,
	}

	analyser.Setting.Logger.Info("global plugin client creation started")
	err = plugins.Load(ctx, PluginYamlFiles)
	if err != nil {
		analyser.Setting.Logger.Error("not able to start global plugins, " + err.Error())
	}
	analyser.Setting.Logger.Info("global plugin client creation completed successfully")

	analyser.Setting.Logger.Debug("global plugin loaded successfully")
	return &plugins
}

func (analyser *Analyser) LoadLanguagePlugin(ctx context.Context, languagePluginPath []*pluginDetails.Language) map[string]*loadPLugins.LanguagePlugin {
	analyser.Setting.Logger.Debug("language plugin loading started")

	allLanguagePlugins := map[string]*loadPLugins.LanguagePlugin{}
	for _, language := range languagePluginPath {
		name := language.Name
		path := language.Path
		analyser.Setting.Logger.Debug(name + "'s plugin loading started")

		analyser.Setting.Logger.Info(name + " plugin's yaml file path reading started")
		pluginYamlFiles, err := utils.SearchFileInDirectory("pluginDetails.yaml", path)
		if err != nil {
			analyser.Setting.Logger.Error("not able to get yaml file of " + name + "'s plugin")
		}
		analyser.Setting.Logger.Info(name + " plugin's yaml file path reading completed")

		plugins := &loadPLugins.LanguagePlugin{}

		analyser.Setting.Logger.Info(name + "'s plugin client creation started")
		err = plugins.Load(ctx, pluginYamlFiles)
		if err != nil {
			analyser.Setting.Logger.Error("not able start plugins of " + name + err.Error())
		}
		analyser.Setting.Logger.Info(name + "'s plugin client creation completed successfully")

		analyser.Setting.Logger.Debug(name + "'s plugin loaded successfully")
		allLanguagePlugins[name] = plugins
	}

	analyser.Setting.Logger.Debug("language plugin loaded successfully")
	return allLanguagePlugins
}
