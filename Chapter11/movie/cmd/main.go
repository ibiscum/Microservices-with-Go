package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/ibiscum/Microservices-with-Go/Chapter11/gen"
	"github.com/ibiscum/Microservices-with-Go/Chapter11/movie/internal/controller/movie"
	metadatagateway "github.com/ibiscum/Microservices-with-Go/Chapter11/movie/internal/gateway/metadata/grpc"
	ratinggateway "github.com/ibiscum/Microservices-with-Go/Chapter11/movie/internal/gateway/rating/grpc"
	grpchandler "github.com/ibiscum/Microservices-with-Go/Chapter11/movie/internal/handler/grpc"
	"github.com/ibiscum/Microservices-with-Go/Chapter11/pkg/discovery"
	"github.com/ibiscum/Microservices-with-Go/Chapter11/pkg/discovery/consul"
	"github.com/ibiscum/Microservices-with-Go/Chapter11/pkg/tracing"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
)

const serviceName = "movie"

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Panic(err)
		}
	}()

	f, err := os.Open("base.yaml")
	if err != nil {
		logger.Fatal("Failed to open configuration", zap.Error(err))
	}
	var cfg config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		logger.Fatal("Failed to parse configuration", zap.Error(err))
	}
	port := cfg.API.Port

	logger.Info("Starting the movie service", zap.Int("port", port))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tp, err := tracing.NewJaegerProvider(cfg.Jaeger.URL, serviceName)
	if err != nil {
		logger.Fatal("Failed to initialize Jaeger provider", zap.Error(err))
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal("Failed to shut down Jaeger prodiver", zap.Error(err))
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				logger.Error("Failed to report healthy state", zap.Error(err))
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer func() {
		err := registry.Deregister(ctx, instanceID, serviceName)
		if err != nil {
			log.Fatal(err)
		}
	}()
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", zap.Error(err))
	}
	srv := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	reflection.Register(srv)
	gen.RegisterMovieServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to start the gRPC server: %v", zap.Error(err))
	}
}

// type limiter struct {
// 	l *rate.Limiter
// }

// func newLimiter(limit int, burst int) *limiter {
// 	return &limiter{rate.NewLimiter(rate.Limit(limit), burst)}
// }

// func (l *limiter) Limit() bool {
// 	return l.l.Allow()
// }
