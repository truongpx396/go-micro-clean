package usecase

import (
	"context"
	"project/common"
	"project/modules/user/entity"
)

func (biz *userUsecase) GetUserProfile(ctx context.Context, requesterId int) (*entity.User, error) {

	user, err := biz.repository.FindUser(ctx, map[string]interface{}{"id": requesterId})

	if err != nil {
		return nil, common.NewUnauthorized(err, entity.ErrCannotGetUser.Error(), "CAN_NOT_GET_USER")
	}

	return user, nil
}

func (biz *userUsecase) GetUsers(ctx context.Context, ids []int) ([]entity.SimpleUser, error) {
	return biz.repository.GetUsers(ctx, ids)
}

func (biz *userUsecase) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*entity.User, error) {
	return biz.repository.FindUser(ctx, conditions, moreInfo...)
}
