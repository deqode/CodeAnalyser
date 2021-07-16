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

//Load It calls load function on all language plugins
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
			languagePlugins.Setting.Logger.Info("detectRuntime plugin client creation started")
			languagePlugins.DetectRunTime.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("detectRuntime plugin client created successfully")
		case "commands":
			languagePlugins.Setting.Logger.Info("commands plugin client creation started")
			languagePlugins.Commands.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("commands plugin client created successfully")
		case "env":
			languagePlugins.Setting.Logger.Info("env plugin client creation started")
			languagePlugins.Env.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("env plugin client created successfully")
		case "preDetectCommand":
			languagePlugins.Setting.Logger.Info("preDetectCommand plugin client creation started")
			languagePlugins.PreDetectCommands.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("preDetectCommand plugin client created successfully")
		case "staticAssets":
			languagePlugins.Setting.Logger.Info("staticAssets plugin client creation started")
			languagePlugins.StaticAssets.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("staticAssets plugin client created successfully")
		case "buildDirectory":
			languagePlugins.Setting.Logger.Info("buildDirectory plugin client creation started")
			languagePlugins.BuildDirectory.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("buildDirectory plugin client created successfully")
		case "testCasesCommands":
			languagePlugins.Setting.Logger.Info("testCasesCommands plugin client creation started")
			languagePlugins.TestCommand.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("testCasesCommands plugin client created successfully")
		case "framework":
			languagePlugins.Setting.Logger.Info("framework plugin client creation started")
			languagePlugins.Framework.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("framework plugin client created successfully")
		case "orm":
			languagePlugins.Setting.Logger.Info("orm plugin client creation started")
			languagePlugins.Orm.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("orm plugin client created successfully")
		case "library":
			languagePlugins.Setting.Logger.Info("library plugin client creation started")
			languagePlugins.Library.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("library plugin client created successfully")
		case "database":
			languagePlugins.Setting.Logger.Info("database plugin client creation started")
			languagePlugins.Db.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("database plugin client created successfully")
		case "getDependencies":
			languagePlugins.Setting.Logger.Info("getDependencies plugin client creation started")
			languagePlugins.Dependencies.Load(pluginYamlFile)
			languagePlugins.Setting.Logger.Info("getDependencies plugin client created successfully")
		}
	}
	languagePlugins.Setting.Logger.Info("language plugin yaml file reading and client initialization completed")

	languagePlugins.Setting.Logger.Debug("language plugin loading completed")
	return nil
}

