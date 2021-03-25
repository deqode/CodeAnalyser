package interfaces

type FrameworkDetector interface {
	GetLibrariesUsed(runtimeVersion, root string) *map[string]string
	GetDetector(libraryVersion, root, libraryUsed string) Framework
}

type Framework interface {
	GetVersionedDetector(runtimeVersion, languageVersionFile, root string) FrameworkVersionDetector
}

type FrameworkVersions interface {
	GetVersionName() (string, error)
	GetSemver() (string, error)
	Detect(runtimeVersion, root string) (bool, error) //todo: can return FrameworkOutput ?
	IsFrameworkFound(runtimeVersion, root string) (bool, error)
	IsFrameworkUsed(runtimeVersion, root string) (bool, error)
	PercentOfFrameworkUsed(runtimeVersion, root string) (int64, error)
	GetFrameworkName() (string, error)
}

type FrameworkVersionDetector struct {
	Default  bool
	Name     string
	Semver   string
	Detector FrameworkVersions
}
