package global

import "code-analyser/detector"

func DetectProc(root string) []string {
	var output []string
	exists, _, _ := detector.CheckExist(root + "/Procfile")
	if exists {
		output = append(output, "Procfile")
	}
	exists, _, _ = detector.CheckExist(root + "/*/Procfile")
	if exists {
		output = append(output, "Procfile")
	}

	return output
}
