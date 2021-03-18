package main

import (
	"code-analyser/analyser"
	"code-analyser/language_detectors"
	"code-analyser/language_detectors/go"
	"log"
	"sync"
)

var languageDetectors = map[string]language_detectors.LanguageSpecificDetector{
	GO:         _go.GoDetector{},
	JAVASCRIPT: nil,
	PYTHON:     nil,
}

const (
	GO         = "Go"
	PYTHON     = "Python"
	JAVASCRIPT = "JavaScript"
)

func main() {
	path := "/home/deqode/Documents/basic_repo/beegoRepo"
	languages, _, _ := analyser.Analyse(path)
	for _, language := range languages {
		detector := languageDetectors[language.Name]
		runtimeVersion, _ := detector.DetectRuntime(nil, path)
		runDetection(detector, runtimeVersion, path)
	}
}

func runDetection(detector language_detectors.LanguageSpecificDetector, runtimeVersion, path string) {
	var wg sync.WaitGroup
	defer wg.Wait()
	//
	//detector.RunPreDetect(nil, version, path)
	//wg.Add(1)
	//go func() {
	//	detector.ParseENVs(nil, path)
	//	wg.Done()
	//}()
	wg.Add(2)
	go func() {
		fw, err := detector.DetectFrameworks(nil, runtimeVersion, path)
		if err != nil {
			log.Println(err)
		}
		log.Println(fw)
		wg.Done()
	}()
	go func() {
		db, _ := detector.DetectDBs(nil, runtimeVersion, path)
		log.Println(db)
		wg.Done()
	}()
	//go func() {
	//	detector.DetectORMs(nil, version, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.DetectDependencies(nil, version, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.DetectLibraries(nil, version, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.GetStaticAssets(nil, version, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.GetStack(nil, version, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.DetectAppserver(nil, version, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.DetectBuildDirectory(nil, version, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.DetectTestCasesRunCommands(nil, version, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.DetectBuildCommands(nil, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.DetectStartUpCommands(nil, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.DetectSeedCommands(nil, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.DetectMigrationCommands(nil, path)
	//	wg.Done()
	//}()
	//go func() {
	//	detector.DetectAdHocScripts(nil, path)
	//	wg.Done()
	//}()

}
