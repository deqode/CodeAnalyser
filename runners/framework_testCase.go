package runners

import versionsPB "code-analyser/protos/pb/versions"

type FrameworkCase struct {
	Input  FrameworkInput
	Output map[string]DependencyDetail
}

type FrameworkInput struct {
	DependenciesList map[string]string
	LangYamlObject   *versionsPB.LanguageVersion
}

var FrameworkCases = []FrameworkCase{
	{
		Input: FrameworkInput{
			DependenciesList: map[string]string{},
			LangYamlObject:   &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: FrameworkInput{
			DependenciesList: map[string]string{
				"gdfsgd": "dsfdsf",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: FrameworkInput{
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
		Input: FrameworkInput{
			DependenciesList: map[string]string{
				"vue":    "2.5",
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
