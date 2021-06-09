package framework

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
	"errors"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of Framework that talks over RPC.
type GRPCClient struct {
	Client plugin.FrameworkClient
}

//TODO add nil check in service input variables
//Detect will detect the usage of framework
func (m *GRPCClient) Detect(input *pb.Input) (*pb.BoolOutput, error) {
	if input != nil {
		ret, err := m.Client.Detect(context.Background(), &pb.Input{
			RuntimeVersion: input.RuntimeVersion,
			RootPath:       input.RootPath,
		})
		return ret, err
	}
	return nil, errors.New("service input nil found")
}

//IsUsed returns true if that framework used
func (m *GRPCClient) IsUsed(input *pb.Input) (*pb.BoolOutput, error) {
	ret, err := m.Client.IsUsed(context.Background(), &pb.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return ret, err
}

//PercentOfUsed returns percentage of framework used
func (m *GRPCClient) PercentOfUsed(input *pb.Input) (*pb.FloatOutput, error) {
	ret, err := m.Client.PercentOfUsed(context.Background(), &pb.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return ret, err
}

//GRPCServer  is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl interfaces.Framework
}

//Detect will detect the usage of framework
func (m *GRPCServer) Detect(ctx context.Context, req *pb.Input) (*pb.BoolOutput, error) {
	v, err := m.Impl.Detect(req)
	return &pb.BoolOutput{
		Value: v.Value,
		Error: v.Error,
	}, err
}

//IsFrameworkUsed returns true if that framework used
func (m *GRPCServer) IsUsed(ctx context.Context, req *pb.Input) (*pb.BoolOutput, error) {
	v, err := m.Impl.IsUsed(req)
	return &pb.BoolOutput{
		Value: v.Value,
		Error: v.Error,
	}, err
}

//PercentOfFrameworkUsed returns percentage of framework used
func (m *GRPCServer) PercentOfUsed(ctx context.Context, req *pb.Input) (*pb.FloatOutput, error) {
	v, err := m.Impl.PercentOfUsed(req)
	return &pb.FloatOutput{
		Value: v.Value,
		Error: v.Error,
	}, err
}
