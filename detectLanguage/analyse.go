package detectlanguage

import (
	"github.com/go-cmd/cmd"
	"strconv"
	"strings"
)

// File is a struct of language name and percentage used
type File struct {
	Name    string
	Percent float64
}

//must be correct path of enry binary
var enryLoc = "./static/"

//GetLanguagesWithPercent analyse any projects with given path
// returns languages like(go,js etc ... ) with their percentages (Go:50%, JS:30%)
// on error returns slice of error(string),error
func GetLanguagesWithPercent(path string) ([]File, []string, error) {
	files := []File{}
	err := []string{}
	//must be correct name of enry binary
	detectCmd := cmd.NewCmd(enryLoc+"enry", path)
	finalStatus := <-detectCmd.Start()

	//if error from shell has length>0
	if len(finalStatus.Stderr) > 0 {
		err = finalStatus.Stderr
		return nil, err, nil
	}
	if finalStatus.Error != nil {
		err = finalStatus.Stderr
		return nil, nil, finalStatus.Error
	}

	for _, item := range finalStatus.Stdout {
		parsed, err := strconv.ParseFloat(strings.TrimSpace(strings.Split(item, "%")[0]), 2)
		if err != nil {
			return nil, nil, err
		}
		files = append(files, File{Name: strings.TrimSpace(strings.Split(item, "%")[1]), Percent: parsed})
	}
	return files, nil, nil
}
