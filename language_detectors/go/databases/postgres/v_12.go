package postgress

import (
	protos "code-analyser/protos/pb"
	"context"
)

func Postgress_v_12(context.Context, string, string) (*protos.DB, error) {
	return &protos.DB{
		Version: "12",
		Name:    "Postgres 12",
		Port:    8080,
	}, nil
}
