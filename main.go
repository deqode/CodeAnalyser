package main

import (
	"code-analyser/detector"
	"code-analyser/detector/analyser"
	"code-analyser/detector/frameworks"
	"code-analyser/detector/frameworks/global"
	"log"
)

const enryLoc = "/home/deqode/Documents/code-analyser/"

func main() {
	//root := "."

	//root := "/home/deqode/Documents/basic_repo/ginRepo"
	//root := "/home/deqode/Documents/basic_repo/beegoRepo"
	//root := "/home/deqode/Documents/go-basics/swapi"
	root := "/home/deqode/Documents/go-basics/gormBasics"

	languages, errarr, err := analyser.Analyse(root, enryLoc)
	if err != nil || len(errarr) > 0 {
		log.Panic(err, errarr)
	}
	for _, language := range languages {
		log.Println("Found", language.Name, language.Percent)
		if detectors, ok := frameworks.Languages[language.Name]; ok {
			detectors.Prerequisites(root)
			for _, framework := range detectors.Frameworks {
				detected := framework.Detector(root)
				log.Println(framework.Name, detected.FrameworkUsed)
			}
			for _, dbChecker := range detectors.DBCheckers {
				_, db, port := dbChecker.Detector(root)
				log.Println(dbChecker.Name, db, port)
			}
			for _, orm := range detectors.Orms {
				eorm := orm.Detector(root)
				log.Println(orm.Name, eorm)
			}
		}
	}
	log.Println(global.DetectDocker(root))
	log.Println(global.DetectProc(root))
	//log.Println(parsers.ParsePackageJSON("/home/deqode/Documents/basic_repo/expressRepo/package.json"))
	//parsers.ParseGoMod("/home/deqode/Documents/code-analyser/go.mod")
	log.Println(detector.ValidateVersion(">=1.2.3-beta.1", "4.17.1"))
}
