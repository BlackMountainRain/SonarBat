/**
 * Package biz
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/22
 */

package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/durationpb"
	v1 "sonar-bat/api/auth/v1"
	"sonar-bat/ent"
	"sonar-bat/internal/auth/biz/helper"
	"sonar-bat/internal/conf"
	"testing"
	"time"
)

type MockAuthRepo struct {
	mock.Mock
}

func (m *MockAuthRepo) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	args := m.Called(ctx, user)
	usr := args.Get(0)
	if usr != nil {
		return usr.(*ent.User), nil
	}
	return nil, args.Error(1)
}

func (m *MockAuthRepo) Save(ctx context.Context, user *ent.User) (*ent.User, error) {
	return nil, nil
}

func (m *MockAuthRepo) Update(ctx context.Context, user *ent.User) (*ent.User, error) {
	return nil, nil
}

func (m *MockAuthRepo) FindByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return nil, nil
}

func (m *MockAuthRepo) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	args := m.Called(ctx, email)
	user := args.Get(0)
	if user != nil {
		return user.(*ent.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthRepo) ListAll(ctx context.Context) ([]*ent.User, error) {
	return nil, nil
}

type MockLogger struct{}

func (m *MockLogger) Log(level log.Level, keyVals ...interface{}) error {
	return nil
}

// Initializes AuthUseCase with valid configuration, repository, and logger
func TestNewAuthUseCase_ValidConfig(t *testing.T) {
	// Arrange
	cfg := &conf.Auth{}
	repo := new(MockAuthRepo)
	var logger log.Logger = new(MockLogger)

	// Act
	useCase := NewAuthUseCase(cfg, repo, logger)

	// Assert
	assert.NotNil(t, useCase)
	assert.Equal(t, cfg, useCase.cfg)
	assert.Equal(t, repo, useCase.repo)
	assert.NotNil(t, useCase.log)
}

// Handles nil configuration gracefully
func TestNewAuthUseCase_NilConfig(t *testing.T) {
	// Arrange
	var cfg *conf.Auth = nil
	repo := new(MockAuthRepo)
	var logger log.Logger = new(MockLogger)

	// Act
	useCase := NewAuthUseCase(cfg, repo, logger)

	// Assert
	assert.NotNil(t, useCase)
	assert.Nil(t, useCase.cfg)
	assert.Equal(t, repo, useCase.repo)
	assert.NotNil(t, useCase.log)
}

// Valid email and password should return a valid JWT token
func TestSignIn_ValidEmailAndPassword(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockAuthRepo)
	mockLog := new(log.Helper)
	cfg := &conf.Auth{
		Jwt: &conf.Auth_Jwt{
			Secret:     "secret",
			Expiration: durationpb.New(time.Hour),
		},
	}
	uc := &AuthUseCase{cfg: cfg, repo: mockRepo, log: mockLog}
	email := "test@example.com"
	password := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)
	user := &ent.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
		Username: "test-user",
		Status:   true,
	}
	mockRepo.On("FindByEmail", ctx, email).Return(user, nil)

	// Act
	token, err := uc.SignIn(ctx, &v1.SignInRequest{Email: email, Password: password})

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// Non-existent email should return ErrEmailOrPasswordIncorrect
func TestSignIn_NonExistentEmail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockAuthRepo)
	mockLog := new(log.Helper)
	cfg := &conf.Auth{
		Jwt: &conf.Auth_Jwt{
			Secret:     "secret",
			Expiration: durationpb.New(time.Hour),
		},
	}
	uc := &AuthUseCase{cfg: cfg, repo: mockRepo, log: mockLog}
	email := "nonexistent@example.com"
	password := "password"
	mockRepo.On("FindByEmail", ctx, email).Return(nil, ErrUserNotFound)

	// Act
	token, err := uc.SignIn(ctx, &v1.SignInRequest{Email: email, Password: password})

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrEmailOrPasswordIncorrect, err)
	assert.Empty(t, token)
}

