package http

import (
	"context"
	"go-micro-clean/internal/auth/entity"
)

type AuthUsecase interface {
	Login(ctx context.Context, data *entity.AuthEmailPassword) (*entity.TokenResponse, error)
	Register(ctx context.Context, data *entity.AuthRegister) error
}

type authHandler struct {
	ctx      context.Context
	business AuthUsecase
}

func NewAuthHandler(ctx context.Context, business AuthUsecase) *authHandler {
	return &authHandler{ctx: ctx, business: business}
}
