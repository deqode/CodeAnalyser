package runners

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/protos/protos"
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
	if _, err := versionDetector.IsDbFound(runTimeVersion, root); err == nil {
		if _, err := versionDetector.IsDbUsed(runTimeVersion, root); err == nil {
			if detected, port, err := versionDetector.Detect(runTimeVersion, root); detected && err == nil {
				version, _ := versionDetector.GetVersionName()
				name, _ := versionDetector.GetVersionName()
				return &protos.DB{
					Version: version,
					Name:    name,
					Port:    port,
				}
			}
		}
	}

	return nil
}
