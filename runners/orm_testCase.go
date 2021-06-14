package runners

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	"code-analyser/protos/pb/pluginDetails"
	"code-analyser/utils"
)

type OrmParsingCase struct {
	Input  OrmParsingInput
	Output map[string]DependencyDetail
}

type OrmParsingInput struct {
	DependenciesList map[string]string
	LangYamlObject   *pluginDetails.LanguagePlugins
}

var OrmParsingCases = []OrmParsingCase{
	{
		Input: OrmParsingInput{
			DependenciesList: map[string]string{},
			LangYamlObject:   &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: OrmParsingInput{
			DependenciesList: map[string]string{
				"sdgdfg": "sfd454",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: OrmParsingInput{
			DependenciesList: map[string]string{
				"bookshelf": "5.6",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"bookshelf": {
				Version: "v1.x",
				Command: "node plugin/js/orm/bookshelf/v1_x/server.js",
			},
		},
	},
	{
		Input: OrmParsingInput{
			DependenciesList: map[string]string{
				"sequelize": "6.8",
				"typeorm":   "0.8",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"sequelize": {
				Version: "v1.x",
				Command: "node plugin/js/orm/sequelize/v1_x/server.js",
			},
			"typeorm": {
				Version: "v1.x",
				Command: "node plugin/js/orm/typeorm/v1_x/server.js",
			},
		},
	},
}

type OrmRunnerCase struct {
	Input  OrmRunnerInput
	Output *languageSpecificPB.OrmOutput
}

type OrmRunnerInput struct {
	OrmList        map[string]DependencyDetail
	RuntimeVersion string
	Root           string
}

var OrmRunnerCases = []OrmRunnerCase{
	{
		Input: OrmRunnerInput{
			OrmList:        map[string]DependencyDetail{},
			RuntimeVersion: "",
			Root:           utils.ProjectPath() + "/testingRepos/emptyRepo",
		},
		Output: &languageSpecificPB.OrmOutput{
			Used: false,
			Orms: []*languageSpecificPB.ORM{},
		},
	},
	{
		Input: OrmRunnerInput{
			OrmList: map[string]DependencyDetail{
				"typeorm": {
					Version: "v1.x",
					Command: "node " + utils.ProjectPath() + "/plugin/js/orm/typeorm/v1_x/server.js",
				},
				"bookshelf": {
					Version: "v1.x",
					Command: "node " + utils.ProjectPath() + "/plugin/js/orm/bookshelf/v1_x/server.js",
				},
			},
			RuntimeVersion: "",
			Root:           "/tefsdfstdfdsfingRepos/empfdsfdstyRepo",
		},
		Output: &languageSpecificPB.OrmOutput{
			Used: false,
			Orms: []*languageSpecificPB.ORM{},
		},
	},
	{
		Input: OrmRunnerInput{
			OrmList: map[string]DependencyDetail{
				"sequelize": {
					Version: "v1.x",
					Command: "node " + utils.ProjectPath() + "/plugin/js/orm/sequelize/v1_x/server.js",
				},
				"bookshelf": {
					Version: "v1.x",
					Command: "node " + utils.ProjectPath() + "/plugin/js/orm/bookshelf/v1_x/server.js",
				},
			},
			RuntimeVersion: "",
			Root:           utils.ProjectPath() +"/testingRepos/orm/repo1",
		},
		Output: &languageSpecificPB.OrmOutput{
			Used: true,
			Orms: []*languageSpecificPB.ORM{
				{
					Name:    "sequelize",
					Version: "v1.x",
				},
				{
					Name:    "bookshelf",
					Version: "v1.x",
				},
			},
		},
	},
	{
		Input: OrmRunnerInput{
			OrmList: map[string]DependencyDetail{
				"typeorm": {
					Version: "v1.x",
					Command: "node " + utils.ProjectPath() + "/plugin/js/orm/typeorm/v1_x/server.js",
				},
			},
			RuntimeVersion: "",
			Root:           utils.ProjectPath() + "/testingRepos/orm/repo2",
		},
		Output: &languageSpecificPB.OrmOutput{
			Used: true,
			Orms: []*languageSpecificPB.ORM{
				{
					Name:    "typeorm",
					Version: "v1.x",
				},
			},
		},
	},
}
