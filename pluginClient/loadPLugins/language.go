package loadPLugins

import (
	"code-analyser/protos/pb"
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
	Setting           *utils.Setting
}

func (languagePlugins *LanguagePlugin) Load(ctx context.Context, pluginYamlFiles []utils.FileDetails) error {
    languagePlugins.Setting.Logger.Debug("language plugin loading started")

	languagePlugins.DetectRunTime = &DetectRunTimePlugin{Setting: languagePlugins.Setting}
	languagePlugins.Commands = &CommandsPlugin{Setting: languagePlugins.Setting}
	languagePlugins.Dependencies = &GetDependenciesPlugin{Setting: languagePlugins.Setting}
	languagePlugins.Framework = &FrameworkPlugin{Setting: languagePlugins.Setting}
	languagePlugins.Env = &EnvPlugin{Setting: languagePlugins.Setting}
	languagePlugins.PreDetectCommands = &PreDetectCommandsPlugin{Setting: languagePlugins.Setting}
	languagePlugins.StaticAssets = &StaticAssetsPlugin{Setting: languagePlugins.Setting}
	languagePlugins.TestCommand = &TestCommandPlugin{Setting: languagePlugins.Setting}
	languagePlugins.BuildDirectory = &BuildDirectoryPlugin{Setting: languagePlugins.Setting}
	languagePlugins.Db = &DbPlugin{Setting: languagePlugins.Setting}
	languagePlugins.Library = &LibraryPlugin{Setting: languagePlugins.Setting}
	languagePlugins.Orm = &OrmPlugin{Setting: languagePlugins.Setting}

    languagePlugins.Setting.Logger.Info("language plugin yaml file reading and client initialization started")
	for _, pluginFile := range pluginYamlFiles {
		parsedRawFile, err := utils.ReadPluginYamlFile(ctx, pluginFile)
		if err != nil {
			log.Println(err)
			return err
		}

		pluginYamlFile := parsedRawFile.PluginDetails

		switch pluginYamlFile.Type {
		case "detectRuntime":
			languagePlugins.DetectRunTime.Load(pluginYamlFile)
		case "commands":
			languagePlugins.Commands.Load(pluginYamlFile)
		case "env":
			languagePlugins.Env.Load(pluginYamlFile)
		case "preDetectCommand":
			languagePlugins.PreDetectCommands.Load(pluginYamlFile)
		case "staticAssets":
			languagePlugins.StaticAssets.Load(pluginYamlFile)
		case "buildDirectory":
			languagePlugins.BuildDirectory.Load(pluginYamlFile)
		case "testCasesCommands":
			languagePlugins.TestCommand.Load(pluginYamlFile)
		case "framework":
			languagePlugins.Framework.Load(pluginYamlFile)
		case "orm":
			languagePlugins.Orm.Load(pluginYamlFile)
		case "library":
			languagePlugins.Library.Load(pluginYamlFile)
		case "database":
			languagePlugins.Db.Load(pluginYamlFile)
		case "getDependencies":
			languagePlugins.Dependencies.Load(pluginYamlFile)
		}
	}
	return nil
}

func (languagePlugins *LanguagePlugin) Run(ctx context.Context, projectRootPath string) (*pb.LanguageSpecificDetections, error) {
	detectedDependencies := &pb.LanguageSpecificDetections{}
	var projectDependencies map[string]string

	res, err := languagePlugins.DetectRunTime.Run(ctx, projectRootPath)
	if err != nil {
		return detectedDependencies, err
	}
	runtimeVersion := res.Value
	detectedDependencies.RuntimeVersion = runtimeVersion

	if languagePlugins.PreDetectCommands.Client != nil {
		err = languagePlugins.PreDetectCommands.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
	}

	if languagePlugins.Dependencies.Dependencies != nil {
		response, err := languagePlugins.Dependencies.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		projectDependencies = response.Value
	}

	if languagePlugins.Env != nil {
		env, err := languagePlugins.Env.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		detectedDependencies.Env = env
	}

	if languagePlugins.TestCommand != nil {
		testCommand, err := languagePlugins.TestCommand.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		detectedDependencies.TestCases = testCommand
	}

	if languagePlugins.Orm.Orms != nil && projectDependencies != nil {
		orms := languagePlugins.Orm.Extract(ctx, projectDependencies)
		ormOutput, err := languagePlugins.Orm.Run(ctx, orms, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		detectedDependencies.Orm = ormOutput
	}

	if languagePlugins.Db.Dbs != nil && projectDependencies != nil {
		databases := languagePlugins.Db.Extract(ctx, projectDependencies)
		dbOutput, err := languagePlugins.Db.Run(ctx, databases, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		detectedDependencies.Db = dbOutput
	}

	if languagePlugins.Library.Libraries != nil && projectDependencies != nil {
		libraries := languagePlugins.Library.Extract(ctx, projectDependencies)
		librariesOutput, err := languagePlugins.Library.Run(ctx, libraries, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		detectedDependencies.Libraries = librariesOutput
	}

	if languagePlugins.Framework.Frameworks != nil && projectDependencies != nil {
		frameworks := languagePlugins.Framework.Extract(ctx, projectDependencies)
		frameworksOutput, err := languagePlugins.Framework.Run(ctx, frameworks, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		detectedDependencies.Framework = frameworksOutput
	}

	if languagePlugins.Commands.Client != nil {
		commands, err := languagePlugins.Commands.Run(ctx, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		detectedDependencies.Commands = commands
	}

	if languagePlugins.BuildDirectory.Client != nil {
		buildDirectoryOutput, err := languagePlugins.BuildDirectory.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		detectedDependencies.BuildDirectory = buildDirectoryOutput
	}

	if languagePlugins.StaticAssets.Client != nil {
		staticAssets, err := languagePlugins.StaticAssets.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		detectedDependencies.StaticAssets = staticAssets
	}

	return detectedDependencies, nil
}
