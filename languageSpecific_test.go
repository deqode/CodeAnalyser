package main

import (
	decisionmakerPB "code-analyser/protos/pb"
	"code-analyser/protos/pb/output/global"
	"code-analyser/protos/pb/output/languageSpecific"
	"code-analyser/utils"
	testingUtil "code-analyser/utils/testing"
	"reflect"
	"strconv"
	"testing"
)

func TestLanguageSpecific(t *testing.T) {
	for i, element := range LanguageSpecificTestCases {
		input := element.Input
		outputArray := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			gotArray := Scrape(input).LanguageSpecificDetection

			outputMap := map[string]*decisionmakerPB.LanguageSpecificDetections{}
			for _, languageData := range outputArray {
				outputMap[languageData.Name] = languageData
			}

			for _, got := range gotArray {
				output := outputMap[got.Name]
				t.Run(got.Name, func(t *testing.T) {
					t.Parallel()
					CheckIndividualDependencies(got, output, t)
				})
			}

		})
	}
}

func CheckIndividualDependencies(got, output *decisionmakerPB.LanguageSpecificDetections, t *testing.T) {
	t.Run("language name", func(t *testing.T) {
		t.Parallel()
		if got.Name != output.Name {
			t.Error("expected this ", output.Name, "\n got this ", got.Name)
		}
	})
	t.Run("language version", func(t *testing.T) {
		t.Parallel()
		if got.RuntimeVersion != output.RuntimeVersion {
			t.Error("expected this ", output.RuntimeVersion, "\n got this ", got.RuntimeVersion)
		}
	})
	t.Run("commands", func(t *testing.T) {
		t.Parallel()
		testingUtil.AllCommandTesting(t, got.Commands, output.Commands)
	})
	t.Run("databases", func(t *testing.T) {
		t.Parallel()
		testingUtil.CheckDbEquality(got.Db, output.Db, t)
	})
	t.Run("envs", func(t *testing.T) {
		t.Parallel()
		if !reflect.DeepEqual(got.Env, output.Env) {
			t.Error("expected this ", output, "\n got this ", got)
		}
	})
	t.Run("framework", func(t *testing.T) {
		t.Parallel()
		testingUtil.CheckFrameworkEquality(got.Framework, output.Framework, t)
	})
	t.Run("orm", func(t *testing.T) {
		t.Parallel()
		testingUtil.CheckOrmEquality(got.Orm, output.Orm, t)
	})
	t.Run("libraries", func(t *testing.T) {
		t.Parallel()
		testingUtil.CheckLibraryEquality(got.Libraries, output.Libraries, t)
	})
}

type LanguageSpecificTestCase struct {
	Input  string
	Output []*decisionmakerPB.LanguageSpecificDetections
}

