/**
 * Package data
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/13
 */

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"sonar-bat/ent"
	"sonar-bat/internal/auth/biz"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *authRepo) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	return r.data.cli.User.Create().
		SetUsername(user.Username).SetEmail(user.Email).SetPassword(user.Password).
		Save(ctx)
}

func (r *authRepo) Save(ctx context.Context, user *ent.User) (*ent.User, error) {
	return nil, nil
}

func (r *authRepo) Update(ctx context.Context, user *ent.User) (*ent.User, error) {
	return nil, nil
}

func (r *authRepo) FindByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return nil, nil
}

func (r *authRepo) ListAll(ctx context.Context) ([]*ent.User, error) {
	return nil, nil
}
