package interfaces

type FrameworkDetector interface {
	GetLibrariesUsed(runtimeVersion, root string) *map[string]string
	GetDetector(libraryVersion, root, libraryUsed string) Framework
}

type Framework interface {
	GetVersionedDetector(runtimeVersion, languageVersionFile, root string) FrameworkVersionDetector
}

type FrameworkVersions interface {
	GetVersionName() string
	GetSemver() string
	Detect(runtimeVersion, root string) bool //todo: can return FrameworkOutput ?
	IsFrameworkFound(runtimeVersion, root string) bool
	IsFrameworkUsed(runtimeVersion, root string) bool
	PercentOfFrameworkUsed(runtimeVersion, root string) float64
	GetFrameworkName() string
}

type FrameworkVersionDetector struct {
	Default  bool
	Name     string
	Semver   string
	Detector FrameworkVersions
}
