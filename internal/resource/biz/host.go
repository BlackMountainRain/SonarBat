/**
 * Package biz
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/27
 */

package biz

import (
	"context"
	"github.com/google/uuid"
	"net"
	v1 "sonar-bat/api/resource/v1"
	"sonar-bat/ent"
	"sonar-bat/ent/schema"
	"time"
)

// HostRepo is a repository interface for Host.
type HostRepo interface {
	Create(context.Context, *ent.Host) (*ent.Host, error)
	Save(context.Context, *ent.Host) (*ent.Host, error)
	Update(context.Context, *ent.Host) (*ent.Host, error)
	FindByID(context.Context, uuid.UUID) (*ent.Host, error)
	ListAll(context.Context) ([]*ent.Host, error)
}

// CreateHost creates a Host, and returns the new Host.
func (uc *Manager) CreateHost(ctx context.Context, params *v1.CreateHostRequest) (string, error) {
	fakeUser := uuid.New()
	ips := make([]schema.IpWithInfo, 0, len(params.Ips))
	for _, v := range params.Ips {
		ip := net.ParseIP(v)
		if ip == nil {
			return "", nil
		}
		ips = append(ips, schema.IpWithInfo{
			IP:        ip,
			IsPrivate: ip.IsPrivate(),
			IsIPv6:    ip.To16() != nil && ip.To4() == nil,
		})
	}
	additions := make(map[string]interface{})
	for k, v := range params.Additions {
		additions[k] = v
	}
	host := &ent.Host{
		UpdatedBy:        fakeUser,
		CreatedBy:        fakeUser,
		Status:           params.Status,
		Name:             params.Name,
		LiveAt:           time.Time{},
		IsAgentInstalled: params.IsAgentInstalled,
		AgentVersion:     params.AgentVersion,
		Ips:              ips,
		NetType:          int16(params.NetType.Number()),
		Additions:        additions,
	}
	host, err := uc.hostRepo.Create(ctx, host)
	if err != nil {
		return "", err
	}
	return host.ID.String(), nil
}

func (uc *Manager) UpdateHost(ctx context.Context, params *v1.UpdateHostRequest) (*ent.Host, error) {
	return nil, nil
}

func (uc *Manager) GetHostByID(ctx context.Context, id uuid.UUID) (*ent.Host, error) {
	return uc.hostRepo.FindByID(ctx, id)
}

func (uc *Manager) GetHosts(ctx context.Context) ([]*ent.Host, error) {
	return uc.hostRepo.ListAll(ctx)
}
