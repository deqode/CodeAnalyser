package interfaces

import (
	"code-analyser/protos/pb/plugin"
)

//FrameworkVersions It defines methods of framework
type FrameworkVersions interface {
	Detect(*plugin.ServiceInput) (*plugin.ServiceOutputBool, error) //todo: can return FrameworkOutput ?
	IsFrameworkUsed(*plugin.ServiceInput) (*plugin.ServiceOutputBool, error)
	PercentOfFrameworkUsed(*plugin.ServiceInput) (*plugin.ServiceOutputFloat, error)
}
