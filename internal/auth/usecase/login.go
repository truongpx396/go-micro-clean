package usecase

import (
	"context"
	"go-micro-clean/common"
	"go-micro-clean/internal/auth/entity"

	"github.com/google/uuid"
)

func (biz *authUsecase) Login(ctx context.Context, data *entity.AuthEmailPassword) (*entity.TokenResponse, error) {
	if err := data.Validate(); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	authData, err := biz.authRepository.GetAuth(ctx, data.Email)

	if err != nil {
		return nil, common.ErrInvalidRequestWithMsg(err, entity.ErrLoginFailed.Error())
	}

	if !biz.hasher.CompareHashPassword(authData.Password, authData.Salt, data.Password) {
		return nil, common.ErrInvalidRequest(entity.ErrLoginFailed)
	}

	uid := common.NewUID(uint32(authData.UserId), 1, 1)
	sub := uid.String()
	tid := uuid.New().String()

	tokenStr, expSecs, err := biz.jwtProvider.IssueToken(ctx, tid, sub)

	if err != nil {
		return nil, common.ErrInternalWithMsg(err, entity.ErrLoginFailed.Error())
	}

	return &entity.TokenResponse{
		AccessToken: entity.Token{
			Token:     tokenStr,
			ExpiredIn: expSecs,
		},
	}, nil
}
