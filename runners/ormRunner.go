package runners

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/pb"
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

	serviceInput:=&pb.ServiceInput{
		RuntimeVersion: runTimeVersion,
		Root:           root,
	}
	if _, err := versionDetector.IsORMFound(serviceInput); err == nil {
		if _, err := versionDetector.IsORMUsed(serviceInput); err == nil {
			if detected, err := versionDetector.Detect(serviceInput); err == nil && detected.Value {
				version, _ := versionDetector.GetVersionName()
				name, _ := versionDetector.GetORMName()
				return &protos.ORM{
					Name:    name.Value,
					Version: version.Value,
				}
			}
		}
	}
	return nil
}
