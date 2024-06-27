/**
 * Package data
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/27
 */

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"sonar-bat/ent"
	"sonar-bat/internal/resource/biz"
)

type hostRepo struct {
	data *Data
	log  *log.Helper
}

// NewHostRepo .
func NewHostRepo(data *Data, logger log.Logger) biz.HostRepo {
	return &hostRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *hostRepo) Create(ctx context.Context, resource *ent.Host) (*ent.Host, error) {
	return r.data.cli.Host.Create().
		SetName(resource.Name).SetStatus(resource.Status).SetLiveAt(resource.LiveAt).
		SetIsAgentInstalled(resource.IsAgentInstalled).SetIps(resource.Ips).SetAdditions(resource.Additions).
		SetCreatedBy(resource.CreatedBy).SetUpdatedBy(resource.UpdatedBy).
		Save(ctx)
}

func (r *hostRepo) Save(ctx context.Context, resource *ent.Host) (*ent.Host, error) {
	return r.data.cli.Host.UpdateOne(resource).Save(ctx)
}

func (r *hostRepo) Update(ctx context.Context, resource *ent.Host) (*ent.Host, error) {
	return r.data.cli.Host.UpdateOne(resource).Save(ctx)
}

func (r *hostRepo) FindByID(ctx context.Context, id uuid.UUID) (*ent.Host, error) {
	return r.data.cli.Host.Get(ctx, id)
}

func (r *hostRepo) ListAll(ctx context.Context) ([]*ent.Host, error) {
	return r.data.cli.Host.Query().All(ctx)
}
