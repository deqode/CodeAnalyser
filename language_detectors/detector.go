package language_detectors

import (
	_go "code-analyser/language_detectors/go"
	"code-analyser/language_detectors/interfaces"
)

var MapLanguageToDetectors = map[string]interfaces.LanguageSpecificDetector{
	GO:         &_go.GoDetector{},
	JAVASCRIPT: nil,
	PYTHON:     nil,
}

const (
	GO         = "Go"
	PYTHON     = "Python"
	JAVASCRIPT = "JavaScript"
)