var LanguageSpecificTestCases = []LanguageSpecificTestCase{
	{
		Input: utils.ProjectPath() + "/testingRepos/languageSpecific/covid19india-react-master",
		Output: []*decisionmakerPB.LanguageSpecificDetections{
			{
				Name:           "JavaScript",
				RuntimeVersion: "12.3.0",
				Env: &languageSpecific.EnvOutput{
					EnvUsed: true,
					EnvKeyValues: map[string]string{
						"ESLINT_NO_DEV_ERRORS": "true",
					},
					Variables: nil,
				},
				Framework: nil,
				Db: &languageSpecific.DBOutput{
					Used:      false,
					Databases: nil,
				},
				Orm: &languageSpecific.OrmOutput{
					Used: false,
					Orms: nil,
				},
				Dependencies: nil,
				Libraries: []*languageSpecific.LibraryOutput{
					{
						Used:    true,
						Name:    "react",
						Version: "v1.x",
					},
				},
				StaticAssets:   nil,
				StackOutput:    nil,
				Appserver:      nil,
				BuildDirectory: nil,
				TestCases: &languageSpecific.TestCasesCommandOutput{
					Used: true,
					Commands: []*global.Command{
						{
							Command: "npm",
							Args:    []string{"test"},
						},
						{
							Command: "npm",
							Args:    []string{"test:map"},
						},
					},
				},
				Commands: &decisionmakerPB.Commands{
					BuildCommands: &global.BuildCommandsOutput{
						Used: true,
						BuildCommands: []*global.Command{
							{
								Command: "npm run ",
								Args:    []string{"build"},
							},
							{
								Command: "npm run ",
								Args:    []string{"postbuild"},
							},
							{
								Command: "npm run ",
								Args:    []string{"build:sw"},
							},
						},
					},
					StartUpCommands: &global.StartUpCommandsOutput{
						Used: true,
						StartUpCommands: []*global.Command{
							{
								Command: "npm ",
								Args:    []string{"start"},
							},
						},
					},
					SeedCommands: &global.SeedCommandsOutput{
						Used: true,
						SeedCommands: []*global.Command{
							{
								Command: "npm run ",
								Args:    []string{"seed"},
							},
							{
								Command: "npm run ",
								Args:    []string{"seed:map"},
							},
						},
					},
					MigrationCommands: &global.MigrationCommandsOutput{
						Used: true,
						MigrationCommands: []*global.Command{
							{
								Command: "npm run ",
								Args:    []string{"migration:sw"},
							},
						},
					},
					AdHocScriptsOutput: &global.AdHocScriptsOutput{
						Used: true,
						AdHocScripts: []*global.Command{
							{
								Command: "npm run ",
								Args:    []string{"adhoc:so"},
							},
						},
					},
				},
			},
		},
	},
	{
		Input: utils.ProjectPath() + "/testingRepos/languageSpecific/spectrum-alpha",
		Output: []*decisionmakerPB.LanguageSpecificDetections{
			{
				Name:           "JavaScript",
				RuntimeVersion: "12.3.0",
				Env:            &languageSpecific.EnvOutput{},
				Framework: []*languageSpecific.FrameworkOutput{
					{
						Used:    true,
						Name:    "express",
						Version: "v1.x",
					},
				},
				Db: &languageSpecific.DBOutput{
					Used: true,
					Databases: []*languageSpecific.DB{
						{
							Name:    "redis",
							Version: "v1.x",
						},
					},
				},
				Orm: &languageSpecific.OrmOutput{},
				Libraries: []*languageSpecific.LibraryOutput{
					{
						Used:    true,
						Name:    "react",
						Version: "v1.x",
					},
				},
				TestCases: &languageSpecific.TestCasesCommandOutput{
					Used: true,
					Commands: []*global.Command{
						{
							Command: "npm",
							Args:    []string{"test:desktop"},
						},
						{
							Command: "npm",
							Args:    []string{"test"},
						},
						{
							Command: "npm",
							Args:    []string{"test:ci"},
						},
						{
							Command: "npm",
							Args:    []string{"test:e2e"},
						},
						{
							Command: "npm",
							Args:    []string{"prestart:api:test"},
						},
						{
							Command: "npm",
							Args:    []string{"start:api:test"},
						},
						{
							Command: "npm",
							Args:    []string{"process:emails:test"},
						},
					},
				},
				Commands: &decisionmakerPB.Commands{
					BuildCommands: &global.BuildCommandsOutput{
						Used: true,
						BuildCommands: []*global.Command{
							{
								Command: "npm run ",
								Args:    []string{"prebuild:api"},
							},
							{
								Command: "npm run ",
								Args:    []string{"build:api"},
							},
							{
								Command: "npm run ",
								Args:    []string{"postbuild:api"},
							},
							{
								Command: "npm run ",
								Args:    []string{"prebuild:athena"},
							},
							{
								Command: "npm run ",
								Args:    []string{"build:athena"},
							},
							{
								Command: "npm run ",
								Args:    []string{"postbuild:athena"},
							},
							{
								Command: "npm run ",
								Args:    []string{"prebuild:hermes"},
							},
							{
								Command: "npm run ",
								Args:    []string{"build:hermes"},
							},
							{
								Command: "npm run ",
								Args:    []string{"postbuild:hermes"},
							},
							{
								Command: "npm run ",
								Args:    []string{"prebuild:chronos"},
							},
							{
								Command: "npm run ",
								Args:    []string{"build:chronos"},
							},
							{
								Command: "npm run ",
								Args:    []string{"postbuild:chronos"},
							},
							{
								Command: "npm run ",
								Args:    []string{"prebuild:mercury"},
							},
							{
								Command: "npm run ",
								Args:    []string{"build:mercury"},
							},
							{
								Command: "npm run ",
								Args:    []string{"postbuild:mercury"},
							},
							{
								Command: "npm run ",
								Args:    []string{"prebuild:vulcan"},
							},
							{
								Command: "npm run ",
								Args:    []string{"build:vulcan"},
							},
							{
								Command: "npm run ",
								Args:    []string{"postbuild:vulcan"},
							},
							{
								Command: "npm run ",
								Args:    []string{"prebuild:hyperion"},
							},
							{
								Command: "npm run ",
								Args:    []string{"build:hyperion"},
							},
							{
								Command: "npm run ",
								Args:    []string{"postbuild:hyperion"},
							},
							{
								Command: "npm run ",
								Args:    []string{"build:web"},
							},
							{
								Command: "npm run ",
								Args:    []string{"build:desktop"},
							},
						},
					},
					StartUpCommands: &global.StartUpCommandsOutput{
						Used: true,
						StartUpCommands: []*global.Command{
							{
								Command: "npm ",
								Args:    []string{"start"},
							},
						},
					},
					SeedCommands: &global.SeedCommandsOutput{
						Used: true,
						SeedCommands: []*global.Command{
							{
								Command: "npm run ",
								Args:    []string{"db:seed"},
							},
						},
					},
					MigrationCommands: &global.MigrationCommandsOutput{
						Used: true,
						MigrationCommands: []*global.Command{
							{
								Command: "npm run ",
								Args:    []string{"db:migrate"},
							},
							{
								Command: "npm run ",
								Args:    []string{"rethinkdb:migrate"},
							},
						},
					},
					AdHocScriptsOutput: &global.AdHocScriptsOutput{},
				},
			},
		},
	},
	{
		Input: utils.ProjectPath() + "/testingRepos/languageSpecific/vue-express-mongo-boilerplate-master",
		Output: []*decisionmakerPB.LanguageSpecificDetections{
			{
				Name:           "JavaScript",
				RuntimeVersion: "12.3.0",
				Env:            &languageSpecific.EnvOutput{},
				Framework: []*languageSpecific.FrameworkOutput{
					{
						Used:    true,
						Name:    "express",
						Version: "v1.x",
					},
					{
						Used:    true,
						Name:    "vue",
						Version: "v1.x",
					},
				},
				Db: &languageSpecific.DBOutput{
					Used: true,
					Databases: []*languageSpecific.DB{
						{
							Name:    "redis",
							Version: "v1.x",
						},
						{
							Name:    "mongoDB",
							Version: "v1.x",
						},
					},
				},
				Orm: &languageSpecific.OrmOutput{},
				Libraries: []*languageSpecific.LibraryOutput{
					{
						Used:    true,
						Name:    "mongoose",
						Version: "v1.x",
					},
				},
				TestCases: &languageSpecific.TestCasesCommandOutput{
					Used: true,
					Commands: []*global.Command{
						{
							Command: "npm",
							Args:    []string{"test:unit"},
						},
						{
							Command: "npm",
							Args:    []string{"test:e2e:nightmare"},
						},
						{
							Command: "npm",
							Args:    []string{"test:e2e"},
						},
						{
							Command: "npm",
							Args:    []string{"test"},
						},
					},
				},
				Commands: &decisionmakerPB.Commands{
					BuildCommands: &global.BuildCommandsOutput{
						Used: true,
						BuildCommands: []*global.Command{
							{
								Command: "npm run ",
								Args:    []string{"build"},
							},
							{
								Command: "npm run ",
								Args:    []string{"build:server"},
							},
						},
					},
					StartUpCommands: &global.StartUpCommandsOutput{
						Used: true,
						StartUpCommands: []*global.Command{
							{
								Command: "npm ",
								Args:    []string{"start"},
							},
						},
					},
					SeedCommands:       &global.SeedCommandsOutput{},
					MigrationCommands:  &global.MigrationCommandsOutput{},
					AdHocScriptsOutput: &global.AdHocScriptsOutput{},
				},
			},
		},
	},
}
