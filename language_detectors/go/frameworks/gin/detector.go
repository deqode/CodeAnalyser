package gin

import (
	"code-analyser/language_detectors/go/frameworks/gin/V_2_X"
	"code-analyser/language_detectors/interfaces"
	"code-analyser/language_detectors/go/frameworks/gin/V_1_X"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

var ginVersions = []*interfaces.FrameworkVersionDetector{
	{
		Default:true,
		Name:"v1.x",
		Semver: "",
		Detector:&V_1_X.Gin_V_1_X{},
	},
	{
		Default:true,
		Name:"v2.x",
		Semver: "",
		Detector:&V_2_X.Gin_V_2_X{},
	},
}

type GinFramework struct {
	Version []*interfaces.FrameworkVersionDetector
}

func (receiver *GinFramework) GetVersionedDetector(runtimeVersion, languageVersionFile, root string) (interfaces.FrameworkVersionDetector,*plugin.Client) {
	utils.Logger("GetVersionedDetector function called")
	return *ginVersions[1],nil
}