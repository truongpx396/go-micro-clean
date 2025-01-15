package usecase

import (
	"context"
	"project/common"

	"github.com/golang-jwt/jwt/v5"
)

func (biz *authUsecase) IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error) {
	claims, err := biz.jwtProvider.ParseToken(ctx, accessToken)

	if err != nil {
		return nil, common.ErrUnauthorized(err)
	}

	return claims, nil
}
