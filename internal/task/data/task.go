/**
 * Package data
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/8
 */

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"sonar-bat/ent"
	"sonar-bat/internal/task/biz"
)

type taskRepo struct {
	data *Data
	log  *log.Helper
}

// NewTaskRepo .
func NewTaskRepo(data *Data, logger log.Logger) biz.TaskRepo {
	return &taskRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *taskRepo) Create(ctx context.Context, task *ent.Task) (*ent.Task, error) {
	return r.data.cli.Task.Create().
		SetName(task.Name).SetStatus(task.Status).SetComment(task.Comment).
		SetCreatedBy(task.CreatedBy).SetUpdatedBy(task.UpdatedBy).
		Save(ctx)
}

func (r *taskRepo) Save(ctx context.Context, task *ent.Task) (*ent.Task, error) {
	return r.data.cli.Task.UpdateOne(task).Save(ctx)
}

func (r *taskRepo) Update(ctx context.Context, task *ent.Task) (*ent.Task, error) {
	return r.data.cli.Task.UpdateOne(task).Save(ctx)
}

func (r *taskRepo) FindByID(ctx context.Context, id uuid.UUID) (*ent.Task, error) {
	return r.data.cli.Task.Get(ctx, id)
}

func (r *taskRepo) ListAll(ctx context.Context) ([]*ent.Task, error) {
	return r.data.cli.Task.Query().All(ctx)
}
