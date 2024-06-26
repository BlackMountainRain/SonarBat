package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sonar-bat/internal/auth/biz"
	"sonar-bat/internal/runtime"

	pb "sonar-bat/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer

	auth        *biz.AuthUseCase
	runtimeInfo runtime.Info
}

func NewAuthService(auth *biz.AuthUseCase, runtimeInfo runtime.Info) *AuthService {
	return &AuthService{
		auth:        auth,
		runtimeInfo: runtimeInfo,
	}
}

func (s *AuthService) HealthCheck(_ context.Context, _ *pb.HealthRequest) (*pb.HealthReply, error) {
	return &pb.HealthReply{
		Status:  pb.Status_UP,
		Version: s.runtimeInfo.Version,
		Uptime:  timestamppb.New(s.runtimeInfo.Uptime),
	}, nil
}

func (s *AuthService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.AuthReply, error) {
	token, err := s.auth.SignIn(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AuthReply{
		Token: token,
	}, nil
}

func (s *AuthService) SignInWithOAuth(ctx context.Context, req *pb.SignInWithOAuthRequest) (*pb.AuthReply, error) {
	token, err := s.auth.SignInWithOAuth(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AuthReply{
		Token: token,
	}, nil
}

func (s *AuthService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.AuthReply, error) {
	token, err := s.auth.SignUp(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AuthReply{
		Token: token,
	}, nil
}

func (s *AuthService) ValidateJWT(ctx context.Context, req *pb.ValidateJWTRequest) (*pb.UserInfoReply, error) {
	return &pb.UserInfoReply{}, nil
}
func (s *AuthService) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.PermissionReply, error) {
	return &pb.PermissionReply{}, nil
}
func (s *AuthService) RequestPasswordReset(ctx context.Context, req *pb.RequestPasswordResetRequest) (*pb.RequestPasswordResetReply, error) {
	return &pb.RequestPasswordResetReply{}, nil
}
func (s *AuthService) VerifyPasswordResetToken(ctx context.Context, req *pb.VerifyPasswordResetTokenRequest) (*pb.VerifyPasswordResetTokenReply, error) {
	return &pb.VerifyPasswordResetTokenReply{}, nil
}
func (s *AuthService) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordReply, error) {
	return &pb.ResetPasswordReply{}, nil
}
