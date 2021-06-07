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

func ExtractLibraryFromProjectDependencies(ctx context.Context, projectDependencies map[string]string, libraryPlugins map[string]*versionsPB.DependencyDetails) map[string]DependencyDetail {
	library := map[string]DependencyDetail{}
	for name, details := range libraryPlugins {
		if usedLibraryVersion, ok := projectDependencies[name]; ok {
			for version, versionDetails := range details.Version {
				if helpers.SemverValidateFromArray(versionDetails.Semver, usedLibraryVersion) {
					library[name] = DependencyDetail{
						Version: version,
						Command: versionDetails.PluginPath,
					}
				}
			}
		}
	}
	return library
}

func ExecuteLibraryPlugins(ctx context.Context, libraryPlugins map[string]DependencyDetail, runtimeVersion, projectRootPath string) []*languageSpecificPB.LibraryOutput {
	var libraryOutput []*languageSpecificPB.LibraryOutput

	for name, details := range libraryPlugins {
		isUsed := ExecuteLibraryPlugin(ctx, name, details, runtimeVersion, projectRootPath)
		if isUsed != nil {
			libraryOutput = append(libraryOutput, isUsed)
		}
	}
	return libraryOutput
}

//ExecuteLibraryPlugin will find and run version detector & returns protos.FrameworkOutput to
func ExecuteLibraryPlugin(ctx context.Context, name string, libraryDetails DependencyDetail, runtimeVersion, projectRootPath string) *languageSpecificPB.LibraryOutput {
	pluginCall, _ := pluginClient.CreateLibraryClient(utils.CallPluginCommand(libraryDetails.Command))

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
		percentUsed, err := pluginCall.PercentOfUsed(pluginInput)
		if err != nil || percentUsed.Error != nil {
			utils.Logger(err, percentUsed)
			return nil
		}

		return &languageSpecificPB.LibraryOutput{
			Type:           languageSpecificPB.LibraryOutput_Type(response.Type),
			Used:           isUsed.Value,
			Name:           name,
			Version:        libraryDetails.Version,
			PercentageUsed: percentUsed.Value,
		}
	}
	return nil
}
