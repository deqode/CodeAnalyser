package interfaces

import (
	"code-analyser/pluginClient/pb"
)

type Library interface {
	Detect(*pb.ServiceInput) (*pb.ServiceOutputLibraryType, error) //todo:
	IsUsed(*pb.ServiceInput) (*pb.ServiceOutputBool, error)
	PercentOfUsed(*pb.ServiceInput) (*pb.ServiceOutputFloat, error)
}
