package parsers

import (
	"code-analyser/detector"
	"log"
	"strings"
)

type GoMod struct {
	ModuleName       string
	GoVersion        string
	required         []Packages
	requiredIndirect []Packages
}

type Packages struct {
	URI    string
	Vesion string
}

func ParseGoMod(modFile string) (*GoMod, error) {
	file, err := detector.ReadFile(modFile)
	if err != nil {
		return nil, err
	}
	log.Println(file)
	lines := strings.Split(file, "\n")
	for _, line := range lines {
		if detector.FindInLine("module ", line) {
			log.Println("Module Name:- ", strings.Split(line, "module ")[1])
		} else if detector.FindInLine("go ", line) {
			log.Println("go version:- ", strings.Split(line, "go ")[1])
		} else {
			log.Println(line)

		}
	}
	return nil, err
}
