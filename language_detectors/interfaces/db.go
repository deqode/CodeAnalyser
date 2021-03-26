package interfaces

import "code-analyser/pluginClient/pb";

type DbVersionDetector struct {
	Default bool
	Name string
	Semver string
	Detector DbVersion
}

type DbVersion interface {
	GetVersionName() (*pb.ServiceOutputString,error)
	GetSemver() (*pb.ServiceOutputString,error)
	Detect(*pb.ServiceInput) (*pb.ServiceOutputBoolInt,error)  // deep level detection
	IsDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool,error)
	IsDbFound(input *pb.ServiceInput) (*pb.ServiceOutputBool,error)
	GetDbName() (*pb.ServiceOutputString,error)
	PercentOfDbUsed(*pb.ServiceInput) (*pb.ServiceOutputFloat,error)
}

type Db interface {
	GetVersionDetector(runtimeVersion,dbVersionFile, root string) (DbVersionDetector,)
}

type DbDetector interface {
	GetLibraryUsed(runtimeVersion ,root string) *map[string]string
    GetDetector(dbVersion,root,libraryUsed string) Db
}
