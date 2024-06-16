package service

import (
	"context"
	"sonar-bat/internal/auth/biz"

	pb "sonar-bat/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer

	auth *biz.AuthUseCase
}

func NewAuthService(auth *biz.AuthUseCase) *AuthService {
	return &AuthService{
		auth: auth,
	}
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
	return &pb.AuthReply{}, nil
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
func (s *AuthService) GetSupportedOAuthProviders(ctx context.Context, req *pb.EmptyRequest) (*pb.SupportedOAuthProvidersReply, error) {
	return &pb.SupportedOAuthProvidersReply{}, nil
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
