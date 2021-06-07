package interfaces

import (
	"code-analyser/protos/pb/plugin"
)

type Library interface {
	Detect(input *plugin.Input) (*plugin.LibraryType, error) //todo:
	IsUsed(*plugin.Input) (*plugin.BoolOutput, error)
	PercentOfUsed(*plugin.Input) (*plugin.FloatOutput, error)
}
