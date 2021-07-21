package interfaces

import (
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
)

//Orm plugin methods
type Orm interface {
	Detect(input *helpers.Input) (*plugin.OrmOutput, error) // deep level detection
	IsUsed(input *helpers.Input) (*helpers.BoolOutput, error)
	PercentOfUsed(input *helpers.Input) (*helpers.FloatOutput, error)
}
