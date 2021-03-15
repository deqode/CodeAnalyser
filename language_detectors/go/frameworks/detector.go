package frameworks

import (
	"code-analyser/language_detectors/go/frameworks/beego"
	protos "code-analyser/protos/pb"
	"code-analyser/versions"
)

var frameworks = map[string]map[string]*versions.FrameworkVersionDetector{
	"beego": beego.BeegoVersions,
}

func GetDetector(root, lib string) *versions.FrameworkVersionDetector {
	//helper.GetLibraries(root)
	//helper.FindVersion(frameworks["beego"])
	//beego.BeegoVersions
	//return Detector
	return nil
}
func DetectFrameWork(detector *versions.FrameworkVersionDetector) *protos.FrameworkOutput {
	return nil
}
