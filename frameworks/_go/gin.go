package _go

import "code-analyser/detector"

func DetectGin(root string) *detector.DetectedOutput {
	report := map[bool]int{true: 0, false: 0}

	res, _, _ := detector.CheckKeyInFile(root+"/go.mod", "gin")
	report[res]++
	exist, files, _ := detector.CheckExist(root + "/*.go")
	report[exist]++

	if exist {
		res, _ := detector.FindInFiles(files, "gin")
		report[res]++
	}
	exist, files, _ = detector.CheckExist(root + "/*/*.go")
	report[exist]++

	if exist {
		res, _ := detector.FindInFiles(files, "gin")
		report[res]++
	}

	exist, _, _ = detector.CheckKeyInDir(root, "Run")
	report[exist]++
	return &detector.DetectedOutput{
		FrameworkUsed: report[true] > report[false],
	}
}
