package usecase

import (
	"context"
	"project/internal/user/entity"
)

type UserRepository interface {
	GetUsers(ctx context.Context, ids []int) ([]entity.SimpleUser, error)
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*entity.User, error)
	CreateUser(ctx context.Context, data *entity.UserCreation) error
}

type userUsecase struct {
	repository UserRepository
}

func NewBusiness(repository UserRepository) *userUsecase {
	return &userUsecase{repository: repository}
}
