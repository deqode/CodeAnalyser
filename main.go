package main

import (
	"code-analyser/analyser"
	"code-analyser/language_detectors"
	"code-analyser/runners"
)

func main() {
	path := "/home/deqode/Documents/basic_repo/beegoRepo"
	Scrape(path)
}

func Scrape(path string) {
	languages, _, _ := analyser.Analyse(path)
	for _, language := range languages {
		languageDetector := language_detectors.MapLanguageToDetectors[language.Name]
		runtimeVersion, _ := languageDetector.DetectRuntime(nil, path)
		runners.RunDetectors(languageDetector, runtimeVersion, path)
	}
}
