package frameworks

import (
	"code-analyser/language_detectors/go/frameworks/beego"
	"code-analyser/language_detectors/go/frameworks/gin"
	"code-analyser/language_detectors/interfaces"
)

var frameworks = map[string]interfaces.Framework{
	"beego": &beego.BeegoFramework{},
	"gin":   &gin.GinFramework{},
}

type GOFrameworkDetector struct {
}

func (d *GOFrameworkDetector) GetLibrariesUsed(runtimeVersion, root string) *map[string]string {
	return &map[string]string{
		"beego": "1.3",
		"gin":   "2.5",
	}
}

func (d *GOFrameworkDetector) GetDetector(libraryVersion, root, libraryUsed string) interfaces.Framework {
	return frameworks[libraryUsed]
}
