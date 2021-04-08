package interfaces

import (
	"code-analyser/protos/pb/plugin"
)
// Env will be implemented for plugin
type Env interface {
	Detect(*plugin.ServiceInput)(*plugin.ServiceOutputEnv,error)
}