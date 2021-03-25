package beego

import (
	"code-analyser/language_detectors/go/frameworks/beego/V_1_X"
	"code-analyser/language_detectors/interfaces"
)

var beegoVersions = []*interfaces.FrameworkVersionDetector{
	{
		Default:  true,
		Name:     "v1.x",
		Detector: &V_1_X.Beego_V_1_X{},
	},
}

type BeegoFramework struct {
	Version []*interfaces.FrameworkVersionDetector
}

func (f *BeegoFramework) GetVersionedDetector(runtimeVersion, languageVersionFile, root string) interfaces.FrameworkVersionDetector {
	///home/deqode/Documents/code-analyser/plugin/go/main

	////TODO: need to add semver check and return correct version
	//pluginClient.FrameworkPluginCall(exec.Command("sh", "-c", "./plugin/go/beego/main"))

	return *beegoVersions[0]
}
