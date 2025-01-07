package usecase

import (
	"context"
	"project/common"
	"project/modules/auth/entity"

	"github.com/golang-jwt/jwt/v5"
)

type AuthRepository interface {
	AddNewAuth(ctx context.Context, data *entity.Auth) error
	GetAuth(ctx context.Context, email string) (*entity.Auth, error)
}

type UserRepository interface {
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
	userRepository UserRepository
	jwtProvider    JWTProviderComponent
	hasher         Hasher
}

func NewAuthUsecase(repository AuthRepository, userRepository UserRepository,
	jwtProvider JWTProviderComponent, hasher Hasher) *authUsecase {
	return &authUsecase{
		authRepository: repository,
		userRepository: userRepository,
		jwtProvider:    jwtProvider,
		hasher:         hasher,
	}
}

func (biz *authUsecase) IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error) {
	claims, err := biz.jwtProvider.ParseToken(ctx, accessToken)

	if err != nil {
		return nil, common.ErrUnauthorized(err)
	}

	return claims, nil
}
