package usecase

import (
	"context"
	"project/internal/user/entity"
)

func (biz *userUsecase) CreateUser(ctx context.Context, data *entity.UserCreation) error {
	return biz.repository.CreateUser(ctx, data)
}
