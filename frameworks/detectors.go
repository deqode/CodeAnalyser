package frameworks

import (
	"code-analyser/detector"
	"code-analyser/frameworks/_go"
)

type FrameWork struct{
	Name string
	Detector  func(string) *detector.DetectedOutput
}

var Languages= map[string][]FrameWork{
	"Go":{{Name: "Gin", Detector: _go.DetectGin},
		{Name: "Beego", Detector: _go.DetectBeego},
		},
}


