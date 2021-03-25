package runners

import (
	"code-analyser/language_detectors/interfaces"
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
	versionDetector := versionedFramework.Detector
	vname, _ := versionDetector.GetVersionName()
	fname, _ := versionDetector.GetFrameworkName()
	if _, err := versionDetector.IsFrameworkFound(runtimeVersion, root); err == nil {
		if _, err := versionDetector.IsFrameworkUsed(runtimeVersion, root); err == nil {
			if _, err := versionDetector.Detect(runtimeVersion, root); err == nil {
				if client != nil {
					client.Kill()
				}
				return &protos.FrameworkOutput{
					Version: vname,
					Name:    fname,
					Used:    true,
				}
			}
		}
	}
	if client != nil {
		client.Kill()
	}
	return &protos.FrameworkOutput{
		Version: vname,
		Name:    fname,
		Used:    false,
	}
}
