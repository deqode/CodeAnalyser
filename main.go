package main

import (
	"github.com/spf13/afero"
	"log"
)

func main() {

	//var x language_detectors.LanguageSpecificDetector = &goDetector.GoDetector{}
	//file := afero.NewBasePathFs(new(afero.OsFs),"/home/deqode/Documents/basic_repo/beegoRepo")
	af := afero.NewOsFs()
	afero.ReadDir(af, "/home/deqode/Documents/basic_repo/beegoRepo")
	dir,_:=afero.ReadDir(af, "/home/deqode/Documents/basic_repo/")
	log.Println(dir[0].Sys())
	//log.Println(x.DetectRuntime(nil, file))
}
