package runners

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
)

//FrameworkRunner will run to find Frameworks & returns its detectors.
// Must be called in LanguageSpecificDetector.DetectFrameworks
func FrameworkRunner(frameworkDetector interfaces.FrameworkDetector, runtimeVersion, languageVersionFile, root string) []*protos.FrameworkOutput {
	librariesUsed := frameworkDetector.GetLibrariesUsed(runtimeVersion, root)
	var frameworkOutputs []*protos.FrameworkOutput
	for libraryUsed, libraryVersion := range *librariesUsed {
		usedFramework := frameworkDetector.GetDetector(libraryVersion, root, libraryUsed)
		frameworkOutputs = append(frameworkOutputs, FrameworkDetectorRunner(usedFramework, languageVersionFile, runtimeVersion, root))
	}
	return frameworkOutputs
}

// FrameworkDetectorRunner will find and run version detector & returns protos.FrameworkOutput to
func FrameworkDetectorRunner(framework interfaces.Framework, languageVersionFile, runtimeVersion, root string) *protos.FrameworkOutput {
	versionedFramework, client := framework.GetVersionedDetector(runtimeVersion, languageVersionFile, root)
	if client != nil {
		defer client.Kill()
	}
	versionDetector := versionedFramework.Detector
	vname, _ := versionDetector.GetVersionName()
	fname, _ := versionDetector.GetFrameworkName()
	input := &pb.ServiceInput{Root: root, RuntimeVersion: runtimeVersion}
	if _, err := versionDetector.IsFrameworkFound(input); err == nil {
		if _, err := versionDetector.IsFrameworkUsed(input); err == nil {
			if _, err := versionDetector.Detect(input); err == nil {

				return &protos.FrameworkOutput{
					Version: vname.Value,
					Name:    fname.Value,
					Used:    true,
				}
			}
		}
	}

	return &protos.FrameworkOutput{
		Version: vname.Value,
		Name:    fname.Value,
		Used:    false,
	}
}
