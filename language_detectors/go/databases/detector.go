package databases

import (
	"code-analyser/language_detectors/go/databases/postgres"
	"code-analyser/language_detectors/interfaces"
	"code-analyser/utils"
)

var databases = map[string]interfaces.Db{
	"postgres":&postgres.PostgresDb{},
}

type GODbDetector struct {
}

func (receiver *GODbDetector) GetDetector(libraryVersion, root, dbUsed string) interfaces.Db{
	utils.Logger("dbUsed",dbUsed)
	return databases[dbUsed]
}

func (receiver *GODbDetector) GetLibraryUsed(runtimeVersion, root string) *map[string]string {
	return &map[string]string{
		"postgres": "1.5",
	}
}

//func getDetector(root, version string, languageVersion *protos.LanguageVersion) []*versions.DBVersionDetector {
//	var output []*versions.DBVersionDetector
//	mod, _ := go_mod.ParseGoMod(root + "/go.mod")
//	libraryFound := mod.Required
//	library := libraryFound[0]
//	fwDetector := detector_helpers.FindDBDetector(databases[library.Name], languageVersion.Databases)
//	output = append(output, fwDetector)
//	//return Detector
//	return output
//}
//func DetectDB(version, root string, languageVersion *protos.LanguageVersion) *protos.DBOutput {
//	var output = protos.DBOutput{}
//	detector := getDetector(root, version, languageVersion)
//	output.Used = true
//	fwDetected, _ := detector[0].Detector(nil, version, root)
//	output.Databases = append(output.Databases, fwDetected)
//	return &output
//}
