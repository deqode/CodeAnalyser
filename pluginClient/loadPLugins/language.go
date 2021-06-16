package loadPLugins

import (
	"code-analyser/utils"
	"golang.org/x/net/context"
	"log"
)

type LanguagePlugin struct {
	Framework         *FrameworkPlugin
	Db                *DbPlugin
	Library           *LibraryPlugin
	BuildDirectory    *BuildDirectoryPlugin
	Commands          *CommandsPlugin
	DockerFile        *DockerFilePlugin
	Env               *EnvPlugin
	Orm               *OrmPlugin
	DetectRunTime     *DetectRunTimePlugin
	PreDetectCommands *PreDetectCommandsPlugin
	StaticAssets      *StaticAssetsPlugin
	Dependencies      *GetDependenciesPlugin
	TestCommand       *TestCommandPlugin
}

func (languagePlugins *LanguagePlugin) Load(ctx context.Context, pluginYamlFiles []utils.FileDetails) error {

	for _, pluginFile := range pluginYamlFiles {
		parsedRawFile, err := utils.ReadPluginYamlFile(ctx, pluginFile)
		if err != nil {
			log.Println(err)
			return err
		}

		pluginYamlFile := parsedRawFile.PluginDetails

		switch pluginYamlFile.Type {
		case "detectRuntime":
			//languagePlugins.DetectRunTime.Load(pluginYamlFile)
		case "commands":
			languagePlugins.Commands = &CommandsPlugin{}
			languagePlugins.Commands.Load(pluginYamlFile)
		case "env":
			//languagePlugins.Env.Load(pluginYamlFile)
		case "preDetectCommand":
			//languagePlugins.PreDetectCommands.Load(pluginYamlFile)
		case "staticAssets":
			//languagePlugins.StaticAssets.Load(pluginYamlFile)
		case "buildDirectory":
			//languagePlugins.BuildDirectory.Load(pluginYamlFile)
		case "testCasesCommands":
			//languagePlugins.TestCommand.Load(pluginYamlFile)
		case "framework":
			//languagePlugins.Framework.Load(pluginYamlFile)
		case "orm":
			//languagePlugins.Orm.Load(pluginYamlFile)
		case "library":
			//languagePlugins.Library.Load(pluginYamlFile)
		case "database":
			//languagePlugins.Db.Load(pluginYamlFile)
			//case "getDependencies":
			//	languagePlugins.RuntimeVersions[pluginYamlFile.Version] = &pluginDetailsPB.DependencyVersionDetails{
			//		Semver:     pluginYamlFile.Semver,
			//		PluginPath: pluginYamlFile.Command,
			//	}
		}
	}
	return nil

}
