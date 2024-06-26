// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.21.12
// source: auth/v1/auth.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAuthCheckPermission = "/api.auth.v1.Auth/CheckPermission"
const OperationAuthHealthCheck = "/api.auth.v1.Auth/HealthCheck"
const OperationAuthRequestPasswordReset = "/api.auth.v1.Auth/RequestPasswordReset"
const OperationAuthResetPassword = "/api.auth.v1.Auth/ResetPassword"
const OperationAuthSignIn = "/api.auth.v1.Auth/SignIn"
const OperationAuthSignInWithOAuth = "/api.auth.v1.Auth/SignInWithOAuth"
const OperationAuthSignUp = "/api.auth.v1.Auth/SignUp"
const OperationAuthValidateJWT = "/api.auth.v1.Auth/ValidateJWT"
const OperationAuthVerifyPasswordResetToken = "/api.auth.v1.Auth/VerifyPasswordResetToken"

type AuthHTTPServer interface {
	CheckPermission(context.Context, *CheckPermissionRequest) (*PermissionReply, error)
	HealthCheck(context.Context, *HealthRequest) (*HealthReply, error)
	RequestPasswordReset(context.Context, *RequestPasswordResetRequest) (*RequestPasswordResetReply, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordReply, error)
	SignIn(context.Context, *SignInRequest) (*AuthReply, error)
	SignInWithOAuth(context.Context, *SignInWithOAuthRequest) (*AuthReply, error)
	SignUp(context.Context, *SignUpRequest) (*AuthReply, error)
	ValidateJWT(context.Context, *ValidateJWTRequest) (*UserInfoReply, error)
	VerifyPasswordResetToken(context.Context, *VerifyPasswordResetTokenRequest) (*VerifyPasswordResetTokenReply, error)
}

func RegisterAuthHTTPServer(s *http.Server, srv AuthHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/auth/health", _Auth_HealthCheck0_HTTP_Handler(srv))
	r.POST("/api/v1/auth/sign-in/password", _Auth_SignIn0_HTTP_Handler(srv))
	r.POST("/api/v1/auth/sign-in/oauth", _Auth_SignInWithOAuth0_HTTP_Handler(srv))
	r.POST("/api/v1/auth/sign-up", _Auth_SignUp0_HTTP_Handler(srv))
	r.POST("/api/v1/auth/validate-jwt", _Auth_ValidateJWT0_HTTP_Handler(srv))
	r.POST("/api/v1/auth/check-permission", _Auth_CheckPermission0_HTTP_Handler(srv))
	r.POST("/api/v1/auth/request-password-reset", _Auth_RequestPasswordReset0_HTTP_Handler(srv))
	r.GET("/api/v1/auth/verify-password-reset-token", _Auth_VerifyPasswordResetToken0_HTTP_Handler(srv))
	r.POST("/api/v1/auth/reset-password", _Auth_ResetPassword0_HTTP_Handler(srv))
}

func _Auth_HealthCheck0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HealthRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthHealthCheck)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.HealthCheck(ctx, req.(*HealthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HealthReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_SignIn0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SignInRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthSignIn)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SignIn(ctx, req.(*SignInRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_SignInWithOAuth0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SignInWithOAuthRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthSignInWithOAuth)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SignInWithOAuth(ctx, req.(*SignInWithOAuthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_SignUp0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SignUpRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthSignUp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SignUp(ctx, req.(*SignUpRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_ValidateJWT0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ValidateJWTRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthValidateJWT)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ValidateJWT(ctx, req.(*ValidateJWTRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_CheckPermission0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CheckPermissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthCheckPermission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CheckPermission(ctx, req.(*CheckPermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_RequestPasswordReset0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RequestPasswordResetRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthRequestPasswordReset)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RequestPasswordReset(ctx, req.(*RequestPasswordResetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RequestPasswordResetReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_VerifyPasswordResetToken0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VerifyPasswordResetTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthVerifyPasswordResetToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VerifyPasswordResetToken(ctx, req.(*VerifyPasswordResetTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VerifyPasswordResetTokenReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_ResetPassword0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResetPasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthResetPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ResetPassword(ctx, req.(*ResetPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ResetPasswordReply)
		return ctx.Result(200, reply)
	}
}

type AuthHTTPClient interface {
	CheckPermission(ctx context.Context, req *CheckPermissionRequest, opts ...http.CallOption) (rsp *PermissionReply, err error)
	HealthCheck(ctx context.Context, req *HealthRequest, opts ...http.CallOption) (rsp *HealthReply, err error)
	RequestPasswordReset(ctx context.Context, req *RequestPasswordResetRequest, opts ...http.CallOption) (rsp *RequestPasswordResetReply, err error)
	ResetPassword(ctx context.Context, req *ResetPasswordRequest, opts ...http.CallOption) (rsp *ResetPasswordReply, err error)
	SignIn(ctx context.Context, req *SignInRequest, opts ...http.CallOption) (rsp *AuthReply, err error)
	SignInWithOAuth(ctx context.Context, req *SignInWithOAuthRequest, opts ...http.CallOption) (rsp *AuthReply, err error)
	SignUp(ctx context.Context, req *SignUpRequest, opts ...http.CallOption) (rsp *AuthReply, err error)
	ValidateJWT(ctx context.Context, req *ValidateJWTRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	VerifyPasswordResetToken(ctx context.Context, req *VerifyPasswordResetTokenRequest, opts ...http.CallOption) (rsp *VerifyPasswordResetTokenReply, err error)
}

type AuthHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthHTTPClient(client *http.Client) AuthHTTPClient {
	return &AuthHTTPClientImpl{client}
}

func (c *AuthHTTPClientImpl) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...http.CallOption) (*PermissionReply, error) {
	var out PermissionReply
	pattern := "/api/v1/auth/check-permission"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthCheckPermission))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) HealthCheck(ctx context.Context, in *HealthRequest, opts ...http.CallOption) (*HealthReply, error) {
	var out HealthReply
	pattern := "/api/v1/auth/health"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthHealthCheck))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) RequestPasswordReset(ctx context.Context, in *RequestPasswordResetRequest, opts ...http.CallOption) (*RequestPasswordResetReply, error) {
	var out RequestPasswordResetReply
	pattern := "/api/v1/auth/request-password-reset"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthRequestPasswordReset))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...http.CallOption) (*ResetPasswordReply, error) {
	var out ResetPasswordReply
	pattern := "/api/v1/auth/reset-password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthResetPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) SignIn(ctx context.Context, in *SignInRequest, opts ...http.CallOption) (*AuthReply, error) {
	var out AuthReply
	pattern := "/api/v1/auth/sign-in/password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthSignIn))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) SignInWithOAuth(ctx context.Context, in *SignInWithOAuthRequest, opts ...http.CallOption) (*AuthReply, error) {
	var out AuthReply
	pattern := "/api/v1/auth/sign-in/oauth"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthSignInWithOAuth))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) SignUp(ctx context.Context, in *SignUpRequest, opts ...http.CallOption) (*AuthReply, error) {
	var out AuthReply
	pattern := "/api/v1/auth/sign-up"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthSignUp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) ValidateJWT(ctx context.Context, in *ValidateJWTRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/api/v1/auth/validate-jwt"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthValidateJWT))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) VerifyPasswordResetToken(ctx context.Context, in *VerifyPasswordResetTokenRequest, opts ...http.CallOption) (*VerifyPasswordResetTokenReply, error) {
	var out VerifyPasswordResetTokenReply
	pattern := "/api/v1/auth/verify-password-reset-token"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthVerifyPasswordResetToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
