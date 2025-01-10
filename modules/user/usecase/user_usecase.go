package usecase

import (
	"context"
	"project/common"
	"project/modules/user/entity"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int) (*entity.User, error)
	GetUsersByIds(ctx context.Context, ids []int) ([]entity.User, error)
	CreateNewUser(ctx context.Context, data *entity.UserCreation) error
}

type userUsecase struct {
	repository UserRepository
}

func NewBusiness(repository UserRepository) *userUsecase {
	return &userUsecase{repository: repository}
}

func (biz *userUsecase) GetUserProfile(ctx context.Context, requesterId int) (*entity.User, error) {

	user, err := biz.repository.GetUserById(ctx, requesterId)

	if err != nil {
		return nil, common.NewUnauthorized(err, entity.ErrCannotGetUser.Error(), "CAN_NOT_GET_USER")
	}

	return user, nil
}
