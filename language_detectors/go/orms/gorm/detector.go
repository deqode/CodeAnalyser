package gorm

import (
	"code-analyser/language_detectors/go/orms/gorm/V_1_X"
	"code-analyser/language_detectors/interfaces"
)

var gormVersions = []*interfaces.ORMVersionDetector{
	{
		Default:  true,
		Name:     "v1.x",
		Semver:   "",
		Detector: &V_1_X.GORM_V_1_X{},
	},
}

type GormORM struct {
	Version []*interfaces.ORMVersionDetector
}

func (reciever *GormORM) GetVersionDetector(runtimeVersion, ormVersionFile, root string) interfaces.ORMVersionDetector {
	return *gormVersions[0]
}
