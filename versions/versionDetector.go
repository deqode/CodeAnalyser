package versions

import (
	"code-analyser/language_detectors"
	protos "code-analyser/protos/pb"
	"context"
)

type FrameworkVersionDetector struct {
	Default bool
	Name string
	Semver string
	Detector func(context.Context, string, language_detectors.Dir) (*protos.FrameworkOutput, error)
}

type DBVersionDetector struct {
	Default bool
	Name string
	Semver string
	Detector func(context.Context, string, language_detectors.Dir) (*protos.FrameworkOutput, error)
}
type RuntimeVersion struct {
	Default bool
	Name string
	Semver string
	Detector func(context.Context, string, language_detectors.Dir) (*protos.FrameworkOutput, error)
}