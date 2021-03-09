package frameworks

import (
	"code-analyser/detector/frameworks/_go"
)

var RegisterGO = LanguagesDetector{
	Frameworks: []*Framework{
		{"Gin", _go.DetectGin, map[string]string{"lowest": "1.x"}},
		{"Beego", _go.DetectBeego, map[string]string{"lowest": "1.x"}},
	},
	DBCheckers: []*DBChecker{
		{"postgres", _go.CheckPostgres, "*"},
	},
	Prerequisites: _go.Pre,
	Orms:          []*OrmChecker{{"GORM", _go.Gorm}},
}
