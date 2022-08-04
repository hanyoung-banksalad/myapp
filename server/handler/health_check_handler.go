package handler

import (
	"context"

	myapp "github.com/hanyoung-banksalad/myapp/idl"
)

type HealthCheckHandlerFunc func(ctx context.Context, req *myapp.HealthCheckRequest) (*myapp.HealthCheckResponse, error)

func HealthCheck() HealthCheckHandlerFunc {
	return func(ctx context.Context, req *myapp.HealthCheckRequest) (*myapp.HealthCheckResponse, error) {
		return &myapp.HealthCheckResponse{}, nil
	}
}
