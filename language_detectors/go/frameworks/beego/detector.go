package beego

import (
	"code-analyser/language_detectors/go/frameworks/beego/V_1_X"
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient"
	"github.com/hashicorp/go-plugin"
	"os/exec"
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

func (f *BeegoFramework) GetVersionedDetector(runtimeVersion, languageVersionFile, root string) (interfaces.FrameworkVersionDetector, *plugin.Client) {
	///home/deqode/Documents/code-analyser/plugin/go/main

	////TODO: need to add semver check and return correct version
	beegoDetector, client := pluginClient.FrameworkPluginCall(exec.Command("sh", "-c", "go run plugin/go/beego/v1_x/main.go"))
	x := interfaces.FrameworkVersionDetector{
		Detector: beegoDetector,
	}
	return x, client
}
