package file

import (
	"errors"
	"path/filepath"
)

type Dir struct {
	path  string
	isDir bool
}

func (d *Dir) SetPath(path string) error {
	if validate(path) {
		d.path = path
		return nil
	}
	return errors.New("invalid path")
}
func (d *Dir) GetPath() string { return d.path }

func validate(path string) bool {
	if len(path) > 0 {
		str, _ := filepath.Glob(path)
		if len(str) > 0 {
			return true
		} else {
			return false
		}
	}
	return false
}
