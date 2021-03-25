package orms

import (
	"code-analyser/language_detectors/go/orms/gorm"
	"code-analyser/language_detectors/interfaces"
)

var orms = map[string]interfaces.ORM{
	"gorm": &gorm.GormORM{},
}

type GoORMDetector struct {
}

func (g *GoORMDetector) GetLibraryUsed(runtimeVersion, root string) *map[string]string {
	return &map[string]string{
		"gorm": "1.5",
	}
}

func (g *GoORMDetector) GetDetector(ORMVersion, root, libraryUsed string) interfaces.ORM {
	return orms[libraryUsed]
}
