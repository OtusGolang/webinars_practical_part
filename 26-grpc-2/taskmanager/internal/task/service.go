package task

import (
	"context"
	"errors"
	"sync"
	"taskmanager/api/proto"
	"time"
)

type Service interface {
	CreateTask(ctx context.Context, req *proto.TaskRequest) (*proto.TaskResponse, error)
	GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.TaskResponse, error)
	GetAllTasks(ctx context.Context, req *proto.GetAllTasksRequest) (*proto.GetAllTasksResponse, error)
	UpdateTask(ctx context.Context, req *proto.TaskRequest) (*proto.TaskResponse, error)
	DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error)
}

type service struct {
	mu    sync.Mutex // For thread-safety
	tasks map[string]*proto.TaskResponse
}

func NewService() Service {
	return &service{tasks: make(map[string]*proto.TaskResponse)}
}

func (s *service) CreateTask(ctx context.Context, req *proto.TaskRequest) (*proto.TaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Assume each task has a unique title for simplicity
	if _, exists := s.tasks[req.Title]; exists {
		return nil, errors.New("task already exists")
	}

	task := &proto.TaskResponse{
		TaskId:      req.Title, // Assume title is unique and use as ID
		Title:       req.Title,
		Description: req.Description,
		Status:      "Pending",
	}
	s.tasks[req.Title] = task

	// Set a timer to mark the task as finished after 1 minute
	time.AfterFunc(1*time.Minute, func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		if task, exists := s.tasks[req.Title]; exists && task.Status != "Deleted" { // Check if task still exists and is not deleted
			task.Status = "Finished"
			// Optionally, log the status change
			// log.Printf("Task %s status changed to Finished", req.Title)
		}
	})

	return task, nil
}

func (s *service) GetAllTasks(ctx context.Context, _ *proto.GetAllTasksRequest) (*proto.GetAllTasksResponse, error) {
	var tasks []*proto.TaskResponse
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return &proto.GetAllTasksResponse{Tasks: tasks}, nil
}

func (s *service) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.TaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[req.TaskId]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (s *service) UpdateTask(ctx context.Context, req *proto.TaskRequest) (*proto.TaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[req.Title] // Assume title is unique and use as ID
	if !exists {
		return nil, errors.New("task not found")
	}

	// Update task details
	task.Description = req.Description
	task.Status = req.Status

	return task, nil
}

func (s *service) DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[req.TaskId]; !exists {
		return nil, errors.New("task not found")
	}

	delete(s.tasks, req.TaskId)
	return &proto.DeleteTaskResponse{Result: "Successfully deleted"}, nil
}
