package interfaces

import (
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
)

//Db It contains methods for Db details
type Db interface {
	Detect(input *helpers.Input) (*plugin.BoolIntOutput, error) // deep level detection
	IsUsed(input *helpers.Input) (*helpers.BoolOutput, error)
	PercentOfUsed(input *helpers.Input) (*helpers.FloatOutput, error)
}
