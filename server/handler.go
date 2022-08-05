package server

import (
	"context"

	myapp "github.com/hanyoung-banksalad/myapp/idl"
	"github.com/hanyoung-banksalad/myapp/server/handler"
)

// verify MyappServer implements all interface methods
var _ myapp.MyappServer = (*MyappServer)(nil)

func (s *MyappServer) HealthCheck(ctx context.Context, req *myapp.HealthCheckRequest) (*myapp.HealthCheckResponse, error) {
	return handler.HealthCheck()(ctx, req)
}

func (s *MyappServer) GetImage(ctx context.Context, req *myapp.GetImageRequest) (*currency.GetImageResponse, error) {
	return handler.GetImage()(ctx, req)
}
