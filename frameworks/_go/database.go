package _go

import (
	"code-analyser/detector"
	"code-analyser/detector/protos"
	"strings"
)

func CheckDB(root string) []*protos.DBPort {
	var dbs []*protos.DBPort
	str, db, port := CheckPostgres(root)
	if db {
		x := protos.DBPort{Name: str, Port: port}
		dbs = append(dbs, &x)
	}
	return dbs
}

func CheckPostgres(root string) (string, bool, string) {
	report := map[bool]int{true: 0, false: 0}
	res, _, _ := detector.CheckKeyInFile(root+"/go.mod", "github.com/lib/pq")
	report[res]++

	exist, files, _ := detector.CheckExist(root + "/*.go")
	report[exist]++

	if exist {
		res, _ := detector.FindInFiles(files, "github.com/lib/pq")
		report[res]++
		res, _ = detector.FindInFiles(files, "postgres")
		report[res]++
	}
	exist, files, _ = detector.CheckExist(root + "/*/*.go")
	report[exist]++
	if exist {
		res, _ := detector.FindInFiles(files, "github.com/lib/pq")
		report[res]++
	}

	exists, filemap, err := detector.CheckKeyInDir(root, "port ")
	if err != nil {
		return "postgres", false, ""
	}
	var port string
	if exists {
		for file, lineo := range filemap {
			if len(strings.Split(detector.GetLine(file, lineo), "=")) > 1 {
				port = strings.TrimSpace(strings.Split(detector.GetLine(file, lineo), "=")[1])
			}
		}
	}
	return "postgres", report[true] > report[false], port
}
