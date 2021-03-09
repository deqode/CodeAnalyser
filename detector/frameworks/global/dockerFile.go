package global

import "code-analyser/detector"

func DetectDocker(root string) []string {
	var output []string
	exists, _, _ := detector.CheckExist(root + "/Dockerfile")
	if exists {
		output = append(output, "Dockerfile")
	}
	exists, _, _ = detector.CheckExist(root + "/docker-compose.yml")
	if exists {
		output = append(output, "docker-compose.yml")
	}
	return output
}
