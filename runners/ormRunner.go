package runners

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/protos/protos"
)

func OrmRunner(ormDetector interfaces.ORMDetector, runTimeVersion, languageVersionFile, root string) *protos.OrmOutput {
	ormsUsed := ormDetector.GetLibraryUsed(runTimeVersion, root)
	ormOutputs := protos.OrmOutput{}
	for ormUsed, OrmVersion := range *ormsUsed {
		usedOrm := ormDetector.GetDetector(OrmVersion, root, ormUsed)
		if ormFound := OrmDetectorRunner(usedOrm, languageVersionFile, runTimeVersion, root); ormFound != nil {
			ormOutputs.Orms = append(ormOutputs.Orms, ormFound)
			ormOutputs.Used = true
		}
	}
	return &ormOutputs
}

func OrmDetectorRunner(orm interfaces.ORM, languageVersionFile, runTimeVersion, root string) *protos.ORM {
	versionedOrm := orm.GetVersionDetector(runTimeVersion, languageVersionFile, root)
	versionDetector := versionedOrm.Detector

	if _, err := versionDetector.IsORMFound(runTimeVersion, root); err == nil {
		if _, err := versionDetector.IsORMUsed(runTimeVersion, root); err == nil {
			if detected, err := versionDetector.Detect(runTimeVersion, root); err == nil && detected {
				version, _ := versionDetector.GetVersionName()
				name, _ := versionDetector.GetORMName()
				return &protos.ORM{
					Name:    name,
					Version: version,
				}
			}
		}
	}
	return nil
}
