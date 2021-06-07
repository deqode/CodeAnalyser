package interfaces

import (
	"code-analyser/protos/pb/plugin"
)

// Env will be implemented for plugin
type Env interface {
	Detect(input *plugin.Input) (*plugin.EnvsOutput, error)
}
