package postgre

import (
	"context"
	"go-micro-clean/common"
	"go-micro-clean/internal/user/entity"

	"gorm.io/gorm"
)

func (s *userRepository) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*entity.User, error) {
	db := s.db.Table(entity.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user entity.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
