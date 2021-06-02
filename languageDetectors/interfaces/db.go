package interfaces

import (
	"code-analyser/protos/pb/plugin"
)

//DbVersion It contains methods for Db details
type DbVersion interface {
	Detect(*plugin.ServiceInput) (*plugin.ServiceOutputBoolInt, error) // deep level detection
	IsDbUsed(input *plugin.ServiceInput) (*plugin.ServiceOutputBool, error)
	PercentOfDbUsed(*plugin.ServiceInput) (*plugin.ServiceOutputFloat, error)
}
