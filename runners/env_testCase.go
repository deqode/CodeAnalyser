package runners

import (
	"code-analyser/protos/pb/output/languageSpecific"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
)

type EnvCase struct {
	Input  EnvInput
	Output *languageSpecific.EnvOutput
}

type EnvInput struct {
	pluginDetails  *versionsPB.LanguageVersion
	RuntimeVersion string
	Root           string
}

var EnvCases = []EnvCase{
	{
		Input: EnvInput{
			pluginDetails: &versionsPB.LanguageVersion{
				DetectEnvCommand: "node " + utils.ProjectPath() + "/plugin/js/env/server.js",
			},
			Root:           utils.ProjectPath() + "/testingRepos/env/repo1",
			RuntimeVersion: "sfs",
		},
		Output: &languageSpecific.EnvOutput{
			EnvUsed: true,
			EnvKeyValues: map[string]string{
				"jk":       "7",
				"fto":      "hjk",
				"yop":      "tyui",
				"Boss":     "1",
				"just_one": "4",
			},
		},
	},
	{
		Input: EnvInput{
			pluginDetails: &versionsPB.LanguageVersion{
				DetectEnvCommand: "node " + utils.ProjectPath() + "/plugin/js/env/server.js",
			},
			Root:           utils.ProjectPath() + "/testingRepos/emptyRepo",
			RuntimeVersion: "sfs",
		},
		Output: &languageSpecific.EnvOutput{
			EnvUsed: false,
		},
	},
	{
		Input: EnvInput{
			pluginDetails: &versionsPB.LanguageVersion{
				DetectEnvCommand: "node " + utils.ProjectPath() + "/plugin/js/env/server.js",
			},
			Root:           "/stfgsftestingRepos/emptyRepo",
			RuntimeVersion: "sfs",
		},
		Output: nil,
	},
	{
		Input: EnvInput{
			pluginDetails: &versionsPB.LanguageVersion{
				DetectEnvCommand: "node " + utils.ProjectPath() + "/plugin/js/env/server.js",
			},
			Root:           utils.ProjectPath() + "/testingRepos/env/repo2",
			RuntimeVersion: "sfs",
		},
		Output: &languageSpecific.EnvOutput{
			EnvUsed: true,
			EnvKeyValues: map[string]string{
				"joss":      "1",
				"justhkjhk": "sdf",
			},
		},
	},
}
