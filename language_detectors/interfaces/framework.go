package interfaces

import (
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

type FrameworkDetector interface {
	GetLibrariesUsed(runtimeVersion, root string) *map[string]string
	GetDetector(libraryVersion, root, libraryUsed string) Framework
}

type Framework interface {
	GetVersionedDetector(runtimeVersion, languageVersionFile, root string) (FrameworkVersionDetector, *plugin.Client)
}

type FrameworkVersions interface {
	GetVersionName() (*pb.ServiceOutputString, error)
	GetSemver() (*pb.ServiceOutputString, error)
	Detect(*pb.ServiceInput) (*pb.ServiceOutputBool, error) //todo: can return FrameworkOutput ?
	IsFrameworkFound(*pb.ServiceInput) (*pb.ServiceOutputBool, error)
	IsFrameworkUsed(*pb.ServiceInput) (*pb.ServiceOutputBool, error)
	PercentOfFrameworkUsed(*pb.ServiceInput) (*pb.ServiceOutputFloat, error)
	GetFrameworkName() (*pb.ServiceOutputString, error)
}

type FrameworkVersionDetector struct {
	Default  bool
	Name     string
	Semver   string
	Detector FrameworkVersions
}
