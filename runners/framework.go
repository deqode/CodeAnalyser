package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"golang.org/x/net/context"
)

//ExtractFrameworksFromProjectDependencies It will filter out frameworks from dependency list
func ExtractFrameworksFromProjectDependencies(ctx context.Context, projectDependencies map[string]string, frameworkPlugins map[string]*versionsPB.DependencyDetails) map[string]DependencyDetail {
	framework := map[string]DependencyDetail{}
	for name, details := range frameworkPlugins {
		for frameworkVersion, versionDetails := range details.Version {
			for _, library := range versionDetails.Libraries {
				if usedLibraryVersion, ok := projectDependencies[library.Name]; ok {
					if helpers.SemverValidateFromArray(library.Semver, usedLibraryVersion) {
						framework[name] = DependencyDetail{
							Version: frameworkVersion,
							Command: versionDetails.PluginPath,
						}
					}
				}
			}
		}
	}
	return framework
}

//ExecuteFrameworkPlugins will run to find Frameworks & returns its detectors.
//Must be called in LanguageSpecificDetector.DetectFrameworks
func ExecuteFrameworkPlugins(ctx context.Context, frameworkPlugins map[string]DependencyDetail, runtimeVersion, projectRootPath string) []*languageSpecificPB.FrameworkOutput {
	var frameworkOutput []*languageSpecificPB.FrameworkOutput
	for name, details := range frameworkPlugins {
		frameworkPluginResponse := ExecuteFrameworkPlugin(ctx, name, details, runtimeVersion, projectRootPath)
		if frameworkPluginResponse != nil {
			frameworkOutput = append(frameworkOutput, frameworkPluginResponse)
		}
	}
	return frameworkOutput
}

//ExecuteFrameworkPlugin will find and run version detector & returns protos.FrameworkOutput to
func ExecuteFrameworkPlugin(ctx context.Context, name string, frameworkPluginDetail DependencyDetail, runtimeVersion, projectRootPath string) *languageSpecificPB.FrameworkOutput {
	pluginCall, _ := pluginClient.CreateFrameworkClient(utils.CallPluginCommand(frameworkPluginDetail.Command))

	pluginInput := &pb.Input{
		RuntimeVersion: runtimeVersion,
		RootPath:       projectRootPath,
	}

	isUsed, err := pluginCall.IsUsed(pluginInput)
	if err != nil || isUsed.Error != nil {
		utils.Logger(err, isUsed)
		return nil
	}
	if isUsed.Value == false {
		return nil
	}

	response, err := pluginCall.Detect(pluginInput)
	if err != nil || response.Error != nil {
		utils.Logger(err, response)
		return nil
	}
	if response.Value {
		return &languageSpecificPB.FrameworkOutput{
			Used:    true,
			Name:    name,
			Version: frameworkPluginDetail.Version,
		}
	}
	return nil
}
