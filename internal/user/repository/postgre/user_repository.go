package postgre

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

func NewPostgreRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
