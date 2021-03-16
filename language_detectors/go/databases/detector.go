package databases

import (
	"code-analyser/detector_helpers"
	postgress "code-analyser/language_detectors/go/databases/postgres"
	"code-analyser/language_detectors/go/parsers/go_mod"
	protos "code-analyser/protos/pb"
	"code-analyser/versions"
)

var databases = map[string]map[string]*versions.DBVersionDetector{
	"postgres": postgress.PostgresVersions,
}

func getDetector(root, version string, languageVersion *protos.LanguageVersion) []*versions.DBVersionDetector {
	var output []*versions.DBVersionDetector
	mod, _ := go_mod.ParseGoMod(root + "/go.mod")
	libraryFound := mod.Required
	library := libraryFound[0]
	fwDetector := detector_helpers.FindDBDetector(databases[library.Name], languageVersion.Databases)
	output = append(output, fwDetector)
	//return Detector
	return output
}
func DetectDB(version, root string, languageVersion *protos.LanguageVersion) *protos.DBOutput {
	var output = protos.DBOutput{}
	detector := getDetector(root, version, languageVersion)
	output.Used = true
	fwDetected, _ := detector[0].Detector(nil, version, root)
	output.Databases = append(output.Databases, fwDetected)
	return &output
}
