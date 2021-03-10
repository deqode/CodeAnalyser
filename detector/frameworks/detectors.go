package frameworks

// Todo: Add version range
import (
	"code-analyser/detector/protos"
)

type FrameworkDetectOutput func(string) *protos.DetectedOutput

type LanguagesDetector struct {
	Frameworks       []*Framework
	DBCheckers       []*DBChecker
	Prerequisites    func(string)
	Orms             []*OrmChecker
	SemverConstraint string
}

type Framework struct {
	Name             string
	Detector         FrameworkDetectOutput //Todo: pointer
	SemverConstraint string
}

type DBChecker struct {
	Name            string
	Detector        func(string) (string, bool, string)
	SemVerValidator string
}

type OrmChecker struct {
	Name     string
	Detector func(root string) bool
}

var Languages = map[string]*LanguagesDetector{
	"Go": &RegisterGO,
}