// Incorrect password should return ErrEmailOrPasswordIncorrect
func TestSignIn_IncorrectPassword(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockAuthRepo)
	mockLog := new(log.Helper)
	cfg := &conf.Auth{
		Jwt: &conf.Auth_Jwt{
			Secret:     "secret",
			Expiration: durationpb.New(time.Hour),
		},
	}
	uc := &AuthUseCase{cfg: cfg, repo: mockRepo, log: mockLog}
	email := "exist@example.com"
	password := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)
	user := &ent.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
		Username: "test-user",
		Status:   true,
	}
	mockRepo.On("FindByEmail", ctx, email).Return(user, nil)

	// Act
	token, err := uc.SignIn(ctx, &v1.SignInRequest{Email: email, Password: "incorrect"})

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrEmailOrPasswordIncorrect, err)
	assert.Empty(t, token)
}

// Sign up with a new email should return a valid JWT token
func TestSignUp_NewEmail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockAuthRepo)
	mockLog := new(log.Helper)
	cfg := &conf.Auth{
		Jwt: &conf.Auth_Jwt{
			Secret:     "secret",
			Expiration: durationpb.New(time.Hour),
		},
	}
	uc := &AuthUseCase{cfg: cfg, repo: mockRepo, log: mockLog}
	email := "newuser@example.com"
	password := "password"
	mockRepo.On("FindByEmail", ctx, email).Return(nil, ErrUserNotFound)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)
	newUser := &ent.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
		Username: email,
		Status:   true,
	}
	mockRepo.On("Create", ctx, mock.MatchedBy(func(user *ent.User) bool {
		return user.Email == newUser.Email && user.Username == newUser.Username
	})).Return(newUser, nil)

	// Act
	token, err := uc.SignUp(ctx, &v1.SignUpRequest{Email: email, Password: password})

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// Sign up with an existing email should return ErrorUserAlreadyExists
func TestSignUp_ExistingEmail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockAuthRepo)
	mockLog := new(log.Helper)
	cfg := &conf.Auth{
		Jwt: &conf.Auth_Jwt{
			Secret:     "secret",
			Expiration: durationpb.New(time.Hour),
		},
	}
	uc := &AuthUseCase{cfg: cfg, repo: mockRepo, log: mockLog}
	email := "user@example.com"
	password := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)
	user := &ent.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
		Username: email,
		Status:   true,
	}
	mockRepo.On("FindByEmail", ctx, email).Return(user, nil)

	// Act
	token, err := uc.SignUp(ctx, &v1.SignUpRequest{Email: email, Password: password})

	// Assert
	assert.Error(t, err)
	assert.Equal(t, v1.ErrorUserAlreadyExists("user already exists"), err)
	assert.Empty(t, token)
}

// Sign up with an invalid password should return ErrPasswordLength
func TestSignUp_InvalidPassword(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockAuthRepo)
	mockLog := new(log.Helper)
	cfg := &conf.Auth{
		Jwt: &conf.Auth_Jwt{
			Secret:     "secret",
			Expiration: durationpb.New(time.Hour),
		},
	}
	uc := &AuthUseCase{cfg: cfg, repo: mockRepo, log: mockLog}
	email := "user@example.com"
	password := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)
	user := &ent.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
		Username: email,
		Status:   true,
	}
	mockRepo.On("FindByEmail", ctx, email).Return(nil, ErrUserNotFound)
	mockRepo.On("Create", ctx, mock.MatchedBy(func(user *ent.User) bool {
		return user.Email == email && user.Username == email
	})).Return(user, nil)

	// Act
	token, err := uc.SignUp(ctx, &v1.SignUpRequest{Email: email, Password: ""})

	// Assert
	assert.Error(t, err)
	assert.Equal(t, helper.ErrPasswordLength, err)
	assert.Empty(t, token)
}
