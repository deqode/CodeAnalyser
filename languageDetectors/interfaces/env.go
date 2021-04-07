package interfaces

import "code-analyser/pluginClient/pb"
// Env will be implemented for plugin
type Env interface {
	Detect(*pb.ServiceInput)(*pb.ServiceOutputEnv,error)
}