//Run It calls run function on all language plugins
func (languagePlugins *LanguagePlugin) Run(ctx context.Context, projectRootPath string) (*pb.LanguageSpecificDetections, error) {
	languagePlugins.Setting.Logger.Debug("language plugins execution started")
	detectedDependencies := &pb.LanguageSpecificDetections{}
	var projectDependencies map[string]string

	languagePlugins.Setting.Logger.Info("detectRuntime plugin methods execution started")
	res, err := languagePlugins.DetectRunTime.Run(ctx, projectRootPath)
	if err != nil {
		return detectedDependencies, err
	}
	languagePlugins.Setting.Logger.Info("detectRuntime plugin methods execution completed")

	runtimeVersion := res.Value
	detectedDependencies.RuntimeVersion = runtimeVersion

	if languagePlugins.PreDetectCommands.Client != nil {
		languagePlugins.Setting.Logger.Info("preDetectCommands plugin methods execution started")
		err = languagePlugins.PreDetectCommands.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		languagePlugins.Setting.Logger.Info("preDetectCommands plugin methods execution completed")
	}

	if languagePlugins.Dependencies.Dependencies != nil {
		languagePlugins.Setting.Logger.Info("detectDependencies plugin methods execution started")
		response, err := languagePlugins.Dependencies.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		languagePlugins.Setting.Logger.Info("detectDependencies plugin methods execution completed")
		projectDependencies = response.Value
	}

	if languagePlugins.Env != nil {
		languagePlugins.Setting.Logger.Info("env plugin methods execution started")
		env, err := languagePlugins.Env.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		languagePlugins.Setting.Logger.Info("env plugin methods execution completed")
		detectedDependencies.Env = env
	}

	if languagePlugins.TestCommand != nil {
		languagePlugins.Setting.Logger.Info("testCommand plugin methods execution started")
		testCommand, err := languagePlugins.TestCommand.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		languagePlugins.Setting.Logger.Info("testCommand plugin methods execution completed")
		detectedDependencies.TestCases = testCommand
	}

	if languagePlugins.Orm.Orms != nil && projectDependencies != nil {
		languagePlugins.Setting.Logger.Info("filtration process of orm's plugin supported by us started")
		orms := languagePlugins.Orm.Extract(ctx, projectDependencies)
		languagePlugins.Setting.Logger.Info("filtration process of orm's plugin completed")

		languagePlugins.Setting.Logger.Info("orm plugin methods execution started")
		ormOutput, err := languagePlugins.Orm.Run(ctx, orms, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		languagePlugins.Setting.Logger.Info("orm plugin methods execution completed")
		detectedDependencies.Orm = ormOutput
	}

	if languagePlugins.Db.Dbs != nil && projectDependencies != nil {
		languagePlugins.Setting.Logger.Info("filtration process of db's plugin supported by us started")
		databases := languagePlugins.Db.Extract(ctx, projectDependencies)
		languagePlugins.Setting.Logger.Info("filtration process of db's plugin completed")

		languagePlugins.Setting.Logger.Info("db plugin methods execution started")
		dbOutput, err := languagePlugins.Db.Run(ctx, databases, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		languagePlugins.Setting.Logger.Info("db plugin methods execution completed")
		detectedDependencies.Db = dbOutput
	}

	if languagePlugins.Library.Libraries != nil && projectDependencies != nil {
		languagePlugins.Setting.Logger.Info("filtration process of library's plugin supported by us started")
		libraries := languagePlugins.Library.Extract(ctx, projectDependencies)
		languagePlugins.Setting.Logger.Info("filtration process of library's plugin completed")

		languagePlugins.Setting.Logger.Info("library plugin methods execution started")
		librariesOutput, err := languagePlugins.Library.Run(ctx, libraries, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		languagePlugins.Setting.Logger.Info("library plugin methods execution completed")
		detectedDependencies.Libraries = librariesOutput
	}

	if languagePlugins.Framework.Frameworks != nil && projectDependencies != nil {
		languagePlugins.Setting.Logger.Info("filtration process of framework's plugin supported by us started")
		frameworks := languagePlugins.Framework.Extract(ctx, projectDependencies)
		languagePlugins.Setting.Logger.Info("filtration process of framework's plugin completed")

		languagePlugins.Setting.Logger.Info("framework plugin methods execution started")
		frameworksOutput, err := languagePlugins.Framework.Run(ctx, frameworks, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		languagePlugins.Setting.Logger.Info("framework plugin methods execution completed")
		detectedDependencies.Framework = frameworksOutput
	}

	if languagePlugins.Commands.Client != nil {
		languagePlugins.Setting.Logger.Info("commands plugin methods execution started")
		commands, err := languagePlugins.Commands.Run(ctx, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		languagePlugins.Setting.Logger.Info("commands plugin methods execution completed")
		detectedDependencies.Commands = commands
	}

	if languagePlugins.BuildDirectory.Client != nil {
		languagePlugins.Setting.Logger.Info("buildDirectory plugin methods execution started")
		buildDirectoryOutput, err := languagePlugins.BuildDirectory.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		languagePlugins.Setting.Logger.Info("buildDirectory plugin methods execution completed")
		detectedDependencies.BuildDirectory = buildDirectoryOutput
	}

	if languagePlugins.StaticAssets.Client != nil {
		languagePlugins.Setting.Logger.Info("staticAssets plugin methods execution started")
		staticAssets, err := languagePlugins.StaticAssets.Run(ctx, runtimeVersion, projectRootPath)
		if err != nil {
			return detectedDependencies, err
		}
		languagePlugins.Setting.Logger.Info("staticAssets plugin methods execution completed")
		detectedDependencies.StaticAssets = staticAssets
	}

	languagePlugins.Setting.Logger.Debug("language plugins execution completed")
	return detectedDependencies, nil
}
