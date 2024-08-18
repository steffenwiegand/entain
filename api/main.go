package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"git.neds.sh/matty/entain/api/proto/racing"
	"git.neds.sh/matty/entain/api/proto/sports"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	apiEndpoint           = flag.String("api-endpoint", "localhost:8000", "API endpoint")
	racingServiceEndpoint = flag.String("grpc-racingService-endpoint", "localhost:9000", "gRPC racing service endpoint")
	sportsServiceEndpoint = flag.String("grpc-sportsService-endpoint", "localhost:10000", "gRPC sports service endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Printf("failed running api server: %s\n", err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	if err := racing.RegisterRacingHandlerFromEndpoint(
		ctx,
		mux,
		*racingServiceEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}
	if err := sports.RegisterSportsHandlerFromEndpoint(
		ctx,
		mux,
		*sportsServiceEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}

	log.Printf("API server listening on: %s\n", *apiEndpoint)

	return http.ListenAndServe(*apiEndpoint, mux)
}
