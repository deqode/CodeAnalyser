package interfaces

import (
	"code-analyser/protos/pb/plugin"
)

//ORMVersion It contains methods to get details of ORM
type ORMVersion interface {
	Detect(*plugin.ServiceInput) (*plugin.ServiceOutputDetectOrm, error) // deep level detection
	IsORMUsed(*plugin.ServiceInput) (*plugin.ServiceOutputBool, error)
	PercentOfORMUsed(*plugin.ServiceInput) (*plugin.ServiceOutputFloat, error)
}
