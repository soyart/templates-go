package grpc

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	envGrpcPort = "GRPC_PORT"
	tcpProtocol = "tcp"
)

func Start(grpcServer *grpc.Server) error {
	netListener, err := net.Listen(tcpProtocol, viper.GetString(envGrpcPort))
	if err != nil {
		return fmt.Errorf("%s: %s: %w", "failed to listen the network", viper.GetString(envGrpcPort), err)
	}

	//Subscribe Signal to graceful stop grpcServer
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	log.Printf("Listening and serving gRPC on %v", netListener.Addr())
	if err = grpcServer.Serve(netListener); err != nil {
		log.Fatal(fmt.Errorf("%s: %v", "failed to start gRPC server", err))
	}
	//Waiting for signal to graceful stop
	<-stopChan
	grpcServer.GracefulStop()

	return nil
}

func NewServer(tlsCredentials credentials.TransportCredentials) (*grpc.Server, error) {
	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)

	// Setup health checking
	// https://github.com/grpc/grpc-go/tree/master/examples/features/health
	healthcheck := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthcheck)

	return grpcServer, nil
}
