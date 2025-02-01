package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"go-micro-clean/common"
	"go-micro-clean/internal/auth/entity"

	"github.com/btcsuite/btcutil/base58"
)

func (biz *authUsecase) Register(ctx context.Context, data *entity.AuthRegister) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	_, err := biz.authRepository.GetAuth(ctx, data.Email)

	if err == nil {
		return common.ErrInvalidRequest(entity.ErrEmailHasExisted)
	} else if err != common.ErrRecordNotFound {
		return common.ErrInternalWithMsg(err, entity.ErrCannotRegister.Error())
	}

	avatar, err := json.Marshal(data.Avatar)
	if err != nil {
		return common.ErrInternalWithMsg(err, entity.ErrCannotRegister.Error())
	}
	avatarData := base58.Encode(avatar)

	newUserId, err := biz.userRepository.CreateUser(ctx, data.FirstName, data.LastName, data.Email, avatarData)

	if err != nil {
		return common.ErrInternalWithMsg(fmt.Errorf("[userRpc error] %v", err), entity.ErrCannotRegister.Error())
	}

	salt, err := biz.hasher.RandomStr(16)

	if err != nil {
		return common.ErrInternalWithMsg(err, entity.ErrCannotRegister.Error())
	}

	passHashed, err := biz.hasher.HashPassword(salt, data.Password)

	if err != nil {
		return common.ErrInternalWithMsg(err, entity.ErrCannotRegister.Error())
	}

	newAuth := entity.NewAuthWithEmailPassword(newUserId, data.Email, salt, passHashed)

	if err := biz.authRepository.AddNewAuth(ctx, &newAuth); err != nil {
		return common.ErrInternalWithMsg(err, entity.ErrCannotRegister.Error())
	}

	return nil
}
