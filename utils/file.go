package utils

import (
	"os"
	"path/filepath"
	"strings"
)

//SearchFileInDirectory search specific file in directory , it will return path of files
func SearchFileInDirectory(fileName, dirPath string) ([]FileDetails, error) {
	var pluginYamlFiles []FileDetails
	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == fileName {
				dir := strings.Split(path, info.Name())[0]
				pluginYamlFiles = append(pluginYamlFiles, FileDetails{path: path, dir: dir})
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return pluginYamlFiles, err
}
