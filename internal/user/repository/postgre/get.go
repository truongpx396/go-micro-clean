package postgre

import (
	"context"
	"go-micro-clean/common"
	"go-micro-clean/internal/user/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *userRepository) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	var data entity.User

	if err := repo.db.
		Table(data.TableName()).
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrNotFound
		}
		return nil, errors.Wrap(err, entity.ErrCannotGetUser.Error())
	}

	return &data, nil
}

func (s *userRepository) GetUsers(ctx context.Context, ids []int) ([]entity.SimpleUser, error) {
	var result []entity.SimpleUser

	if err := s.db.Table(entity.SimpleUser{}.TableName()).
		Where("id in (?)", ids).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
