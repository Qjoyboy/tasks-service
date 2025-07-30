package grpc

import (
	"context"
	"fmt"

	taskpb "github.com/Qjoyboy/project-proto/proto/task"
	userpb "github.com/Qjoyboy/project-proto/proto/user"
	"github.com/Qjoyboy/tasks-service/internal/task"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc        task.TaskService
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc task.TaskService, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}

func (h Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	t, err := h.svc.CreateTaskByUserId(req.Title, req.UserId, false)
	if err != nil {
		return nil, err
	}

	return &taskpb.CreateTaskResponse{Task: &taskpb.Task{Id: t.ID, UserId: t.UserID, Title: t.Text, IsDone: t.IsDone}}, nil
}

func (h Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	t, err := h.svc.GetTaskByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &taskpb.GetTaskResponse{Task: &taskpb.Task{Id: t.ID, UserId: t.UserID, Title: t.Text, IsDone: t.IsDone}}, nil
}

func (h Handler) ListTasks(ctx context.Context, _ *emptypb.Empty) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.GetTasks()
	if err != nil {
		return nil, err
	}
	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:     t.ID,
			UserId: t.UserID,
			Title:  t.Text,
			IsDone: t.IsDone,
		})
	}
	return &taskpb.ListTasksResponse{Task: pbTasks}, nil
}

// func (h Handler) GetTasksForUser(ctx context.Context, req *taskpb.)

func (h Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	taskId := req.Task.Id
	taskToUpdate, err := h.svc.GetTaskByID(taskId)

	if err != nil {
		return nil, err
	}

	taskToUpdate.Text = req.Task.Title
	taskToUpdate.IsDone = req.Task.IsDone

	upTask, err := h.svc.UpdateTask(taskToUpdate.ID, taskToUpdate.Text, taskToUpdate.IsDone)
	if err != nil {
		return nil, err
	}
	pbTask := &taskpb.Task{
		Id:     upTask.ID,
		Title:  upTask.Text,
		IsDone: upTask.IsDone,
		UserId: upTask.UserID,
	}
	return &taskpb.UpdateTaskResponse{Task: pbTask}, nil
}

func (h Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	err := h.svc.DeleteTask(req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete task %d: %w", req.Id, err)
	}

	return &taskpb.DeleteTaskResponse{}, err

}
