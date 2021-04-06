package interfaces

import "code-analyser/pluginClient/pb"

type DbVersion interface {
	Detect(*pb.ServiceInput) (*pb.ServiceOutputBoolInt, error) // deep level detection
	IsDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error)
	PercentOfDbUsed(*pb.ServiceInput) (*pb.ServiceOutputFloat, error)
}
