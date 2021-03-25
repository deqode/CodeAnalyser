package interfaces


type DbVersionDetector struct {
	Default bool
	Name string
	Semver string
	Detector DbVersion
}

type DbVersion interface {
	GetVersionName() string
	GetSemver() string
	Detect(runtimeVersion ,root string) (bool,int64)  // deep level detection
	IsDbUsed(runtimeVersion ,root string) bool
	IsDbFound(runtimeVersion ,root string) bool
	GetDbName() string
	PercentOfDbUsed(runtimeVersion ,root string) float64
}

type Db interface {
	GetVersionDetector(runtimeVersion,dbVersionFile, root string) DbVersionDetector
}

type DbDetector interface {
	GetLibraryUsed(runtimeVersion ,root string) *map[string]string
    GetDetector(dbVersion,root,libraryUsed string) Db
}
