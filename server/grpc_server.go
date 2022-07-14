package server

import (
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/banksalad/go-banksalad"
	"/Users/han058/go/myapp"
	"github.com/banksalad/myapp/config"
)

type MyappServer struct {
	cfg config.Config
}

func NewMyappServer(cfg config.Config) (*MyappServer, error) {
	return &MyappServer{cfg: cfg}, nil
}

func (s *MyappServer) Config() config.Config {
	return s.cfg
}

func (s *MyappServer) RegisterServer(srv *grpc.Server) {
	myapp.RegisterMyappServer(srv, s)
}

func NewGRPCServer(cfg config.Config) (*grpc.Server, error) {
	logrus.ErrorKey = "grpc.error"

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			banksalad.UnaryServerInterceptor(cfg.StatsdClient().CloneWithPrefixExtension(".grpc"), log),
		),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 30 * time.Second,
		}),
	)

	myappServer, err := NewMyappServer(cfg)
	if err != nil {
		return nil, err
	}

	myapp.RegisterMyappServer(grpcServer, myappServer)
	reflection.Register(grpcServer)

	return grpcServer, nil
}
