package V_1_X

import "code-analyser/pluginClient/pb"

type Beego_V_1_X struct {
}

func (b *Beego_V_1_X) GetVersionName() (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Error: nil,
		Value: "1.x",
	}, nil
}

func (b Beego_V_1_X) GetSemver() (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Error: nil,
		Value: ">=2.x",
	}, nil
}

func (b *Beego_V_1_X) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Error: nil,
		Value: true,
	}, nil
}

func (b *Beego_V_1_X) IsFrameworkFound(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Error: nil,
		Value: true,
	}, nil
}

func (b *Beego_V_1_X) IsFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Error: nil,
		Value: true,
	}, nil
}

func (b *Beego_V_1_X) PercentOfFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	return &pb.ServiceOutputFloat{
		Error: nil,
		Value: 88,
	}, nil
}

func (b *Beego_V_1_X) GetFrameworkName() (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Error: nil,
		Value: "beego",
	}, nil
}
