package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hanyoung-banksalad/myapp/idl"
)

func TestHealthCheck(t *testing.T) {
	ctx := context.Background()
	req := &myapp.HealthCheckRequest{}

	resp, err := HealthCheck()(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, &myapp.HealthCheckResponse{}, resp)
}
