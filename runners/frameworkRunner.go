package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"code-analyser/utils"
	"os/exec"
)

func ParseFrameworkFromDependencies(dependenciesList map[string]string, langYamlObject *protos.LanguageVersion) map[string]*protos.PluginSemver {
	//framework
	framework := map[string]*protos.PluginSemver{}
	for _, supportedFramework := range langYamlObject.Framework {
		if versionUsed, ok := dependenciesList[supportedFramework.Name]; ok {
			for _, v := range supportedFramework.Versions {
				if helpers.SeverValidate(v.Semver, versionUsed) {
					framework[supportedFramework.Name] = v
				}
			}
		}
	}
	return framework
}

//FrameworkRunner will run to find Frameworks & returns its detectors.
// Must be called in LanguageSpecificDetector.DetectFrameworks
func FrameworkRunner(frameworkList map[string]*protos.PluginSemver, runtimeVersion, root string) []*protos.FrameworkOutput {
	var frameworkOutputs []*protos.FrameworkOutput
	for frameworkUsed, frameworkDetails := range frameworkList {
		isUsed := FrameworkDetectorRunner(frameworkUsed, frameworkDetails, runtimeVersion, root)
		if isUsed != nil {
			frameworkOutputs = append(frameworkOutputs, isUsed)
		}
	}
	return frameworkOutputs
}

//
// FrameworkDetectorRunner will find and run version detector & returns protos.FrameworkOutput to
func FrameworkDetectorRunner(name string, framworkDetails *protos.PluginSemver, runtimeVersion, root string) *protos.FrameworkOutput {
	frameworkResponse, client := pluginClient.FrameworkPluginCall(exec.Command("sh", "-c", "go run plugin/go/framework/beego/v1_x/main.go"))
	defer client.Kill()
	isUsed, err := frameworkResponse.IsFrameworkUsed(&pb.ServiceInput{
		RuntimeVersion: runtimeVersion,
		Root:           root,
	})
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if isUsed.Value == true {
		detection, err := frameworkResponse.Detect(&pb.ServiceInput{
			RuntimeVersion: runtimeVersion,
			Root:           root,
		})
		if err != nil {
			utils.Logger(err)
			return nil
		}
		if detection.Value == true {
			percentageUsed, err := frameworkResponse.PercentOfFrameworkUsed(&pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           root,
			})
			if err != nil {
				utils.Logger(err)
				return nil
			}
			return &protos.FrameworkOutput{
				Used:           false,
				Name:           name,
				Version:        framworkDetails.Name,
				PercentageUsed: percentageUsed.Value,
			}
		}
	}
	return nil
}
