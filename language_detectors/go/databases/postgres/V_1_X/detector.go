package V_1_X

import "code-analyser/pluginClient/pb"

type Postgres_V_1_X struct {
}

func (receiver *Postgres_V_1_X) GetSemver() (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Error: nil,
		Value: ">=2.x",
	}, nil
}

func (receiver *Postgres_V_1_X) IsDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Value: true,
		Error: nil,
	}, nil
}

func (receiver *Postgres_V_1_X) PercentOfDbUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	return &pb.ServiceOutputFloat{
		Error: nil,
		Value: 88,
	}, nil
}

// it returns port as well
func (receiver *Postgres_V_1_X) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBoolInt, error) {
	return &pb.ServiceOutputBoolInt{
		Value:    true,
		IntValue: 6565,
	}, nil
}
