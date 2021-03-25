package interfaces

type ORMVersionDetector struct {
	Default bool
	Name string
	Semver string
	Detector ORMVersion
}

type ORMVersion interface {
	GetVersionName() (string,error)
	GetSemver() (string,error)
	Detect(runtimeVersion ,root string) (bool,error)  // deep level detection
	IsORMUsed(runtimeVersion ,root string) (bool,error)
	IsORMFound(runtimeVersion ,root string) (bool,error)
	GetORMName() (string,error)
	PercentOfORMUsed(runtimeVersion ,root string) (float64,error)
}

type ORM interface {
	GetVersionDetector(runtimeVersion,ormVersionFile, root string) ORMVersionDetector
}

type ORMDetector interface {
	GetLibraryUsed(runtimeVersion ,root string) *map[string]string
	GetDetector(ORMVersion,root,libraryUsed string) ORM
}