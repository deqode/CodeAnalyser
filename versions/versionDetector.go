package versions

import (
	"code-analyser/protos/protos"
	"context"
)

type DBVersionDetector struct {
	Default  bool
	Name     string
	Semver   string
	Detector func(context.Context, string, string) (*protos.DB, error)
}
type ORMDetector struct {
	Default  bool
	Name     string
	Semver   string
	Detector func(context.Context, string, string) (*protos.OrmOutput, error)
}

type RuntimeVersion struct {
	Default  bool
	Name     string
	Semver   string
	Detector func(context.Context, string, string) (string, error)
}
