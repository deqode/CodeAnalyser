package interfaces

import (
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
)

// Env will be implemented for plugin
type Env interface {
	Detect(input *helpers.Input) (*languageSpecific.Envs, error)
}
