package frameworks

// Todo: Add version range
import (
	"code-analyser/detector/protos"
)

type FrameworkDetectOutput func(string) *protos.DetectedOutput

type LanguagesDetector struct {
	Frameworks    []*Framework
	DBCheckers    []*DBChecker
	Prerequisites func(string)
	Orms          []*OrmChecker
}

type Framework struct {
	Name         string
	Detector     FrameworkDetectOutput //Todo: pointer
	VersionRange map[string]string
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

var Languages = map[string]*LanguagesDetector{
	"Go": &RegisterGO,
}
