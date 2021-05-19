package runners

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	versionsPB "code-analyser/protos/pb/versions"
)

type FrameworkParsingCase struct {
	Input  FrameworkParsingInput
	Output map[string]DependencyDetail
}

type FrameworkParsingInput struct {
	DependenciesList map[string]string
	LangYamlObject   *versionsPB.LanguageVersion
}

var FrameworkParsingCases = []FrameworkParsingCase{
	{
		Input: FrameworkParsingInput{
			DependenciesList: map[string]string{},
			LangYamlObject:   &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: FrameworkParsingInput{
			DependenciesList: map[string]string{
				"gdfsgd": "dsfdsf",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: FrameworkParsingInput{
			DependenciesList: map[string]string{
				"next":    "10.1",
				"express": "4.5",
				"ddfd":    "4545",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"nextJs": {
				Version: "v1.x",
				Command: "node plugin/js/framework/next.js/v1_x/server.js",
			},
			"express": {
				Version: "v1.x",
				Command: "node plugin/js/framework/express/v1_x/server.js",
			},
		},
	},
	{
		Input: FrameworkParsingInput{
			DependenciesList: map[string]string{
				"vue": "2.5",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"vue": {
				Version: "v1.x",
				Command: "node plugin/js/framework/vue/v1_x/server.js",
			},
		},
	},
}

type FrameworkRunnerCase struct {
	Input  FrameworkRunnerInput
	Output []*languageSpecificPB.FrameworkOutput
}

type FrameworkRunnerInput struct {
	FrameworkList  map[string]DependencyDetail
	RuntimeVersion string
	Root           string
}

var FrameworkRunnerCases = []FrameworkRunnerCase{
	{
		Input: FrameworkRunnerInput{
			FrameworkList:  nil,
			RuntimeVersion: "",
			Root:           "",
		},
		Output: nil,
	},
}
