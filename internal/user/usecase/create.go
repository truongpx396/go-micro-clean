package usecase

import (
	"context"
	"go-micro-clean/internal/user/entity"
)

func (biz *userUsecase) CreateUser(ctx context.Context, data *entity.UserCreation) error {
	return biz.repository.CreateUser(ctx, data)
}
