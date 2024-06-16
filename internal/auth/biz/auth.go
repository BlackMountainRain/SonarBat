/**
 * Package biz
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/13
 */

package biz

import (
	"context"
	"github.com/google/uuid"
	"sonar-bat/api/auth/v1"
	"sonar-bat/ent"
	"sonar-bat/internal/auth/biz/util"
	"sonar-bat/internal/conf"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound             = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
	ErrEmailOrPasswordIncorrect = errors.NotFound(v1.ErrorReason_INVALID_EMAIL_OR_PASSWORD.String(), "email or password incorrect")
)

// AuthRepo is a Auth repo.
type AuthRepo interface {
	Create(context.Context, *ent.User) (*ent.User, error)
	Save(context.Context, *ent.User) (*ent.User, error)
	Update(context.Context, *ent.User) (*ent.User, error)
	FindByID(context.Context, uuid.UUID) (*ent.User, error)
	FindByEmail(context.Context, string) (*ent.User, error)
	ListAll(context.Context) ([]*ent.User, error)
}

// AuthUseCase is a Auth use case.
type AuthUseCase struct {
	cfg  *conf.Auth
	repo AuthRepo
	log  *log.Helper
}

// NewAuthUseCase creates a new AuthUseCase
func NewAuthUseCase(c *conf.Auth, repo AuthRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		cfg:  c,
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// SignIn signs in and returns a token.
func (uc *AuthUseCase) SignIn(ctx context.Context, params *v1.SignInRequest) (string, error) {
	usr, err := uc.repo.FindByEmail(ctx, params.Email)
	if err != nil || usr == nil {
		return "", ErrEmailOrPasswordIncorrect
	}

	if !util.CompareHashAndPassword(params.Password, usr.Password) {
		return "", ErrEmailOrPasswordIncorrect
	}

	return util.GenerateToken([]byte(uc.cfg.Jwt.Secret), uc.cfg.Jwt.Expiration.AsDuration(), map[string]interface{}{
		"uid":      usr.ID,
		"username": usr.Username,
		"email":    usr.Email,
		"status":   usr.Status,
	})
}

// SignUp creates a new user.
func (uc *AuthUseCase) SignUp(ctx context.Context, params *v1.SignUpRequest) (string, error) {
	// confirm the user does not exist
	usr, err := uc.repo.FindByEmail(ctx, params.Email)
	if err != nil && !ErrUserNotFound.Is(err) {
		return "", err
	}
	if usr != nil {
		return "", v1.ErrorUserAlreadyExists("user already exists")
	}

	hashedPass, err := util.HashPassword(params.Password)
	if err != nil {
		return "", err
	}

	// create the user
	data := &ent.User{
		Username: params.Email,
		Email:    params.Email,
		Password: hashedPass,
	}
	user, err := uc.repo.
		Create(ctx, data)
	if err != nil {
		return "", err
	}

	// generate the token
	return util.GenerateToken([]byte(uc.cfg.Jwt.Secret), uc.cfg.Jwt.Expiration.AsDuration(), map[string]interface{}{
		"uid":      user.ID,
		"username": user.Username,
		"email":    user.Email,
		"status":   user.Status,
	})
}

func (uc *AuthUseCase) UpdateAuth(ctx context.Context, t *ent.User) (*ent.User, error) {
	return uc.repo.Update(ctx, t)
}

func (uc *AuthUseCase) GetAuthByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *AuthUseCase) GetAuths(ctx context.Context) ([]*ent.User, error) {
	return uc.repo.ListAll(ctx)
}
