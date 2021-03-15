package beego

import (
	"code-analyser/versions"
)

var BeegoVersions = map[string]*versions.FrameworkVersionDetector{
	"v1.x": {
		Default: true,
		Name:     "v1.x",
		Detector: Beego_v_1_6,
	},
	"v2.x": {
		Name:     "v1.x",
		Detector: Beego_v_2_6,
	},
}