package myapp

import (
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/banksalad/go-banksalad"
	myapp "github.com/hanyoung-banksalad/myapp/idl"
)

//go:generate mockgen -package myapp -destination ./mock_client.go -mock_names MyappClient=MockMyappClient github.com/hanyoung-banksalad/myapp/idl MyappClient
const serviceConfig = `{"loadBalancingPolicy":"round_robin"}`

var (
	once sync.Once
	cli  myapp.MyappClient

	// verify MockMyappClient implements all MyappClient interface methods
	_ myapp.MyappClient = (*MockMyappClient)(nil)
)

func GetMyappClient(serviceHost, caller string) myapp.MyappClient {
	once.Do(func() {
		conn, _ := grpc.Dial(
			serviceHost,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(serviceConfig),
			grpc.WithUnaryInterceptor(banksalad.UnaryClientInterceptor(caller)),
		)

		cli = myapp.NewMyappClient(conn)
	})

	return cli
}
