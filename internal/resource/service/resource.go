package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "sonar-bat/api/resource/v1"
	"sonar-bat/internal/resource/biz"
	"sonar-bat/internal/runtime"
)

type ResourceService struct {
	resource    biz.Resource
	runtimeInfo runtime.Info

	pb.UnimplementedResourceServer
}

func NewResourceService(resource biz.Resource, runtimeInfo runtime.Info) *ResourceService {
	return &ResourceService{
		resource:    resource,
		runtimeInfo: runtimeInfo,
	}
}

func (s *ResourceService) HealthCheck(_ context.Context, _ *pb.HealthRequest) (*pb.HealthReply, error) {
	return &pb.HealthReply{
		Status:  pb.Status_UP,
		Version: s.runtimeInfo.Version,
		Uptime:  timestamppb.New(s.runtimeInfo.Uptime),
	}, nil
}
