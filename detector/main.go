package detector

import (
	"errors"
	"github.com/masterminds/semver"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Check file exists on that path can use *.go  or /*/*.go
func CheckExist(path string) (bool, []string, error) {
	str, err := filepath.Glob(path)
	if len(str) > 0 {
		return true, str, err
	} else {
		return false, str, err
	}
}

// Provide directory return list of all files
func ListAll(root string) []string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil
	}
	return files
}

// Perform deep search of keyword in dir, returns exists,map of filePath and line no & error
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

// Check Key in that file returns exists, Map of line code and line no & error
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

// This will search key in files path, returns existed,line no
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

// This will match regex of the provided string
func FindInLine(regex, line string) bool {
	matched, err := regexp.MatchString(regex, line)
	if err != nil {
		return false
	} else if matched {
		return true
	}
	return false
}

// This will return file content from file path
func ReadFile(path string) (string, error) {
	existed, _, _ := CheckExist(path)
	if existed {
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			return "", err

		} else {
			return string(dat), err
		}
	} else {
		return "", errors.New("file not existed")
	}
}

// This will return line content from line no of filePath
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

// This will compare constraint and version, returns bool,errors (if false),error in versioning
//https://pkg.go.dev/github.com/masterminds/semver
func ValidateVersion(constraint, version string) (bool, []error, error) {
	c, err := semver.NewConstraint(constraint)
	if err != nil {
		return false, nil, errors.New("invalid checker")
	}

	v, err := semver.NewVersion(version)
	if err != nil {
		return false, nil, errors.New("invalid version")
	}

	// Validate a version against a constraint.
	res, errs := c.Validate(v)
	return res, errs, nil

}
