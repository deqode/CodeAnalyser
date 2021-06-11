package interfaces

import (
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
)

//Orm It contains methods to get details of ORM
type Orm interface {
	Detect(input *helpers.Input) (*plugin.OrmOutput, error) // deep level detection
	IsUsed(input *helpers.Input) (*helpers.BoolOutput, error)
	PercentOfUsed(input *helpers.Input) (*helpers.FloatOutput, error)
}
