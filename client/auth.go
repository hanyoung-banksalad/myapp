package client

//go:generate mockgen -package client -destination ./auth_mock.go -mock_names AuthClient=MockAuthClient github.com/banksalad/idl/gen/go/apis/v1/auth AuthClient

import (
	"sync"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/banksalad/go-banksalad"
	"github.com/banksalad/idl/gen/go/apis/v1/auth"
)

const authServiceConfig = `{"loadBalancingPolicy":"round_robin"}`

var (
	onceAuth sync.Once
	authCli  auth.AuthClient
)

func GetAuthClient(serviceHost, caller string) auth.AuthClient {
	onceAuth.Do(func() {
		conn, err := grpc.Dial(
			serviceHost,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(authServiceConfig),
			grpc.WithUnaryInterceptor(banksalad.UnaryClientInterceptor(caller)),
		)
		if err != nil {
			logrus.WithError(err).Panic("failed to get auth client")
		}

		authCli = auth.NewAuthClient(conn)
	})

	return authCli
}
