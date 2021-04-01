package V_1_X

import "code-analyser/pluginClient/pb"

type GORM_V_1_X struct {
}

func (reciever *GORM_V_1_X) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Value: true,
		Error: nil,
	}, nil
}

func (reciever *GORM_V_1_X) IsORMUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Value: true,
		Error: nil,
	}, nil
}

func (reciever *GORM_V_1_X) PercentOfORMUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	return &pb.ServiceOutputFloat{
		Value: 14.6,
		Error: nil,
	}, nil
}
