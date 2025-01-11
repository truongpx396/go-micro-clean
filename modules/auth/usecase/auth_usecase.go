package usecase

import (
	"context"
	"project/modules/auth/entity"

	"github.com/golang-jwt/jwt/v5"
)

type AuthRepository interface {
	AddNewAuth(ctx context.Context, data *entity.Auth) error
	GetAuth(ctx context.Context, email string) (*entity.Auth, error)
}

type UserRpcClient interface {
	CreateUser(ctx context.Context, firstName, lastName, email, avatar string) (newId int, err error)
}

type Hasher interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(hashedPassword, salt, password string) bool
}

type JWTProviderComponent interface {
	IssueToken(ctx context.Context, id, sub string) (token string, expSecs int, err error)
	ParseToken(ctx context.Context, tokenString string) (claims *jwt.RegisteredClaims, err error)
}

type authUsecase struct {
	authRepository AuthRepository
	userRepository UserRpcClient
	jwtProvider    JWTProviderComponent
	hasher         Hasher
}

func NewAuthUsecase(repository AuthRepository, userRepository UserRpcClient,
	jwtProvider JWTProviderComponent, hasher Hasher) *authUsecase {
	return &authUsecase{
		authRepository: repository,
		userRepository: userRepository,
		jwtProvider:    jwtProvider,
		hasher:         hasher,
	}
}
