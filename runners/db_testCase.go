package runners

import versionsPB "code-analyser/protos/pb/versions"

type DbCase struct {
	Input  DbInput
	Output map[string]DependencyDetail
}

type DbInput struct {
	DependenciesList map[string]string
	LangYamlObject   *versionsPB.LanguageVersion
}

var DbCases = []DbCase{
	{
		Input: DbInput{
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
		Input: DbInput{
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
		Input: DbInput{
			DependenciesList: map[string]string{
				"cassandra-driver": "4.5",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"cassandra": {
				Version: "v1.x",
				Command: "node plugin/js/db/cassandra/v1_x/server.js",
			},
		},
	},
	{
		Input: DbInput{
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
		Input: DbInput{
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
		Input: DbInput{
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
		Input: DbInput{
			DependenciesList: map[string]string{
				"mysql2": "2.4",
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
		Input: DbInput{
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
		Input: DbInput{
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
		Input: DbInput{
			DependenciesList: map[string]string{
				"postgres": "1.8",
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
		Input: DbInput{
			DependenciesList: map[string]string{
				"postgresql-client": "1.8",
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
		Input: DbInput{
			DependenciesList: map[string]string{
				"pg-connection-string": "2.4",
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
		Input: DbInput{
			DependenciesList: map[string]string{
				"redis": "3.4",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"redis": {
				Version: "v1.x",
				Command: "node plugin/js/db/redis/v1_x/server.js",
			},
		},
	},
	{
		Input: DbInput{
			DependenciesList: map[string]string{
				"handy-redis": "2.4",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"redis": {
				Version: "v1.x",
				Command: "node plugin/js/db/redis/v1_x/server.js",
			},
		},
	},
	{
		Input: DbInput{
			DependenciesList: map[string]string{
				"ioredis": "4.4",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"redis": {
				Version: "v1.x",
				Command: "node plugin/js/db/redis/v1_x/server.js",
			},
		},
	},
	{
		Input: DbInput{
			DependenciesList: map[string]string{
				"redis-fast-driver": "2.8",
				"express":           "7.8",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"redis": {
				Version: "v1.x",
				Command: "node plugin/js/db/redis/v1_x/server.js",
			},
		},
	},
	{
		Input: DbInput{
			DependenciesList: map[string]string{
				"tedis": "0.8",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"redis": {
				Version: "v1.x",
				Command: "node plugin/js/db/redis/v1_x/server.js",
			},
		},
	},
	{
		Input: DbInput{
			DependenciesList: map[string]string{
				"sqlite": "4.8",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"sqlite": {
				Version: "v1.x",
				Command: "node plugin/js/db/sqlite/v1_x/server.js",
			},
		},
	},
	{
		Input: DbInput{
			DependenciesList: map[string]string{
				"sqlite3": "5.5",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"sqlite": {
				Version: "v1.x",
				Command: "node plugin/js/db/sqlite/v1_x/server.js",
			},
		},
	},
	{
		Input: DbInput{
			DependenciesList: map[string]string{
				"better-sqlite3": "7.8",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"sqlite": {
				Version: "v1.x",
				Command: "node plugin/js/db/sqlite/v1_x/server.js",
			},
		},
	},
	{
		Input: DbInput{
			DependenciesList: map[string]string{
				"sql.js": "1.8",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"sqlite": {
				Version: "v1.x",
				Command: "node plugin/js/db/sqlite/v1_x/server.js",
			},
		},
	},
	{
		Input: DbInput{
			DependenciesList: map[string]string{},
			LangYamlObject:   &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: DbInput{
			DependenciesList: map[string]string{
				"dfgdgdfs": "fgdfsd78789",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
}
