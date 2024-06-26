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
	"sonar-bat/ent/user"
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
	return r.data.cli.User.UpdateOneID(user.ID).
		SetUsername(user.Username).SetEmail(user.Email).SetAvatarURL(user.AvatarURL).
		Save(ctx)
}

func (r *authRepo) FindByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return nil, nil
}

func (r *authRepo) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	usr, err := r.data.cli.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, biz.ErrUserNotFound
	}
	return usr, err
}

func (r *authRepo) ListAll(ctx context.Context) ([]*ent.User, error) {
	return nil, nil
}
