package runners

import (
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
	"code-analyser/protos/pb/pluginDetails"
	commonUtils "code-analyser/utils"
	"golang.org/x/net/context"
)

type DependencyCase struct {
	Input  DependencyInput
	Output map[string]map[string]DependencyDetail
}

type DependencyInput struct {
	Ctx             context.Context
	LanguageVersion string
	Path            string
	PluginDetails   *pluginDetails.LanguagePlugins
}

var DependencyCases = []DependencyCase{
	{
		Input: DependencyInput{
			Ctx:             nil,
			LanguageVersion: "10.19.0",
			Path:            commonUtils.ProjectPath() + "/testingRepos/parsedDependencies/testRepo1",
			PluginDetails:   &SupportedDependencies,
		},
		Output: map[string]map[string]DependencyDetail{
			Framework: {
				"express": {
					Version: "v1.x",
					Command: "node plugin/js/framework/express/v1_x/server.js",
				},
			},
			DB: {
				"postgres": {
					Version: "v1.x",
					Command: "node plugin/js/db/postgres/v1_x/server.js",
				},
				"redis": {
					Version: "v1.x",
					Command: "node plugin/js/db/redis/v1_x/server.js",
				},
			},
			Library: {},
			ORM:     {},
		},
	},
	{
		Input: DependencyInput{
			Ctx:             nil,
			LanguageVersion: "10.19.0",
			Path:            commonUtils.ProjectPath() + "/testingRepos/parsedDependencies/testRepo2",
			PluginDetails:   &SupportedDependencies,
		},
		Output: map[string]map[string]DependencyDetail{
			Framework: {
				"nestJs": {
					Version: "v1.x",
					Command: "node plugin/js/framework/nest/v1_x/server.js",
				},
			},
			DB: {
				"MariaDB": {
					Version: "v1.x",
					Command: "node plugin/js/db/mariadb/v1_x/server.js",
				},
			},
			Library: {
				"react": {
					Version: "v1.x",
					Command: "node plugin/js/libraries/react/v1_x/server.js"},
			},
			ORM: {
				"typeorm": {
					Version: "v1.x",
					Command: "node plugin/js/orm/typeorm/v1_x/server.js"}},
		},
	},
	{
		Input: DependencyInput{
			Ctx:             nil,
			LanguageVersion: "10.19.0",
			Path:            commonUtils.ProjectPath() + "/testingRepos/emptyRepo",
			PluginDetails:   &SupportedDependencies,
		},
		Output: map[string]map[string]DependencyDetail{
			Framework: {},
			DB:        {},
			Library:   {},
			ORM:       {},
		},
	},
}

type DetectRunTimeCase struct {
	Input  DetectRunTimeInput
	Output string
}

type DetectRunTimeInput struct {
	Ctx            context.Context
	Path           string
	YamlLangObject *pluginDetails.LanguagePlugins
}

var DetectRunTimeCases = []DetectRunTimeCase{
	{
		Input: DetectRunTimeInput{
			Ctx:            nil,
			Path:           commonUtils.ProjectPath() + "/testingRepos/detectRunTime/repo1",
			YamlLangObject: &SupportedDependencies,
		},
		Output: "12.3",
	},
}

type DetectTestCommandCase struct {
	Input  DetectTestCommandInput
	Output *languageSpecific.TestCasesCommand
}

type DetectTestCommandInput struct {
	Ctx           context.Context
	Input         *helpers.Input
	PluginDetails *pluginDetails.LanguagePlugins
}

var DetectTestCommandCases = []DetectTestCommandCase{
	{
		Input: DetectTestCommandInput{
			Ctx: nil,
			Input: &helpers.Input{
				RuntimeVersion: "",
				RootPath:       commonUtils.ProjectPath() + "/testingRepos/detectTestCase/repo1",
			},
			PluginDetails: &SupportedDependencies,
		},
		Output: &languageSpecific.TestCasesCommand{
			Used: true,
			Commands: []*helpers.Command{
				{
					Command: "npm",
					Args:    []string{"test:map"},
				},
				{
					Command: "npm",
					Args:    []string{"test"},
				},
			},
		},
	},
	{
		Input: DetectTestCommandInput{
			Ctx: nil,
			Input: &helpers.Input{
				RuntimeVersion: "",
				RootPath:       commonUtils.ProjectPath() + "/testingRepos/emptyRepo",
			},
			PluginDetails: &SupportedDependencies,
		},
		Output: &languageSpecific.TestCasesCommand{},
	},
}

type DetectCommandCase struct {
	Input  DetectCommandInput
	Output *languageSpecific.Commands
}

type DetectCommandInput struct {
	Ctx           context.Context
	Input         *helpers.Input
	PluginDetails *pluginDetails.LanguagePlugins
}

var DetectCommandCases = []DetectCommandCase{
	{
		Input: DetectCommandInput{
			Ctx: nil,
			Input: &helpers.Input{
				RuntimeVersion: "",
				RootPath:       commonUtils.ProjectPath() + "/testingRepos/commands/repo1",
			},
			PluginDetails: &SupportedDependencies,
		},
		Output: &languageSpecific.Commands{
			SeedCommands: &languageSpecific.CommandOutput{
				Used: true,
				Commands: []*helpers.Command{
					{
						Command: "npm run ",
						Args:    []string{"seed:map"},
					},
					{
						Command: "npm run ",
						Args:    []string{"seed"},
					},
				},
			},
			BuildCommands: &languageSpecific.CommandOutput{
				Used: true,
				Commands: []*helpers.Command{
					{
						Command: "npm run ",
						Args:    []string{"build:sw"},
					},
					{
						Command: "npm run ",
						Args:    []string{"build"},
					},
					{
						Command: "npm run ",
						Args:    []string{"postbuild"},
					},
				},
			},
			MigrationCommands: &languageSpecific.CommandOutput{
				Used: true,
				Commands: []*helpers.Command{
					{
						Command: "npm run ",
						Args:    []string{"migration:sw"},
					},
				},
			},
			StartUpCommands: &languageSpecific.CommandOutput{
				Used: true,
				Commands: []*helpers.Command{
					{
						Command: "npm ",
						Args:    []string{"start"},
					},
				},
			},
			AdHocScripts: &languageSpecific.CommandOutput{
				Used: true,
				Commands: []*helpers.Command{
					{
						Command: "npm run ",
						Args:    []string{"adhoc:so"},
					},
				},
			},
		},
	},
	{
		Input: DetectCommandInput{
			Ctx: nil,
			Input: &helpers.Input{
				RuntimeVersion: "",
				RootPath:       commonUtils.ProjectPath() + "/testingRepos/emptyRepo",
			},
			PluginDetails: &SupportedDependencies,
		},
		Output: &languageSpecific.Commands{
			SeedCommands: &languageSpecific.CommandOutput{
				Used:     false,
				Commands: []*helpers.Command{},
			},
			BuildCommands: &languageSpecific.CommandOutput{
				Used:     false,
				Commands: []*helpers.Command{},
			},
			MigrationCommands: &languageSpecific.CommandOutput{
				Used:     false,
				Commands: []*helpers.Command{},
			},
			StartUpCommands: &languageSpecific.CommandOutput{
				Used:     false,
				Commands: []*helpers.Command{},
			},
			AdHocScripts: &languageSpecific.CommandOutput{
				Used:     false,
				Commands: []*helpers.Command{},
			},
		},
	},
}
