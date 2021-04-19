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

func ParseLibraryFromDependencies(dependenciesList map[string]string, langYamlObject *versionsPB.LanguageVersion) map[string]DependencyDetail {
	library := map[string]DependencyDetail{}
	for key, supportedLibrary := range langYamlObject.Libraries {
		if versionUsed, ok := dependenciesList[key]; ok {
			for versionName, v := range supportedLibrary.Version {
				if helpers.SemverValidateFromArray(v.Semver, versionUsed) {
					library[key] = DependencyDetail{
						Version: versionName,
						Command: v.Plugincommand,
					}
				}
			}
		}
	}
	return library
}

func LibraryRunner(libraryList map[string]DependencyDetail, runtimeVersion, root string) []*languageSpecificPB.LibraryOutput {
	var libraryOutputs []*languageSpecificPB.LibraryOutput
	for libraryUsed, libraryDetails := range libraryList {
		isUsed := LibraryDetectorRunner(libraryUsed, libraryDetails, runtimeVersion, root)
		if isUsed != nil {
			libraryOutputs = append(libraryOutputs, isUsed)
		}
	}
	return libraryOutputs
}

//FrameworkDetectorRunner will find and run version detector & returns protos.FrameworkOutput to
func LibraryDetectorRunner(name string, libraryDetails DependencyDetail, runtimeVersion, root string) *languageSpecificPB.LibraryOutput {
	libraryResponse, client := pluginClient.LibraryPluginCall(exec.Command("sh", "-c", libraryDetails.Command))
	for client.Exited() {
		client.Kill()
	}
	isUsed, err := libraryResponse.IsUsed(&pb.ServiceInput{
		RuntimeVersion: runtimeVersion,
		Root:           root,
	})
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if isUsed.Value {
		detection, err := libraryResponse.Detect(&pb.ServiceInput{
			RuntimeVersion: runtimeVersion,
			Root:           root,
		})
		if err != nil {
			utils.Logger(err)
			return nil
		}
		if detection.Value {
			percentUsed, err := libraryResponse.PercentOfUsed(&pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           root,
			})
			if err != nil {
				utils.Logger(err)
				return nil
			}
			return &languageSpecificPB.LibraryOutput{
				Type:           languageSpecificPB.LibraryOutput_Type(detection.Type),
				Used:           isUsed.Value,
				Name:           name,
				Version:        libraryDetails.Version,
				PercentageUsed: percentUsed.Value,
			}
		}
	}
	return nil
}
