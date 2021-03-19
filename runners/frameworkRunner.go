package runners

import (
	"code-analyser/language_detectors/interfaces"
	protos "code-analyser/protos/pb"
)

func FrameworkRunner(frameworkDetector interfaces.FrameworkDetector, runtimeVersion, languageVersionFile, root string, ) []*protos.FrameworkOutput {
	librariesUsed := frameworkDetector.GetLibrariesUsed(runtimeVersion, root)
	var frameworkOutputs []*protos.FrameworkOutput
	for libraryUsed, libraryVersion := range *librariesUsed {
		usedFramework := frameworkDetector.GetDetector(libraryVersion, root, libraryUsed)
		frameworkOutputs = append(frameworkOutputs, FrameworkDetectorRunner(usedFramework, languageVersionFile, runtimeVersion, root))
	}
	return frameworkOutputs
}

func FrameworkDetectorRunner(framework interfaces.Framework, languageVersionFile, runtimeVersion, root string) *protos.FrameworkOutput {
	versionedFramework := framework.GetVersionedDetector(runtimeVersion, languageVersionFile, root)
	versionDetector := versionedFramework.Detector
	if versionDetector.IsFrameworkFound(runtimeVersion, root) {
		if versionDetector.IsFrameworkUsed(runtimeVersion, root) {
			if versionDetector.Detect(runtimeVersion, root) {
				return &protos.FrameworkOutput{
					Version: versionDetector.GetVersionName(),
					Name:    versionDetector.GetFrameworkName(),
					Used:    true,
				}
			}
		}
	}
	return &protos.FrameworkOutput{
		Version: versionDetector.GetVersionName(),
		Name:    versionDetector.GetFrameworkName(),
		Used:    false,
	}
}
