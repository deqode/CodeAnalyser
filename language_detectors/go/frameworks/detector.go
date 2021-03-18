package frameworks

import (
	"code-analyser/detector_helpers"
	"code-analyser/language_detectors/go/frameworks/beego"
	"code-analyser/language_detectors/go/parsers/go_mod"
	protos "code-analyser/protos/pb"
	"code-analyser/versions"
)

var frameworks = map[string]map[string]*versions.FrameworkVersionDetector{
	"beego": beego.BeegoVersions,
}


func getDetector(runtimeVersion, root string, languageVersion *protos.LanguageVersion) []*versions.FrameworkVersionDetector {
	var output []*versions.FrameworkVersionDetector
	mod, _ := go_mod.ParseGoMod(root + "/go.mod")
	libraryFound := mod.Required
	library := libraryFound[0]
	fwDetector := detector_helpers.FindFrameworkDetector(frameworks[library.Name], languageVersion.Framework)
	output = append(output, fwDetector)
	//return Detector
	return output
}

func DetectFrameWork(root, runtimeVersion string, languageVersion *protos.LanguageVersion) []*protos.FrameworkOutput {
	output := []*protos.FrameworkOutput{}
	detector := getDetector(root, runtimeVersion, languageVersion)
	fwDetected, _ := detector[0].Detector(nil, runtimeVersion, root)
	output = append(output, fwDetected)
	return output
}
