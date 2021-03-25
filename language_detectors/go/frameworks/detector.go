package frameworks

import (
	"code-analyser/language_detectors/go/frameworks/beego"
	"code-analyser/language_detectors/interfaces"
)

var frameworks = map[string]interfaces.Framework{
	"main": &beego.BeegoFramework{},
}

type GOFrameworkDetector struct {
}

func (d *GOFrameworkDetector) GetLibrariesUsed(runtimeVersion, root string) *map[string]string {
	return &map[string]string{
		"main": "1.3",
	}
}

func (d *GOFrameworkDetector) GetDetector(libraryVersion, root, libraryUsed string) interfaces.Framework {
	return frameworks[libraryUsed]
}
