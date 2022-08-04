package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/banksalad/go-banksalad/grpcgateway/v2"
	"github.com/hanyoung-banksalad/myapp"
	"github.com/hanyoung-banksalad/myapp/config"
)

func NewHTTPServer(ctx context.Context, cfg config.Config) (*http.Server, error) {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		),
		runtime.WithErrorHandler(grpcgateway.ExternalHTTPErrorHandler),
		runtime.WithIncomingHeaderMatcher(grpcgateway.IncomingHTTPHeaderMatcher()),
		runtime.WithOutgoingHeaderMatcher(grpcgateway.OutgoingHTTPHeaderMatcher()),
	)
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := myapp.RegisterMyappHandlerFromEndpoint(
		ctx,
		mux,
		cfg.Setting().GRPCServerEndpoint,
		options,
	); err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:    ":" + cfg.Setting().HTTPServerPort,
		Handler: mux,
	}

	return server, nil
}
