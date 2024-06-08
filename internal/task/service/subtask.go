package service

import (
	"context"
	"sonar-bat/internal/task/biz"

	pb "sonar-bat/api/task/v1"
)

type SubtaskService struct {
	pb.UnimplementedSubtaskServer
}

func NewSubtaskService(biz biz.TaskRepo) *SubtaskService {
	return &SubtaskService{}
}

func (s *SubtaskService) CreateSubtask(ctx context.Context, req *pb.CreateSubtaskRequest) (*pb.CreateSubtaskReply, error) {
	return &pb.CreateSubtaskReply{}, nil
}
func (s *SubtaskService) UpdateSubtask(ctx context.Context, req *pb.UpdateSubtaskRequest) (*pb.UpdateSubtaskReply, error) {
	return &pb.UpdateSubtaskReply{}, nil
}
func (s *SubtaskService) OverwriteSubtask(ctx context.Context, req *pb.OverwriteSubtaskRequest) (*pb.OverwriteSubtaskReply, error) {
	return &pb.OverwriteSubtaskReply{}, nil
}
func (s *SubtaskService) DeleteSubtask(ctx context.Context, req *pb.DeleteSubtaskRequest) (*pb.DeleteSubtaskReply, error) {
	return &pb.DeleteSubtaskReply{}, nil
}
func (s *SubtaskService) GetSubtask(ctx context.Context, req *pb.GetSubtaskRequest) (*pb.GetSubtaskReply, error) {
	return &pb.GetSubtaskReply{}, nil
}
func (s *SubtaskService) GetSubtasks(ctx context.Context, req *pb.GetSubtasksRequest) (*pb.GetSubtasksReply, error) {
	return &pb.GetSubtasksReply{}, nil
}
