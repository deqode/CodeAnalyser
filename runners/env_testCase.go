package runners

import (
	"code-analyser/protos/pb/output/languageSpecific"
	"code-analyser/protos/pb/pluginDetails"
	"code-analyser/utils"
)

type EnvCase struct {
	Input  EnvInput
	Output *languageSpecific.Envs
}

type EnvInput struct {
	pluginDetails  *pluginDetails.LanguagePlugins
	RuntimeVersion string
	Root           string
}

var EnvCases = []EnvCase{
	{
		Input: EnvInput{
			pluginDetails: &SupportedDependencies,
			Root:           utils.RootDirPath() + "/testingRepos/env/repo1",
			RuntimeVersion: "sfs",
		},
		Output: &languageSpecific.Envs{
			Used: true,
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
			pluginDetails: &SupportedDependencies,
			Root:           utils.RootDirPath() + "/testingRepos/emptyRepo",
			RuntimeVersion: "sfs",
		},
		Output: &languageSpecific.Envs{
			Used: false,
		},
	},
	{
		Input: EnvInput{
			pluginDetails: &SupportedDependencies,
			Root:           "/stfgsftestingRepos/emptyRepo",
			RuntimeVersion: "sfs",
		},
		Output: nil,
	},
	{
		Input: EnvInput{
			pluginDetails: &SupportedDependencies,
			Root:           utils.RootDirPath() + "/testingRepos/env/repo2",
			RuntimeVersion: "sfs",
		},
		Output: &languageSpecific.Envs{
			Used: true,
			EnvKeyValues: map[string]string{
				"joss":      "1",
				"justhkjhk": "sdf",
			},
		},
	},
}
