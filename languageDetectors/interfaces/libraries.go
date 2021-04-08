package interfaces

import (
	"code-analyser/protos/pb/plugin"
)

type Library interface {
	Detect(*plugin.ServiceInput) (*plugin.ServiceOutputLibraryType, error) //todo:
	IsUsed(*plugin.ServiceInput) (*plugin.ServiceOutputBool, error)
	PercentOfUsed(*plugin.ServiceInput) (*plugin.ServiceOutputFloat, error)
}
