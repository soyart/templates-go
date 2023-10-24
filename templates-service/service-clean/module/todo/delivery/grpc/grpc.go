package grpc

import (
	"example.com/service-clean/module/todo/delivery/grpc/model/gen/grpc/protocol_buffer"
	"google.golang.org/grpc"

	"example.com/service-clean/module/todo/domain"
)

type grpcDelivery struct {
	protocol_buffer.UnimplementedTodoServiceServer
	useCase domain.UseCase
}

func New(server *grpc.Server, useCase domain.UseCase) {
	protocol_buffer.RegisterTodoServiceServer(server, &grpcDelivery{useCase: useCase})
}
