package global

import "code-analyser/detector"

func DetectMakefile(root string) []string {
	var output []string
	exists, _, _ := detector.CheckExist(root + "/Makefile")
	if exists {
		output = append(output, "Makefile")
	}
	exists, _, _ = detector.CheckExist(root + "/*/Makefile")
	if exists {
		output = append(output, "Makefile")
	}
	exists, _, _ = detector.CheckExist(root + "/makefile")
	if exists {
		output = append(output, "makefile")
	}
	exists, _, _ = detector.CheckExist(root + "/*/makefile")
	if exists {
		output = append(output, "makefile")
	}
	return output
}
