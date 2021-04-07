package interfaces

import "code-analyser/pluginClient/pb"

//ORMVersion It contains methods to get details of ORM
type ORMVersion interface {
	Detect(*pb.ServiceInput) (*pb.ServiceOutputDetectOrm, error) // deep level detection
	IsORMUsed(*pb.ServiceInput) (*pb.ServiceOutputBool, error)
	PercentOfORMUsed(*pb.ServiceInput) (*pb.ServiceOutputFloat, error)
}
