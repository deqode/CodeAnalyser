package runners

import (
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
