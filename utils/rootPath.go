package utils

import (
	"path"
	"path/filepath"
	"runtime"
)

//RootDirPath It gives root path of project
func RootDirPath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
