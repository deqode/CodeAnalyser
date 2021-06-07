package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
)

//ParseFrameworkFromDependencies It will filter out frameworks from dependency list
func ParseFrameworkFromDependencies(dependenciesList map[string]string, langYamlObject *versionsPB.LanguageVersion) map[string]DependencyDetail {
	framework := map[string]DependencyDetail{}
	for key, supportedFramework := range langYamlObject.Framework {
		for frameworkVersion, dependencyVersionDetails := range supportedFramework.Version {
			for _, library := range dependencyVersionDetails.Libraries {
				if libraryUsedVersion, ok := dependenciesList[library.Name]; ok {
					if helpers.SemverValidateFromArray(library.Semver, libraryUsedVersion) {
						framework[key] = DependencyDetail{
							Version: frameworkVersion,
							Command: dependencyVersionDetails.Plugincommand,
						}
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
	frameworkResponse, client := pluginClient.CreateFrameworkClient(utils.CallPluginCommand(framworkDetails.Command))
	for client.Exited() {
		client.Kill()
	}
	isUsed, err := frameworkResponse.IsUsed(&pb.Input{
		RuntimeVersion: runtimeVersion,
		RootPath:           root,
	})
	if err != nil || isUsed.Error != nil {
		utils.Logger(err, isUsed.Error)
		return nil
	}
	if isUsed.Value {
		detection, err := frameworkResponse.Detect(&pb.Input{
			RuntimeVersion: runtimeVersion,
			RootPath:           root,
		})
		if err != nil || detection.Error != nil {
			utils.Logger(err, detection.Error)
			return nil
		}
		if detection.Value {
			percentageUsed, err := frameworkResponse.PercentOfUsed(&pb.Input{
				RuntimeVersion: runtimeVersion,
				RootPath:           root,
			})
			if err != nil || detection.Error != nil {
				utils.Logger(err, detection.Error)
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
