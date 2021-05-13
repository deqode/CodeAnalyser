package parseDependencies

import (
	"code-analyser/protos/pb/versions"
	"code-analyser/runners"
	"golang.org/x/net/context"
)

type Case struct {
	Input  Input
	Output map[string]map[string]runners.DependencyDetail
}

type Input struct {
	Ctx             context.Context
	LanguageVersion string
	Path            string
	PluginDetails   *versions.LanguageVersion
}

var Cases = []Case{
	{
		Input: Input{
			Ctx:             nil,
			LanguageVersion: "",
			Path:            "",
			PluginDetails:   nil,
		},
		Output: map[string]map[string]runners.DependencyDetail{},
	},
}
