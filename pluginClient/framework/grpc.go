package framework

import (
	"code-analyser/languageDetectors/interfaces"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/plugin"
	"errors"
	"golang.org/x/net/context"
)

//GRPCClient is an implementation of Db that talks over gRPC to GRPCServer
type GRPCClient struct {
	Client plugin.FrameworkClient
}

//TODO add nil check in service input variables

//Detect method call on client side over gRPC
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

//IsUsed method call on client side over gRPC
func (m *GRPCClient) IsUsed(input *pb.Input) (*pb.BoolOutput, error) {
	ret, err := m.Client.IsUsed(context.Background(), &pb.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return ret, err
}

//PercentOfUsed method call on client side over gRPC
func (m *GRPCClient) PercentOfUsed(input *pb.Input) (*pb.FloatOutput, error) {
	ret, err := m.Client.PercentOfUsed(context.Background(), &pb.Input{
		RuntimeVersion: input.RuntimeVersion,
		RootPath:       input.RootPath,
	})
	return ret, err
}

//GRPCServer is the gRPC server to communicate with GRPCClient
type GRPCServer struct {
	// This is the real implementation
	Impl interfaces.Framework
}

//Detect will detect if given framework found
func (m *GRPCServer) Detect(ctx context.Context, req *pb.Input) (*pb.BoolOutput, error) {
	v, err := m.Impl.Detect(req)
	return &pb.BoolOutput{
		Value: v.Value,
		Error: v.Error,
	}, err
}

//IsUsed will check is given framework used
func (m *GRPCServer) IsUsed(ctx context.Context, req *pb.Input) (*pb.BoolOutput, error) {
	v, err := m.Impl.IsUsed(req)
	return &pb.BoolOutput{
		Value: v.Value,
		Error: v.Error,
	}, err
}

//PercentOfUsed will check percent of given framework used
func (m *GRPCServer) PercentOfUsed(ctx context.Context, req *pb.Input) (*pb.FloatOutput, error) {
	v, err := m.Impl.PercentOfUsed(req)
	return &pb.FloatOutput{
		Value: v.Value,
		Error: v.Error,
	}, err
}
