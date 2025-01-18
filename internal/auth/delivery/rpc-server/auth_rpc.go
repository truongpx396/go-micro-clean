package rpcserver

import (
	"context"

	"project/internal/auth/entity"
	"project/pkg/rpcclient"
	"project/proto/auth"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
)

type AuthUsecase interface {
	Login(ctx context.Context, data *entity.AuthEmailPassword) (*entity.TokenResponse, error)
	Register(ctx context.Context, data *entity.AuthRegister) error
	IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error)
}

type authServer struct {
	business      AuthUsecase
	userRpcClient *rpcclient.UserRpcClient
}

func NewAuthServer(business AuthUsecase, userRpcClient *rpcclient.UserRpcClient) *authServer {
	return &authServer{
		business:      business,
		userRpcClient: userRpcClient,
	}
}

func StartAuthServer(business AuthUsecase, userRpcClient *rpcclient.UserRpcClient, server *grpc.Server) {
	authService := NewAuthServer(business, userRpcClient)
	auth.RegisterAuthServiceServer(server, authService)
}

func (s *authServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	data := &entity.AuthEmailPassword{}
	data.FromProto(req)

	response, err := s.business.Login(ctx, data)
	if err != nil {
		return nil, err
	}

	return response.ToProto(), nil
}

func (s *authServer) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	data := &entity.AuthRegister{}
	data.FromProto(req)

	err := s.business.Register(ctx, data)
	if err != nil {
		return nil, err
	}

	return data.ToProto(), nil
}

func (s *authServer) IntrospectToken(ctx context.Context, req *auth.IntrospectReq) (*auth.IntrospectResp, error) {
	claims, err := s.business.IntrospectToken(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	return &auth.IntrospectResp{
		Tid: claims.ID,
		Sub: claims.Subject,
	}, nil
}
