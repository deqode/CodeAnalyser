package runners

//DbRunner will run to detect its dbs and return its detectors
//func DbRunner(dbDetector interfaces.DbVersion, runTimeVersion, languageVersionFile, root string) *protos.DBOutput {
//	dbsUsed := dbDetector.GetLibraryUsed(runTimeVersion, root)
//	dbOutputs := protos.DBOutput{}
//	for dbUsed, dbVersion := range *dbsUsed {
//		usedDb := dbDetector.GetDetector(dbVersion, root, dbUsed)
//		if dbFound := DbDetectorRunner(usedDb, languageVersionFile, runTimeVersion, root); dbFound != nil {
//			dbOutputs.Databases = append(dbOutputs.Databases, dbFound)
//			dbOutputs.Used = true
//		}
//	}
//	return &dbOutputs
//}
//
////TODO handle errors on every method calls
//func DbDetectorRunner(db interfaces.Db, languageVersionFile, runTimeVersion, root string) *protos.DB {
//	versionedDb,dbClient := db.GetVersionDetector(runTimeVersion, languageVersionFile, root)
//	versionDetector := versionedDb.Detector
//	serviceInput:=&pb.ServiceInput{
//		RuntimeVersion: runTimeVersion,
//		Root:           root,
//	}
//	if dbClient != nil {
//		defer dbClient.Kill()
//	}
//	if _, err := versionDetector.IsDbFound(serviceInput); err == nil {
//		if _, err := versionDetector.IsDbUsed(serviceInput); err == nil {
//			if detected, err := versionDetector.Detect(serviceInput); detected.Value && err == nil {
//				version, _ := versionDetector.GetVersionName()
//				name, _ := versionDetector.GetDbName()
//				return &protos.DB{
//					Version: version.Value,
//					Name:    name.Value,
//					Port:    detected.IntValue,
//				}
//			}
//		}
//	}
//
//	return nil
//}
