package postgre

import (
	"context"
	"go-micro-clean/common"
	"go-micro-clean/internal/user/entity"
)

func (s *userRepository) CreateUser(ctx context.Context, data *entity.UserCreation) error {
	db := s.db.Begin()
	data.PrepareForInsert()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
