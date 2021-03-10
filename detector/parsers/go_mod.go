package parsers

import (
	"code-analyser/detector"
	"log"
	"strings"
)

type GoMod struct {
	ModuleName string
	GoVersion  string
	required   []Package
}

type Package struct {
	URI      string
	Version  string
	Indirect bool
}

func ParseGoMod(modFile string) (*GoMod, error) {
	file, err := detector.ReadFile(modFile)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(file, "\n")
	output := GoMod{}
	for lno, line := range lines {
		if detector.FindInLine("go [0-9]", line) {
			output.GoVersion = strings.Split(line, "go ")[1]
		} else if detector.FindInLine("module ", line) {
			output.ModuleName = strings.Split(line, "module ")[1]
		} else if detector.FindInLine("require ", line) {
			requiresCount := lno
			if strings.Split(lines[requiresCount], "require ")[1] == "(" {
				for {
					requiresCount++
					if lines[requiresCount] == ")" {
						break
					}
					output.required = append(output.required, *parseURI(lines[requiresCount]))
				}
			} else {
				output.required = append(output.required, *parseURI(strings.Split(lines[requiresCount], "require ")[1]))
			}
		}
	}

	log.Println(output)
	return nil, nil
}

func parseURI(str string) *Package {
	out := Package{}
	if detector.FindInLine(" // indirect", str) {
		out.Indirect = true
	}
	s := strings.TrimSpace(strings.Split(str, " // indirect")[0])
	out.URI = strings.TrimSpace(strings.Split(s, " v")[0])
	out.Version = strings.TrimSpace(strings.Split(s, " v")[1])
	return &out
}
