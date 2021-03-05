package analyser

import (
	"strconv"
	"strings"

	"github.com/go-cmd/cmd"
)

type File struct {
	Name    string
	Percent float64
}

// func main() {
// 	log.Println(Analyse("./"))

// }

//Analyse returns Files , slice of error, string, error
func Analyse(path, enryLoc string) ([]File, []string, error) {
	files := []File{}
	err := []string{}
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
