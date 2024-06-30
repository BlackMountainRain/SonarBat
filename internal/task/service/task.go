package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sonar-bat/ent"
	"sonar-bat/internal/runtime"
	"sonar-bat/internal/task/biz"
	"time"

	pb "sonar-bat/api/task/v1"
)

type TaskService struct {
	task        *biz.TaskUseCase
	runtimeInfo runtime.Info

	pb.UnimplementedTaskServer
}

func NewTaskService(task *biz.TaskUseCase, runtimeInfo runtime.Info) *TaskService {
	return &TaskService{
		task:        task,
		runtimeInfo: runtimeInfo,
	}
}

func (s *TaskService) HealthCheck(_ context.Context, _ *pb.HealthRequest) (*pb.HealthReply, error) {
	return &pb.HealthReply{
		Status:  pb.Status_UP,
		Version: s.runtimeInfo.Version,
		Uptime:  timestamppb.New(s.runtimeInfo.Uptime),
	}, nil
}

func (s *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskReply, error) {
	task := &ent.Task{
		Name:    req.Name,
		Status:  req.Status,
		Comment: req.Comment,
	}
	task, err := s.task.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTaskReply{
		Id: task.ID.String(),
	}, nil
}
func (s *TaskService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskReply, error) {
	return &pb.UpdateTaskReply{}, nil
}
func (s *TaskService) OverwriteTask(ctx context.Context, req *pb.OverwriteTaskRequest) (*pb.OverwriteTaskReply, error) {
	return &pb.OverwriteTaskReply{}, nil
}
func (s *TaskService) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskReply, error) {
	return &pb.DeleteTaskReply{}, nil
}
func (s *TaskService) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskReply, error) {
	return &pb.GetTaskReply{}, nil
}
func (s *TaskService) GetTasks(ctx context.Context, req *pb.GetTasksRequest) (*pb.GetTasksReply, error) {
	tasks, err := s.task.GetTasks(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.SingleTask, 0, len(tasks))
	for _, task := range tasks {
		res = append(res, &pb.SingleTask{
			Id:        task.ID.String(),
			Status:    task.Status,
			Name:      task.Name,
			Comment:   task.Comment,
			UpdatedBy: task.UpdatedBy.String(),
			CreatedBy: task.CreatedBy.String(),
			CreatedAt: task.CreatedAt.Format(time.RFC3339),
			UpdatedAt: task.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &pb.GetTasksReply{
		Tasks: res,
	}, nil
}
