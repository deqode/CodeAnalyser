package runners

import (
	"code-analyser/language_detectors/interfaces"
	protos "code-analyser/protos/pb"
	"code-analyser/utils"
	"sync"
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
// RunDetectors will run all the language specific detectors & returns language wise output to decision maker
// Must be called independently on detection of language
func RunDetectors(detector interfaces.LanguageSpecificDetector, runtimeVersion, path string) {
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go func() {
		fw, err := detector.DetectFrameworks(nil, runtimeVersion, path)
		if err != nil {
			utils.Logger(err)
		}
		utils.Logger(fw)
		wg.Done()
	}()

}
