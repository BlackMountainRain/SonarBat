/**
 * Package biz
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/8
 */

package biz

import (
	"context"
	"github.com/google/uuid"
	"sonar-bat/ent"

	v1 "sonar-bat/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// TaskRepo is a Task repo.
type TaskRepo interface {
	Create(context.Context, *ent.Task) (*ent.Task, error)
	Save(context.Context, *ent.Task) (*ent.Task, error)
	Update(context.Context, *ent.Task) (*ent.Task, error)
	FindByID(context.Context, uuid.UUID) (*ent.Task, error)
	ListAll(context.Context) ([]*ent.Task, error)
}

// TaskUseCase is a Task use case.
type TaskUseCase struct {
	repo TaskRepo
	log  *log.Helper
}

// NewTaskUseCase creates a new TaskUseCase
func NewTaskUseCase(repo TaskRepo, logger log.Logger) *TaskUseCase {
	return &TaskUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateTask creates a Task, and returns the new Task.
func (uc *TaskUseCase) CreateTask(ctx context.Context, t *ent.Task) (*ent.Task, error) {
	return uc.repo.Create(ctx, t)
}

func (uc *TaskUseCase) UpdateTask(ctx context.Context, t *ent.Task) (*ent.Task, error) {
	return uc.repo.Update(ctx, t)
}

func (uc *TaskUseCase) GetTaskByID(ctx context.Context, id uuid.UUID) (*ent.Task, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *TaskUseCase) GetTasks(ctx context.Context) ([]*ent.Task, error) {
	return uc.repo.ListAll(ctx)
}
