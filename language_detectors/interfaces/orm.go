package interfaces

import "code-analyser/pluginClient/pb"

type ORMVersionDetector struct {
	Default bool
	Name string
	Semver string
	Detector ORMVersion
}

type ORMVersion interface {
	GetVersionName() (*pb.ServiceOutputString,error)
	GetSemver() (*pb.ServiceOutputString,error)
	Detect(*pb.ServiceInput) (*pb.ServiceOutputBool,error)  // deep level detection
	IsORMUsed(*pb.ServiceInput) (*pb.ServiceOutputBool,error)
	IsORMFound(*pb.ServiceInput) (*pb.ServiceOutputBool,error)
	GetORMName() (*pb.ServiceOutputString,error)
	PercentOfORMUsed(*pb.ServiceInput) (*pb.ServiceOutputFloat,error)
}

type ORM interface {
	GetVersionDetector(runtimeVersion,ormVersionFile, root string) ORMVersionDetector
}

type ORMDetector interface {
	GetLibraryUsed(runtimeVersion ,root string) *map[string]string
	GetDetector(ORMVersion,root,libraryUsed string) ORM
}