package runners

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
)

type DbParsingCase struct {
	Input  DbParsingInput
	Output map[string]DependencyDetail
}

type DbParsingInput struct {
	DependenciesList map[string]string
	LangYamlObject   *versionsPB.LanguageVersion
}

var DbCases = []DbParsingCase{
	{
		Input: DbParsingInput{
			DependenciesList: map[string]string{
				"mongodb":  "3.5",
				"mongoose": "5.5",
				"express":  "6.9",
				"alp":      "gdfd",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"mongoDB": {
				Version: "v1.x",
				Command: "node plugin/js/db/mongodb/v1_x/server.js",
			},
		},
	},
	{
		Input: DbParsingInput{
			DependenciesList: map[string]string{
				"mongoose": "5.5",
				"express":  "6.9",
				"alp":      "gdfd",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"mongoDB": {
				Version: "v1.x",
				Command: "node plugin/js/db/mongodb/v1_x/server.js",
			},
		},
	},
	{
		Input: DbParsingInput{
			DependenciesList: map[string]string{
				"mariadb": "2.5",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"MariaDB": {
				Version: "v1.x",
				Command: "node plugin/js/db/mariadb/v1_x/server.js",
			},
		},
	},
	{
		Input: DbParsingInput{
			DependenciesList: map[string]string{
				"tedious": "11.3",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"Microsoft SQL server": {
				Version: "v1.x",
				Command: "node plugin/js/db/microsoftSQLServer/v1_x/server.js",
			},
		},
	},
	{
		Input: DbParsingInput{
			DependenciesList: map[string]string{
				"mysql": "2.5",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"mysql": {
				Version: "v1.x",
				Command: "node plugin/js/db/mysql/v1_x/server.js",
			},
		},
	},
	{
		Input: DbParsingInput{
			DependenciesList: map[string]string{
				"oracledb": "5.5",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"oracle": {
				Version: "v1.x",
				Command: "node plugin/js/db/oracle/v1_x/server.js",
			},
		},
	},
	{
		Input: DbParsingInput{
			DependenciesList: map[string]string{
				"pg": "8.5",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"postgres": {
				Version: "v1.x",
				Command: "node plugin/js/db/postgres/v1_x/server.js",
			},
		},
	},
	{
		Input: DbParsingInput{
			DependenciesList: map[string]string{},
			LangYamlObject:   &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: DbParsingInput{
			DependenciesList: map[string]string{
				"dfgdgdfs": "fgdfsd78789",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
}

type DbRunnerCase struct {
	Input  DbRunnerInput
	Output *languageSpecificPB.DBOutput
}

type DbRunnerInput struct {
	DbList         map[string]DependencyDetail
	RuntimeVersion string
	Root           string
}

var DbRunnerCases = []DbRunnerCase{
	{
		Input: DbRunnerInput{
			DbList: map[string]DependencyDetail{
				"mongodb": {
					Version: "v1.x",
					Command: "node " + utils.ProjectPath() + "/plugin/js/db/mongodb/v1_x/server.js",
				},
			},
			RuntimeVersion: "",
			Root:           utils.ProjectPath() + "/testingRepos/db/repo1",
		},
		Output: &languageSpecificPB.DBOutput{
			Used: true,
			Databases: []*languageSpecificPB.DB{
				{
					Name:    "mongodb",
					Version: "v1.x",
				},
			},
		},
	},
	{
		Input: DbRunnerInput{
			DbList:         map[string]DependencyDetail{},
			RuntimeVersion: "",
			Root:           utils.ProjectPath() + "/testingRepos/emptyRepo",
		},
		Output: &languageSpecificPB.DBOutput{
			Used:      false,
			Databases: []*languageSpecificPB.DB{},
		},
	},
	{
		Input: DbRunnerInput{
			DbList: map[string]DependencyDetail{
				"mariadb": {
					Version: "v1.x",
					Command: "node " + utils.ProjectPath() + "/plugin/js/db/mariadb/v1_x/server.js",
				},
				"redis": {
					Version: "v1.x",
					Command: "node " + utils.ProjectPath() + "/plugin/js/db/redis/v1_x/server.js",
				},
			},
			RuntimeVersion: "",
			Root:           utils.ProjectPath() + "/testingRepos/db/repo2",
		},
		Output: &languageSpecificPB.DBOutput{
			Used: true,
			Databases: []*languageSpecificPB.DB{
				{
					Name:    "mariadb",
					Version: "v1.x",
				},
				{
					Name:    "redis",
					Version: "v1.x",
				},
			},
		},
	},
	{
		Input: DbRunnerInput{
			DbList: map[string]DependencyDetail{
				"mariadb": {
					Version: "v1.x",
					Command: "node " + utils.ProjectPath() + "/plugin/js/db/mariadb/v1_x/server.js",
				},
				"redis": {
					Version: "v1.x",
					Command: "node " + utils.ProjectPath() + "/plugin/js/db/redis/v1_x/server.js",
				},
			},
			RuntimeVersion: "",
			Root:           utils.ProjectPath() + "/testingRepos/dbfsdsf/repo2",
		},
		Output: &languageSpecificPB.DBOutput{
			Used: false,
			Databases: []*languageSpecificPB.DB{},
		},
	},
}
