package interfaces

import "code-analyser/pluginClient/pb"

type ORMVersion interface {
	Detect(*pb.ServiceInput) (*pb.ServiceOutputDetectOrm,error)  // deep level detection
	IsORMUsed(*pb.ServiceInput) (*pb.ServiceOutputBool,error)
	PercentOfORMUsed(*pb.ServiceInput) (*pb.ServiceOutputFloat,error)
}
