package runners

import (
	"code-analyser/protos/pb/output/utils"
	"code-analyser/protos/pb/versions"
	commonUtils "code-analyser/utils"
)

var SupportedDependencies = versions.LanguageVersion{
	Detectruntimecommand: "node " + commonUtils.ProjectPath() + "/plugin/js/detectRunTime/server.js",
	Runtimeversions: map[string]*versions.DependencyVersionDetails{"v1.x": {
		Semver:        []string{">=1.x,<16.x"},
		Plugincommand: "node " + commonUtils.ProjectPath() + "/plugin/js/getDependencies/server.js",
	}},
	Framework: map[string]*versions.DependencyDetails{
		"Gatsby v3": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "gatsby",
							Semver: []string{">=1.x,<4.x"},
						},
					},
					Plugincommand: "node plugin/js/framework/gatsby/v1_x/server.js",
				},
			},
		},
		"angular": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "@angular/core",
							Semver: []string{">=1.x,<12.x"},
						},
					},
					Plugincommand: "node plugin/js/framework/angular/v1_x/server.js",
				},
			},
		},
		"express": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "express",
							Semver: []string{">=1.x,<5.x"},
						},
					},
					Plugincommand: "node plugin/js/framework/express/v1_x/server.js",
				},
			},
		},
		"nestJs": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "@nestjs/core",
							Semver: []string{">=1.x,<8.x"},
						},
					},
					Plugincommand: "node plugin/js/framework/nest/v1_x/server.js",
				},
			},
		},
		"nextJs": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "next",
							Semver: []string{">=1.x,<11.x"},
						},
					},
					Plugincommand: "node plugin/js/framework/next.js/v1_x/server.js",
				},
			},
		},
		"vue": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "vue",
							Semver: []string{">=1.x,<3.x"},
						},
					},
					Plugincommand: "node plugin/js/framework/vue/v1_x/server.js",
				},
			},
		},
	},
	Databases: map[string]*versions.DependencyDetails{
		"cassandra": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "cassandra-driver",
							Semver: []string{">=1.x,<5.x"},
						},
					},
					Plugincommand: "node plugin/js/db/cassandra/v1_x/server.js",
				},
			},
		},
		"MariaDB": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "mariadb",
							Semver: []string{">=1.x,<3.x"},
						},
					},
					Plugincommand: "node plugin/js/db/mariadb/v1_x/server.js",
				},
			},
		},
		"Microsoft SQL server": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "tedious",
							Semver: []string{">=1.x,<12.x"},
						},
					},
					Plugincommand: "node plugin/js/db/microsoftSQLServer/v1_x/server.js",
				},
			},
		},
		"mongoDB": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "mongodb",
							Semver: []string{">=1.x,<4.x"},
						},
						{
							Name:   "mongoose",
							Semver: []string{">=1.x,<6.x"},
						},
					},
					Plugincommand: "node plugin/js/db/mongodb/v1_x/server.js",
				},
			},
		},
		"mysql": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "mysql",
							Semver: []string{">=1.x,<3.x"},
						},
						{
							Name:   "mysql2",
							Semver: []string{">=1.x,<3.x"},
						},
					},
					Plugincommand: "node plugin/js/db/mysql/v1_x/server.js",
				},
			},
		},
		"oracle": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "oracledb",
							Semver: []string{">=1.x,<6.x"},
						},
					},
					Plugincommand: "node plugin/js/db/oracle/v1_x/server.js",
				},
			},
		},
		"postgres": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "pg",
							Semver: []string{">=1.x,<9.x"},
						},
						{
							Name:   "postgres",
							Semver: []string{">=1.x,<2.x"},
						},
						{
							Name:   "postgresql-client",
							Semver: []string{">=1.x,<2.x"},
						},
						{
							Name:   "pg-connection-string",
							Semver: []string{">=1.x,<3.x"},
						},
					},
					Plugincommand: "node plugin/js/db/postgres/v1_x/server.js",
				},
			},
		},
		"redis": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
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
					Plugincommand: "node plugin/js/db/redis/v1_x/server.js",
				},
			},
		},
		"sqlite": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Libraries: []*utils.Library{
						{
							Name:   "sqlite",
							Semver: []string{">=1.x,<5.x"},
						},
						{
							Name:   "sqlite3",
							Semver: []string{">=1.x,<6.x"},
						},
						{
							Name:   "better-sqlite3",
							Semver: []string{">=1.x,<8.x"},
						},
						{
							Name:   "sql.js",
							Semver: []string{">=1.x,<2.x"},
						},
					},
					Plugincommand: "node plugin/js/db/sqlite/v1_x/server.js",
				},
			},
		},
	},
	Orms: map[string]*versions.DependencyDetails{
		"typeorm": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Semver:        []string{">=0.x,<1.x"},
					Plugincommand: "node plugin/js/orm/typeorm/v1_x/server.js",
				},
			},
		},
		"bookshelf": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Semver:        []string{">=1.x,<6.x"},
					Plugincommand: "node plugin/js/orm/bookshelf/v1_x/server.js",
				},
			},
		},
		"sequelize": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Semver:        []string{">=1.x,<7.x"},
					Plugincommand: "node plugin/js/orm/sequelize/v1_x/server.js",
				},
			},
		},
	},
	Libraries: map[string]*versions.DependencyDetails{
		"react": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Semver:        []string{">=1.x,<18.x"},
					Plugincommand: "node plugin/js/libraries/react/v1_x/server.js",
				},
			},
		},
		"kafka-node": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Semver:        []string{">=1.x,<6.x"},
					Plugincommand: "node plugin/js/libraries/kafka/v1_x/server.js",
				},
			},
		},
		"knex": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Semver:        []string{">=0.x,<1.x"},
					Plugincommand: "node plugin/js/libraries/knex/v1_x/server.js",
				},
			},
		},
		"mongoose": {
			Version: map[string]*versions.DependencyVersionDetails{
				"v1.x": {
					Semver:        []string{">=1.x,<6.x"},
					Plugincommand: "node plugin/js/libraries/mongoose/v1_x/server.js",
				},
			},
		},
	},
	Dependencies:           nil,
	DetectEnvCommand:       "node "+commonUtils.ProjectPath()+"/plugin/js/env/server.js",
	PreDetectCommand:       "",
	StaticAssetsCommand:    "",
	BuildDirectoryCommand:  "",
	DetectTestCasesCommand: "node "+commonUtils.ProjectPath()+"/plugin/js/testCasesCommands/server.js",
	Commands:               "node "+commonUtils.ProjectPath()+"/plugin/js/commands/server.js",
}
