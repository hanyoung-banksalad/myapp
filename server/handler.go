package server

import (
	"context"

	"github.com/hanyoung-banksalad/myapp/idl"
	"github.com/hanyoung-banksalad/myapp/server/handler"
)

// verify MyappServer implements all interface methods
var _ myapp.MyappServer = (*MyappServer)(nil)

func (s *MyappServer) HealthCheck(ctx context.Context, req *myapp.HealthCheckRequest) (*myapp.HealthCheckResponse, error) {
	return handler.HealthCheck()(ctx, req)
}
