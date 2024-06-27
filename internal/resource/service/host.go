/**
 * Package service
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/27
 */

package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "sonar-bat/api/resource/v1"
)

func (s *ResourceService) CreateHost(ctx context.Context, req *pb.CreateHostRequest) (*pb.CreateHostReply, error) {
	hostId, err := s.resource.CreateHost(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateHostReply{
		Id: hostId,
	}, nil
}
func (s *ResourceService) UpdateHost(ctx context.Context, req *pb.UpdateHostRequest) (*pb.UpdateHostReply, error) {
	return &pb.UpdateHostReply{}, nil
}
func (s *ResourceService) OverwriteHost(ctx context.Context, req *pb.OverwriteHostRequest) (*pb.OverwriteHostReply, error) {
	return &pb.OverwriteHostReply{}, nil
}
func (s *ResourceService) DeleteHost(ctx context.Context, req *pb.DeleteHostRequest) (*pb.DeleteHostReply, error) {
	return &pb.DeleteHostReply{}, nil
}
func (s *ResourceService) GetHost(ctx context.Context, req *pb.GetHostRequest) (*pb.GetHostReply, error) {
	return &pb.GetHostReply{}, nil
}
func (s *ResourceService) GetHosts(ctx context.Context, _ *pb.GetHostsRequest) (*pb.GetHostsReply, error) {
	hosts, err := s.resource.GetHosts(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.SingleHost, 0, len(hosts))
	for _, host := range hosts {
		ips := make([]string, 0, len(host.Ips))
		for _, ip := range host.Ips {
			ips = append(ips, ip.IP.String())
		}
		additions := make(map[string]string)
		for k, v := range host.Additions {
			additions[k] = v.(string)
		}
		res = append(res, &pb.SingleHost{
			Id:               host.ID.String(),
			Status:           host.Status,
			Name:             host.Name,
			LiveAt:           timestamppb.New(host.LiveAt),
			IsAgentInstalled: host.IsAgentInstalled,
			AgentVersion:     host.AgentVersion,
			Ips:              ips,
			NetType:          pb.NetType(host.NetType),
			Additions:        additions,
		})
	}
	return &pb.GetHostsReply{
		Hosts: res,
	}, nil
}
