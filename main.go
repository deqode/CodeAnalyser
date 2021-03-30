package main

import (
	"code-analyser/analyser"
	"code-analyser/language_detectors"
	"code-analyser/runners"
	"log"
)

func main() {
	path := "./"
	Scrape(path)
}

func Scrape(path string) {
	languages, _, _ := analyser.Analyse(path)
	log.Println(languages)
	for _, language := range languages {
		if languageDetector, ok := language_detectors.MapLanguageToDetectors[language.Name]; ok {
			runtimeVersion, _ := languageDetector.DetectRuntime(nil, path)
			runners.RunDetectors(languageDetector, runtimeVersion, path)
		}
	}
}
