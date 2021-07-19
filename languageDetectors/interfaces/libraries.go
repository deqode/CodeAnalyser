package interfaces

import (
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
)

//Library plugin methods
type Library interface {
	Detect(input *helpers.Input) (*plugin.LibraryType, error) //todo:
	IsUsed(*helpers.Input) (*helpers.BoolOutput, error)
	PercentOfUsed(*helpers.Input) (*helpers.FloatOutput, error)
}
