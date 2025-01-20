package postgre

import (
	"context"
	"go-micro-clean/common"
	"go-micro-clean/internal/auth/entity"

	"gorm.io/gorm"
)

type postgreRepo struct {
	db *gorm.DB
}

func NewPostgreRepository(db *gorm.DB) *postgreRepo {
	return &postgreRepo{db: db}
}

func (repo *postgreRepo) AddNewAuth(ctx context.Context, data *entity.Auth) error {
	if err := repo.db.Table(data.TableName()).Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (repo *postgreRepo) GetAuth(ctx context.Context, email string) (*entity.Auth, error) {
	var data entity.Auth

	if err := repo.db.
		Table(data.TableName()).
		Where("email = ?", email).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return &data, nil
}
