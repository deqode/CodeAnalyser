package _go

import (
	"code-analyser/detector"
)

func Gorm(root string) bool {
	report := map[bool]int{true: 0, false: 0}
	res, _, _ := detector.CheckKeyInFile(root+"/go.mod", "github.com/jinzhu/gorm")
	report[res]++

	exist, files, _ := detector.CheckExist(root + "/*.go")
	report[exist]++

	if exist {
		res, _ := detector.FindInFiles(files, "github.com/jinzhu/gorm")
		report[res]++
	}
	exist, files, _ = detector.CheckExist(root + "/*/*.go")
	report[exist]++
	if exist {
		res, _ := detector.FindInFiles(files, "github.com/jinzhu/gorm")
		report[res]++
	}
	exist, _, _ = detector.CheckKeyInDir(root, "gorm.Model")
	report[exist]++
	return report[true] > report[false]

}
