package beego

import (
	protos "code-analyser/protos/pb"
	"context"
)

func Beego_v_1_6(context.Context, string, string) (*protos.FrameworkOutput, error) {
	return &protos.FrameworkOutput{
		Name:    "beego",
		Version: "1.6",
		Used:    true,
	}, nil
}


