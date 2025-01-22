package rpcserver

import (
	"context"

	"go-micro-clean/common"
	"go-micro-clean/config"
	"go-micro-clean/internal/auth/entity"
	"go-micro-clean/internal/auth/repository/postgre"
	"go-micro-clean/internal/auth/usecase"
	"go-micro-clean/pkg/rpcclient"
	"go-micro-clean/proto/auth"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
)

type authUsecase interface {
	Login(ctx context.Context, data *entity.AuthEmailPassword) (*entity.TokenResponse, error)
	Register(ctx context.Context, data *entity.AuthRegister) error
	IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error)
}

type authServer struct {
	business authUsecase
}

func StartAuthServer(ctx context.Context, server *grpc.Server) error {

	db := config.SetupDatabase()

	jwtComp := common.NewJWT("jwt")

	authRepo := postgre.NewPostgreRepository(db)
	hasher := new(common.Hasher)

	userRepo := rpcclient.NewUser(ctx)
	business := usecase.NewAuthUsecase(authRepo, userRepo, jwtComp, hasher)

	authService := &authServer{business}
	auth.RegisterAuthServiceServer(server, authService)
	return nil
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
