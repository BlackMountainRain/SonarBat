/**
 * Package biz
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/27
 */

package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	v1 "sonar-bat/api/resource/v1"
	"sonar-bat/ent"
)

// Resource is a resource manager.
type Resource interface {
	CreateHost(context.Context, *v1.CreateHostRequest) (string, error)
	UpdateHost(context.Context, *v1.UpdateHostRequest) (*ent.Host, error)
	GetHostByID(context.Context, uuid.UUID) (*ent.Host, error)
	GetHosts(context.Context) ([]*ent.Host, error)
}

// Manager is a Host manager.
type Manager struct {
	hostRepo HostRepo
	log      *log.Helper
}

// NewHostManager creates a new Manager
func NewHostManager(repo HostRepo, logger log.Logger) Resource {
	return &Manager{hostRepo: repo, log: log.NewHelper(logger)}
}
