package interfaces

import (
	"code-analyser/protos/pb/plugin"
)

//Framework It defines methods of framework
type Framework interface {
	Detect(input *plugin.Input) (*plugin.BoolOutput, error) //todo: can return FrameworkOutput ?
	IsUsed(input *plugin.Input) (*plugin.BoolOutput, error)
	PercentOfUsed(input *plugin.Input) (*plugin.FloatOutput, error)
}
