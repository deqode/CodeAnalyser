package runners

import (
	"code-analyser/protos/pb/output/global"
	"code-analyser/protos/pb/output/languageSpecific"
	pluginPb "code-analyser/protos/pb/plugin"
	"code-analyser/protos/pb/versions"
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
	PluginDetails   *versions.LanguageVersion
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
	YamlLangObject *versions.LanguageVersion
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
	Output *languageSpecific.TestCasesCommandOutput
}

type DetectTestCommandInput struct {
	Ctx           context.Context
	Input         *pluginPb.ServiceInput
	PluginDetails *versions.LanguageVersion
}

var DetectTestCommandCases = []DetectTestCommandCase{
	{
		Input: DetectTestCommandInput{
			Ctx:           nil,
			Input:         nil,
			PluginDetails: &SupportedDependencies,
		},
		Output: nil,
	},
}

type DetectCommandCase struct {
	Input  DetectCommandInput
	Output DetectCommandOutput
}

type DetectCommandOutput struct {
	SeedCommand      *global.SeedCommandsOutput
	BuildCommand     *global.BuildCommandsOutput
	MigrationCommand *global.MigrationCommandsOutput
	StartUpCommand   *global.StartUpCommandsOutput
	AdHocScripts     *global.AdHocScriptsOutput
}

type DetectCommandInput struct {
	Ctx           context.Context
	Input         *pluginPb.ServiceInput
	PluginDetails *versions.LanguageVersion
}
