package main

import (
	"example.com/service-clean/libs/config"
	grpcdelivery "example.com/service-clean/module/todo/delivery/grpc"
	"example.com/service-clean/module/todo/usecase"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gormlib "example.com/service-clean/libs/gorm"
	grpclib "example.com/service-clean/libs/grpc"

	"example.com/service-clean/module/todo/repository"
)

func init() {
	config.BindEnv(config.EnvList())
}

func main() {
	// Init gRPC
	grpcServer := initGrpcServer()
	// Init Gorm
	//gormDB := initGormDB()

	gormDB := &gorm.DB{}
	// Init
	repo := repository.New(gormDB)

	// Init usecase
	useCase := usecase.New(repo)

	// Init delivery
	grpcdelivery.New(grpcServer, useCase)

	//Run gRPC service
	if err := grpclib.Start(grpcServer); err != nil {
		log.Fatal(err)
	}
}

func initGrpcServer() *grpc.Server {
	grpcServer, err := grpclib.NewServer(insecure.NewCredentials())
	if err != nil {
		log.Fatal(err)
	}
	return grpcServer
}

func initGormDB() *gorm.DB {
	gormDB, err := gormlib.NewGormProvider(postgres.New(postgres.Config{
		DSN:                  gormlib.GetDSNByEnv(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}))
	if err != nil {
		log.Fatal(err)
	}

	return gormDB
}
