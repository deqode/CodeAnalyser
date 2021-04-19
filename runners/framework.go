package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"os/exec"
)

//ParseFrameworkFromDependencies It will filter out frameworks from dependency list
func ParseFrameworkFromDependencies(dependenciesList map[string]string, langYamlObject *versionsPB.LanguageVersion) map[string]DependencyDetail {
	framework := map[string]DependencyDetail{}
	for key, supportedFramework := range langYamlObject.Framework {
		if versionUsed, ok := dependenciesList[key]; ok {
			for versionName, v := range supportedFramework.Version {
				if helpers.SemverValidateFromArray(v.Semver, versionUsed) {
					framework[key] = DependencyDetail{
						Version: versionName,
						Command: v.Plugincommand,
					}
				}
			}
		}
	}
	return framework
}

//FrameworkRunner will run to find Frameworks & returns its detectors.
//Must be called in LanguageSpecificDetector.DetectFrameworks
func FrameworkRunner(frameworkList map[string]DependencyDetail, runtimeVersion, root string) []*languageSpecificPB.FrameworkOutput {
	var frameworkOutputs []*languageSpecificPB.FrameworkOutput
	for frameworkUsed, frameworkDetails := range frameworkList {
		isUsed := FrameworkDetectorRunner(frameworkUsed, frameworkDetails, runtimeVersion, root)
		if isUsed != nil {
			frameworkOutputs = append(frameworkOutputs, isUsed)
		}
	}
	return frameworkOutputs
}

//FrameworkDetectorRunner will find and run version detector & returns protos.FrameworkOutput to
func FrameworkDetectorRunner(name string, framworkDetails DependencyDetail, runtimeVersion, root string) *languageSpecificPB.FrameworkOutput {
	frameworkResponse, client := pluginClient.FrameworkPluginCall(exec.Command("sh", "-c", framworkDetails.Command))
	for client.Exited() {
		client.Kill()
	}
	isUsed, err := frameworkResponse.IsFrameworkUsed(&pb.ServiceInput{
		RuntimeVersion: runtimeVersion,
		Root:           root,
	})
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if isUsed.Value {
		detection, err := frameworkResponse.Detect(&pb.ServiceInput{
			RuntimeVersion: runtimeVersion,
			Root:           root,
		})
		if err != nil {
			utils.Logger(err)
			return nil
		}
		if detection.Value {
			percentageUsed, err := frameworkResponse.PercentOfFrameworkUsed(&pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           root,
			})
			if err != nil {
				utils.Logger(err)
				return nil
			}
			return &languageSpecificPB.FrameworkOutput{
				Used:           isUsed.Value,
				Name:           name,
				Version:        framworkDetails.Version,
				PercentageUsed: percentageUsed.Value,
			}
		}
	}
	return nil
}
