package interfaces

import "code-analyser/pluginClient/pb"

type Env interface {
	Detect(*pb.ServiceInput)(*pb.ServiceOutputEnv,error)
	IsUsed(*pb.ServiceInput)(*pb.ServiceOutputBool,error)
}