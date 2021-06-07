package interfaces

import (
	"code-analyser/protos/pb/plugin"
)

//ORMVersion It contains methods to get details of ORM
type ORMVersion interface {
	Detect(input *plugin.Input) (*plugin.OrmOutput, error) // deep level detection
	IsUsed(input *plugin.Input) (*plugin.BoolOutput, error)
	PercentOfUsed(input *plugin.Input) (*plugin.FloatOutput, error)
}
