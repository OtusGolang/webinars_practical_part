package task

import (
	"context"
	"taskmanager/api/proto"
)

type Handler struct {
	proto.UnimplementedTaskServiceServer // Embed the unimplemented server
	service                              Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateTask(ctx context.Context, req *proto.TaskRequest) (*proto.TaskResponse, error) {
	return h.service.CreateTask(ctx, req)
}

func (h *Handler) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.TaskResponse, error) {
	return h.service.GetTask(ctx, req)
}

func (h *Handler) GetAllTasks(ctx context.Context, req *proto.GetAllTasksRequest) (*proto.GetAllTasksResponse, error) {
	return h.service.GetAllTasks(ctx, req)
}

func (h *Handler) UpdateTask(ctx context.Context, req *proto.TaskRequest) (*proto.TaskResponse, error) {
	return h.service.UpdateTask(ctx, req)
}

func (h *Handler) DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	return h.service.DeleteTask(ctx, req)
}
