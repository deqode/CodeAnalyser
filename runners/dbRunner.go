package runners

import (
	"code-analyser/language_detectors/interfaces"
	protos "code-analyser/protos/pb"
)

//DbRunner will run to detect its dbs and return its detectors
func DbRunner(dbDetector interfaces.DbDetector, runTimeVersion, languageVersionFile, root string) *protos.DBOutput {
	dbsUsed := dbDetector.GetLibraryUsed(runTimeVersion, root)
	dbOutputs := protos.DBOutput{}
	for dbUsed, dbVersion := range *dbsUsed {
		usedDb := dbDetector.GetDetector(dbVersion, root, dbUsed)
		if dbFound := DbDetectorRunner(usedDb, languageVersionFile, runTimeVersion, root); dbFound != nil {
			dbOutputs.Databases = append(dbOutputs.Databases, dbFound)
			dbOutputs.Used = true
		}
	}
	return &dbOutputs
}

func DbDetectorRunner(db interfaces.Db, languageVersionFile, runTimeVersion, root string) *protos.DB {
	versionedDb := db.GetVersionDetector(runTimeVersion, languageVersionFile, root)
	versionDetector := versionedDb.Detector
	if versionDetector.IsDbFound(runTimeVersion, root) {
		if versionDetector.IsDbUsed(runTimeVersion, root) {
			if detected, port := versionDetector.Detect(runTimeVersion, root); detected {
				return &protos.DB{
					Version: versionDetector.GetVersionName(),
					Name:    versionDetector.GetVersionName(),
					Port:    port,
				}
			}
		}
	}

	return nil
}
