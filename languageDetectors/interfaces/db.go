package interfaces

import (
	"code-analyser/protos/pb/plugin"
)

//Db It contains methods for Db details
type Db interface {
	Detect(input *plugin.Input) (*plugin.BoolIntOutput, error) // deep level detection
	IsUsed(input *plugin.Input) (*plugin.BoolOutput, error)
	PercentOfUsed(input *plugin.Input) (*plugin.FloatOutput, error)
}
