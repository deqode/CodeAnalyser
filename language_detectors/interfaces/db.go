package interfaces


type DbVersionDetector struct {
	Default bool
	Name string
	Semver string
	Detector DbVersion
}

type DbVersion interface {
	GetVersionName() (string,error)
	GetSemver() (string,error)
	Detect(runtimeVersion ,root string) (bool,int64,error)  // deep level detection
	IsDbUsed(runtimeVersion ,root string) (bool,error)
	IsDbFound(runtimeVersion ,root string) (bool,error)
	GetDbName() (string,error)
	PercentOfDbUsed(runtimeVersion ,root string) (float64,error)
}

type Db interface {
	GetVersionDetector(runtimeVersion,dbVersionFile, root string) (DbVersionDetector,)
}

type DbDetector interface {
	GetLibraryUsed(runtimeVersion ,root string) *map[string]string
    GetDetector(dbVersion,root,libraryUsed string) Db
}
