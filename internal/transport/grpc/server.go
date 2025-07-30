package grpc

import (
	"fmt"
	"net"

	taskpb "github.com/Qjoyboy/project-proto/proto/task"
	userpb "github.com/Qjoyboy/project-proto/proto/user"
	"github.com/Qjoyboy/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc task.TaskService, uc userpb.UserServiceClient) error {
	lis, _ := net.Listen("tcp", ":50052")
	grpcSrv := grpc.NewServer()
	handler := NewHandler(svc, uc)
	taskpb.RegisterTaskServiceServer(grpcSrv, handler)
	fmt.Println("Starting gRPC server on :50052...")

	return grpcSrv.Serve(lis)
}

//dd31a983-5913-4a6e-9451-0929ce9f011d
//a8e859da-a694-4266-8142-a56961f4fa85
