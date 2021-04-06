package interfaces

import (
	"code-analyser/pluginClient/pb"
)
//FrameworkVersions It defines methods of framework
type FrameworkVersions interface {
	Detect(*pb.ServiceInput) (*pb.ServiceOutputBool, error) //todo: can return FrameworkOutput ?
	IsFrameworkUsed(*pb.ServiceInput) (*pb.ServiceOutputBool, error)
	PercentOfFrameworkUsed(*pb.ServiceInput) (*pb.ServiceOutputFloat, error)
}
