package loadPLugins

import (
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
)

type TestDbPluginLoadCase struct {
	Input  *pbUtils.Details
	Output *DbPlugin
}

var TestDbPluginLoadCases = []TestDbPluginLoadCase{
	{
		Input: &pbUtils.Details{
			Type:    "database",
			Name:    "postgres",
			Version: "v1.x",
			Semver:  []string{">=1.x,<2.x"},
			Command: "node " + utils.RootDirPath() + "/plugin/js/db/postgres/v1_x/server.js",
			Libraries: []*pbUtils.Library{
				{
					Name: "pg", Semver: []string{">=1.x,<9.x"},
				},
				{
					Name: "postgres", Semver: []string{">=1.x,<2.x"},
				},
				{
					Name: "postgresql-client", Semver: []string{">=1.x,<2.x"},
				},
				{
					Name: "pg-connection-string", Semver: []string{">=1.x,<3.x"},
				},
			},
		},
		Output: &DbPlugin{
			Dbs: map[string]*DbVersion{
				"postgres": {
					Version: map[string]*DbPluginDetails{
						"v1.x": {
							Libraries: []*pbUtils.Library{
								{
									Name: "pg", Semver: []string{">=1.x,<9.x"},
								},
								{
									Name: "postgres", Semver: []string{">=1.x,<2.x"},
								},
								{
									Name: "postgresql-client", Semver: []string{">=1.x,<2.x"},
								},
								{
									Name: "pg-connection-string", Semver: []string{">=1.x,<3.x"},
								},
							},
							Setting: &Set,
						},
					},
				},
			},
			Setting: &Set,
		},
	},
	{
		Input: &pbUtils.Details{
			Type:    "database",
			Name:    "redis",
			Version: "v1.x",
			Semver:  []string{">=1.x,<4.x"},
			Command: "node " + utils.RootDirPath() + "/plugin/js/db/redis/v1_x/server.js",
			Libraries: []*pbUtils.Library{
				{
					Name:   "redis",
					Semver: []string{">=1.x,<4.x"},
				},
				{
					Name:   "handy-redis",
					Semver: []string{">=1.x,<3.x"},
				},
				{
					Name:   "ioredis",
					Semver: []string{">=1.x,<5.x"},
				},
				{
					Name:   "redis-fast-driver",
					Semver: []string{">=1.x,<3.x"},
				},
				{
					Name:   "tedis",
					Semver: []string{">=0.x,<1.x"},
				},
			},
		},
		Output: &DbPlugin{
			Dbs: map[string]*DbVersion{
				"redis": {
					Version: map[string]*DbPluginDetails{
						"v1.x": {
							Libraries: []*pbUtils.Library{
								{
									Name:   "redis",
									Semver: []string{">=1.x,<4.x"},
								},
								{
									Name:   "handy-redis",
									Semver: []string{">=1.x,<3.x"},
								},
								{
									Name:   "ioredis",
									Semver: []string{">=1.x,<5.x"},
								},
								{
									Name:   "redis-fast-driver",
									Semver: []string{">=1.x,<3.x"},
								},
								{
									Name:   "tedis",
									Semver: []string{">=0.x,<1.x"},
								},
							},
							Setting: &Set,
						},
					},
				},
			},
			Setting: &Set,
		},
	},
}

/*
var DbRunnerCases = []DbRunnerCase{
	{
		Input: DbRunnerInput{
			DbList: map[string]DependencyDetail{
				"mongodb": {
					Version: "v1.x",
					Command: "node " + utils.RootDirPath() + "/plugin/js/db/mongodb/v1_x/server.js",
				},
			},
			RuntimeVersion: "",
			Root:           utils.RootDirPath() + "/testingRepos/db/repo1",
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
			Root:           utils.RootDirPath() + "/testingRepos/emptyRepo",
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
					Command: "node " + utils.RootDirPath() + "/plugin/js/db/mariadb/v1_x/server.js",
				},
				"redis": {
					Version: "v1.x",
					Command: "node " + utils.RootDirPath() + "/plugin/js/db/redis/v1_x/server.js",
				},
			},
			RuntimeVersion: "",
			Root:           utils.RootDirPath() + "/testingRepos/db/repo2",
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
					Command: "node " + utils.RootDirPath() + "/plugin/js/db/mariadb/v1_x/server.js",
				},
				"redis": {
					Version: "v1.x",
					Command: "node " + utils.RootDirPath() + "/plugin/js/db/redis/v1_x/server.js",
				},
			},
			RuntimeVersion: "",
			Root:           utils.RootDirPath() + "/testingRepos/dbfsdsf/repo2",
		},
		Output: &languageSpecificPB.DBOutput{
			Used: false,
			Databases: []*languageSpecificPB.DB{},
		},
	},
}
*/
