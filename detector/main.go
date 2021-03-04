package detector

// Todo: we use go.mod parser https://github.com/uudashr/go-module/blob/master/parser.go

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)
//
//func main() {
//	//CheckKeyIn("/home/deqode/Documents/basic_repo", "gin")
//	log.Println(CheckExist("/home/deqode/Documents/alfred/*/*.go"))
//}

func CheckExist(path string) (bool, []string, error) {
	str, err := filepath.Glob(path)
	if len(str) > 0 {
		return true, str, err
	} else {
		return false, str, err
	}
}

func ListAll(root string) []string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func CheckKeyIn(filepath, key string) (map[string]int, error) {
	existed, _, _ := CheckExist(filepath)
	var outPut map[string]int

	if existed {

		cmd, err := exec.Command("grep", "-rnw", filepath, "-e", key).Output()
		if err != nil {
			return outPut, err
		}
		if len(cmd) < 0 {
			return outPut, nil
		}
		lines := strings.Split(string(cmd), "\n")
		for i := 0; i < len(lines); i++ {
			if len(strings.Split(lines[i], ":")) > 2 {
				lno := strings.Split(lines[i], ":")[1]
				keyPath := strings.Split(lines[i], ":")[0]
				log.Println(lno, keyPath)
			}
		}

		return outPut, nil
	} else {
		return outPut, nil

	}
}

func FindInLine(regex, line string) bool {
	matched, err := regexp.MatchString(regex, line)
	if err != nil {
		return false
	} else if matched {
		return true
	}
	return false
}

func ReadFile(path string) (string, error) {
	existed, _, _ := CheckExist(path)
	if existed {
		dat, err := ioutil.ReadFile(path)
		return string(dat), err
	} else {
		return "", errors.New("file not existed")
	}

}
