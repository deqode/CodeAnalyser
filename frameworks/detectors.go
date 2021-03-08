package frameworks

import (
	"code-analyser/detector"
	"code-analyser/frameworks/_go"
)

type FrameworkDetectOutput func(string) *detector.DetectedOutput

type LanguagesDetector struct {
	Frameworks    []*Framework
	DBCheckers    []*DBChecker
	Prerequisites func(string)
	Orms          []*OrmChecker
}

type Framework struct {
	Name     string
	Detector FrameworkDetectOutput
	Version  string
}

type DBChecker struct {
	Name     string
	Detector func(string) (string, bool, string)
	Version  string
}

type OrmChecker struct {
	Name     string
	Detector func(root string) bool
}

var Languages = map[string]LanguagesDetector{
	"Go": {
		Frameworks: []*Framework{
			{Name: "Gin", Detector: _go.DetectGin, Version: "*"},
			{Name: "Beego", Detector: _go.DetectBeego, Version: "*"},
		},
		DBCheckers: []*DBChecker{
			{
				Name:     "postgres",
				Version:  "*",
				Detector: _go.CheckPostgres,
			},
		},
		Prerequisites: _go.Pre,
		Orms: []*OrmChecker{
			{"GORM", _go.Gorm},
		},
	},
}
