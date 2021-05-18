package runners

import (
	versionsPB "code-analyser/protos/pb/versions"
)

type OrmCase struct {
	Input  OrmInput
	Output map[string]DependencyDetail
}

type OrmInput struct {
	DependenciesList map[string]string
	LangYamlObject   *versionsPB.LanguageVersion
}

var OrmCases = []OrmCase{
	{
		Input: OrmInput{
			DependenciesList: map[string]string{},
			LangYamlObject:   &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: OrmInput{
			DependenciesList: map[string]string{
				"sdgdfg": "sfd454",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: OrmInput{
			DependenciesList: map[string]string{
				"bookshelf": "5.6",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"bookshelf": {
				Version: "v1.x",
				Command: "node plugin/js/orm/bookshelf/v1_x/server.js",
			},
		},
	},
	{
		Input: OrmInput{
			DependenciesList: map[string]string{
				"sequelize": "6.8",
				"typeorm":   "0.8",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"sequelize": {
				Version: "v1.x",
				Command: "node plugin/js/orm/sequelize/v1_x/server.js",
			},
			"typeorm": {
				Version: "v1.x",
				Command: "node plugin/js/orm/typeorm/v1_x/server.js",
			},
		},
	},
}