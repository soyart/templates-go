package grpc

import (
	"context"
	"example.com/service-clean/module/todo/delivery/grpc/model/gen/grpc/protocol_buffer"
	"example.com/service-clean/module/todo/domain"
)

func (delivery grpcDelivery) CreateTodoHandler(ctx context.Context, req *protocol_buffer.TodoCreateRequestModel) (*protocol_buffer.TodoCreateResponseModel, error) {
	todoDate, err := convertStrToTime(req.TodoDate)
	if err != nil {
		return nil, err
	}
	if err = delivery.useCase.CreateTodo(ctx, domain.TodoCreateInput{
		Title:       req.Title,
		Status:      toAppStatus(req.Status),
		Description: req.Description,
		TodoDate:    *todoDate,
	}); err != nil {
		return nil, err
	}

	return &protocol_buffer.TodoCreateResponseModel{}, nil
}
