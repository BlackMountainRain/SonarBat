package service

import (
	pb "sonar-bat/api/resource/v1"
	"sonar-bat/internal/resource/biz"
)

type ResourceService struct {
	resource biz.Resource

	pb.UnimplementedResourceServer
}

func NewResourceService(resource biz.Resource) *ResourceService {
	return &ResourceService{
		resource: resource,
	}
}
