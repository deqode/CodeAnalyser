package interfaces

import (
	"code-analyser/protos/pb/helpers"
)

//Framework It defines methods of framework
type Framework interface {
	Detect(input *helpers.Input) (*helpers.BoolOutput, error) //todo: can return FrameworkOutput ?
	IsUsed(input *helpers.Input) (*helpers.BoolOutput, error)
	PercentOfUsed(input *helpers.Input) (*helpers.FloatOutput, error)
}
