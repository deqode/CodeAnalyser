package _go

import "code-analyser/detector"

func DetectBeego(root string) *detector.DetectedOutput {
	report := map[bool]int{true: 0, false: 0}

	res, _, _ := detector.CheckKeyInFile(root+"/go.mod", "beego")
	report[res]++

	exist, files, _ := detector.CheckExist(root + "/*.go")
	report[exist]++

	if exist {
		res, _ := detector.FindInFiles(files, "beego")
		report[res]++
	}
	exist, files, _ = detector.CheckExist(root + "/*/*.go")
	report[exist]++

	if exist {
		res, _ := detector.FindInFiles(files, "beego")
		report[res]++
	}

	exist, _, _ = detector.CheckKeyInDir(root, "Run")
	report[exist]++
	return &detector.DetectedOutput{
		FrameWorkUsed: report[true] > report[false],
	}
}
