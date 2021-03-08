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
	"strconv"
	"strings"
)

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

func CheckKeyInDir(dirPath, key string) (bool, map[string]int, error) {
	existed, _, _ := CheckExist(dirPath)
	var outPut = map[string]int{}

	if existed {
		cmd, err := exec.Command("grep", "-rnw", dirPath, "-e", key).Output()
		if err != nil {
			return false, outPut, err
		}
		if len(cmd) < 0 {
			return false, outPut, nil
		}
		lines := strings.Split(string(cmd), "\n")
		for i := 0; i < len(lines); i++ {
			if len(strings.Split(lines[i], ":")) > 1 {
				lno := strings.Split(lines[i], ":")[1]
				keyPath := strings.Split(lines[i], ":")[0]
				outPut[keyPath], _ = strconv.Atoi(lno)
			}
		}

		return true, outPut, nil
	} else {
		return false, outPut, nil

	}
}

func CheckKeyInFile(filepath, key string) (bool, map[string]int, error) {
	existed, _, _ := CheckExist(filepath)
	var outPut = map[string]int{}

	if existed {
		cmd, err := exec.Command("grep", "-rnw", filepath, "-e", key).Output()
		if err != nil {
			return false, outPut, err
		}
		if len(cmd) < 0 {
			return false, outPut, nil
		}
		lines := strings.Split(string(cmd), "\n")
		for i := 0; i < len(lines); i++ {
			if len(strings.Split(lines[i], ":")) > 1 {
				lno := strings.Split(lines[i], ":")[0]
				keyPath := strings.Split(lines[i], ":")[1]
				outPut[keyPath], _ = strconv.Atoi(lno)
			}
		}

		return true, outPut, nil
	} else {
		return false, outPut, nil

	}
}

func FindInFiles(files []string, key string) (bool, []map[string]int) {
	outPut := []map[string]int{}
	exists := false
	for _, file := range files {
		existed, res, err := CheckKeyInFile(file, key)
		if existed || err == nil {
			exists = true
			outPut = append(outPut, res)
		}
	}
	return exists, outPut
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

func GetLine(filePath string, lineno int) string {
	file, err := ReadFile(filePath)
	if err != nil {
		log.Println(err)
	}
	lines := strings.Split(file, "\n")
	if lineno-1 > len(lines) {
		return ""
	} else {
		return lines[lineno-1]
	}
}
