package grpc

import (
	protobuf "example.com/service-clean/module/todo/delivery/grpc/model/gen/grpc/protocol_buffer"
	"example.com/service-clean/module/todo/domain"
)

func toAppStatus(status protobuf.Status) domain.Status {
	switch status {
	case protobuf.Status_Todo:
		return domain.StatusTodo
	case protobuf.Status_Done:
		return domain.StatusDone
	case protobuf.Status_InProgress:
		return domain.StatusInProgress
	default:
		return domain.StatusTodo
	}
}
