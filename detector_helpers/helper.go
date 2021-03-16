package detector_helpers

import (
	postgress "code-analyser/language_detectors/go/databases/postgres"
	"code-analyser/language_detectors/go/frameworks/beego"
	protos "code-analyser/protos/pb"
	"code-analyser/versions"
)

func FindFrameworkDetector(map[string]*versions.FrameworkVersionDetector, []*protos.Versions) *versions.FrameworkVersionDetector {
	return &versions.FrameworkVersionDetector{
		Default:  true,
		Name:     "v1.x",
		Detector: beego.Beego_v_1_6,
	}
}

func FindDBDetector(map[string]*versions.DBVersionDetector, []*protos.Versions) *versions.DBVersionDetector {
	return &versions.DBVersionDetector{
		Default:  true,
		Name:     "v12",
		Detector: postgress.Postgress_v_12,
	}
}
