package handler

import (
	"context"

	"/Users/han058/go/myapp"
)

type HealthCheckHandlerFunc func(ctx context.Context, req *myapp.HealthCheckRequest) (*myapp.HealthCheckResponse, error)

func HealthCheck() HealthCheckHandlerFunc {
	return func(ctx context.Context, req *myapp.HealthCheckRequest) (*myapp.HealthCheckResponse, error) {
		return &myapp.HealthCheckResponse{}, nil
	}
}
