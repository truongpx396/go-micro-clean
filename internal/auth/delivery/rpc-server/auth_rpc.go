package rpcserver

import (
	"context"

	"project/proto/auth"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase interface {
	IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error)
}

type grpcService struct {
	business AuthUsecase
}

func NewService(business AuthUsecase) *grpcService {
	return &grpcService{business: business}
}

func (s *grpcService) IntrospectToken(ctx context.Context, req *auth.IntrospectReq) (*auth.IntrospectResp, error) {
	claims, err := s.business.IntrospectToken(ctx, req.AccessToken)

	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }

	if err != nil {
		return nil, err
	}

	return &auth.IntrospectResp{
		Tid: claims.ID,
		Sub: claims.Subject,
	}, nil
}
