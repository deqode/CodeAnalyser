package runners

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/utils"
	"sync"
)

// RunDetectors will run all the language specific detectors & returns language wise output to decision maker
// Must be called independently on detection of language
func RunDetectors(detector interfaces.LanguageSpecificDetector, runtimeVersion, path string) {
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(2)
	go func() {
		fw, err := detector.DetectFrameworks(nil, runtimeVersion, path)
		if err != nil {
			utils.Logger(err)
		}
		utils.Logger(fw)
		wg.Done()
	}()
	go func() {
		fw, err := detector.DetectDBs(nil, runtimeVersion, path)
		if err != nil {
			utils.Logger(err)
		}
		utils.Logger(fw)
		wg.Done()
	}()
}