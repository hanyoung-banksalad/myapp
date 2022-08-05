package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/hanyoung-banksalad/myapp/idl"
	"github.com/hanyoung-banksalad/myapp/client"
	"github.com/hanyoung-banksalad/myapp/config"
	"github.com/hanyoung-banksalad/myapp/server"
	"github.com/banksalad/go-banksalad"
)

func main() {
	ctx := context.Background()

	logrus.SetFormatter(&logrus.JSONFormatter{})
	log := logrus.StandardLogger()

	setting := config.NewSetting()

	statsdCli := banksalad.NewStatsdClient(banksalad.StatsdOptions{
		StatsdAddr: fmt.Sprintf("localhost:%s", setting.StatsdUDPPort),
		Prefix:     fmt.Sprintf("%s.%s", setting.Namespace, setting.ServiceName),
		Logger:     log,
	})
	defer func() {
		if err := statsdCli.Close(); err != nil {
			log.WithError(err).Error("failed to close statsd client")
		}
	}()

	logrus.AddHook(banksalad.NewLogLevelStatHook(statsdCli))
	logrus.AddHook(banksalad.NewLogStacktraceHook())

	closeTracer, err := banksalad.InitializeTracer(banksalad.TracerOptions{
		ServiceName: setting.ServiceName,
		Namespace:   setting.Namespace,
		Env:         setting.Env,
	})
	if err != nil {
		log.WithError(err).Fatal("failed to initialize jaeger client")
	}
	defer closeTracer()

	authCli := client.GetAuthClient(setting.AuthGRPCEndpoint, setting.ServiceName)

	cfg := config.NewConfig(
		setting,
		statsdCli,
		authCli,
	)

	grpcServer, err := server.NewGRPCServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	httpServer, err := server.NewHTTPServer(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		lis, err := net.Listen("tcp", ":"+cfg.Setting().GRPCServerPort)
		if err != nil {
			log.WithError(err).Fatal("net.Listen")
		}

		log.WithField("port", cfg.Setting().GRPCServerPort).Info("starting myapp gRPC server")
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatal(err)
		}
	}()

	go func() {
		log.WithField("port", cfg.Setting().HTTPServerPort).Info("starting myapp HTTP server")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	<-quit

	time.Sleep(time.Duration(cfg.Setting().GracefulShutdownTimeoutMs) * time.Millisecond)

	log.Info("stopping myapp HTTP server")
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Info("stopping myapp gRPC server")
	grpcServer.GracefulStop()
}

type imageServer struct{}

var _ pb.ImageServer = (*imageServer)(nil)

func (s *imageServer) GetImage(req *pb.GetImageRequest, stream pb.Image_GetImageServer) error {
	f, err := os.Open("images/" + req.Path)
	if err != nil {
		return status.Error(codes.NotFound, "file not found")
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.WithError(err).Error("failed to close file")
		}
	}()

	// Maximum 16KB size per stream.
	buf := make([]byte, 16*2<<10)
	for {
		num, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err := stream.Send(&pb.GetImageResponse{Data: buf[:num]}); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
