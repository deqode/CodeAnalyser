package V_2_X

import "code-analyser/pluginClient/pb"

type Gin_V_2_X struct {
}

func (b *Gin_V_2_X) GetVersionName() (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Error: nil,
		Value: "2.x",
	}, nil
}

func (b *Gin_V_2_X) GetSemver() (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Error: nil,
		Value: ">=2.x",
	}, nil
}

func (b *Gin_V_2_X) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Error: nil,
		Value: true,
	}, nil
}

func (b *Gin_V_2_X) IsFrameworkFound(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Error: nil,
		Value: true,
	}, nil
}

func (b *Gin_V_2_X) IsFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Error: nil,
		Value: true,
	}, nil
}

func (b *Gin_V_2_X) PercentOfFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputFloat, error) {
	return &pb.ServiceOutputFloat{
		Error: nil,
		Value: 88,
	}, nil
}

func (b *Gin_V_2_X) GetFrameworkName() (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Error: nil,
		Value: "gin",
	}, nil
}
