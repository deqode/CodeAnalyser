package main

import (
	"code-analyser/analyser"
	"code-analyser/frameworks"
	"log"
)

type AnalysedOutput struct {
	Languages []*Language
	Databases []*Database
}
type Language struct {
	Name       string
	FrameWorks []string
}
type Database struct {
	Names []string
}

const enryLoc = "/home/deqode/Documents/code-analyser/"

func main() {
	//root := "/home/deqode/Documents/basic_repo/ginRepo"
	root := "/home/deqode/Documents/basic_repo/beegoRepo"
	//root := "/home/deqode/Documents/basic_repo"

	languages, errarr, err := analyser.Analyse(root, enryLoc)
	if err != nil || len(errarr) > 0 {
		log.Panic(err, errarr)
	}
	for _, language := range languages {
		for _, framework := range frameworks.Languages[language.Name] {
			log.Println(framework.Name, framework.Detector(root).FrameWorkUsed)
		}
	}
}
