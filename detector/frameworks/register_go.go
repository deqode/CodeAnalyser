package frameworks

import "code-analyser/detector/frameworks/_go"

var RegisterGO = LanguagesDetector{
	Frameworks: []*Framework{
		{"Gin", _go.DetectGin, ">=1.6.x,<=2.x"},
		{"Beego", _go.DetectBeego, ">=1.x,<=2.x"},
	},
	DBCheckers: []*DBChecker{
		{"postgres", _go.CheckPostgres, "*"},
	},
	Prerequisites:    _go.Pre,
	Orms:             []*OrmChecker{{"GORM", _go.Gorm}},
	SemverConstraint: "*",
}
