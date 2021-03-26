package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/framework"
	"code-analyser/pluginClient/pb"
	"github.com/hashicorp/go-plugin"
)

type Beego_V_1_x struct{}

func (b Beego_V_1_x) GetVersionName() (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Error: nil,
		Value: "1.x",
	}, nil
}

func (b Beego_V_1_x) GetSemver() (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Error: nil,
		Value: ">=2.x",
	}, nil
}

func (b Beego_V_1_x) Detect(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Error: nil,
		Value: true,
	}, nil
}

func (b Beego_V_1_x) IsFrameworkFound(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Error: nil,
		Value: true,
	}, nil
}

func (b Beego_V_1_x) IsFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputBool, error) {
	return &pb.ServiceOutputBool{
		Error: nil,
		Value: true,
	}, nil
}

func (b Beego_V_1_x) PercentOfFrameworkUsed(input *pb.ServiceInput) (*pb.ServiceOutputInt, error) {
	return &pb.ServiceOutputInt{
		Error: nil,
		Value: 88,
	}, nil
}

func (b Beego_V_1_x) GetFrameworkName() (*pb.ServiceOutputString, error) {
	return &pb.ServiceOutputString{
		Error: nil,
		Value: "beego",
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserFramework: &framework.FrameworkGRPCPlugin{Impl: &Beego_V_1_x{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
