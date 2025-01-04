package common

import (
	"time"

	"gorm.io/gorm"
)

type SQLModelCreate struct {
	Id        uint      `json:"-" gorm:"column:id;primaryKey"`
	Status    int       `json:"status" gorm:"column:status;"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at;"`
	UpdateAt  time.Time `json:"-" gorm:"column:updated_at;"`
}

func (sqlModel *SQLModelCreate) PrepareForInsert() {
	prepareForInsert(&sqlModel.Id, &sqlModel.Status, &sqlModel.CreatedAt, &sqlModel.UpdateAt)
}

type SQLModel struct {
	gorm.Model
	Id        uint       `json:"id" gorm:"column:id;primaryKey"`
	Status    int        `json:"status" gorm:"column:status;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

func (sqlModel *SQLModel) PrepareForInsert() {
	prepareForInsert(&sqlModel.Id, &sqlModel.Status, &sqlModel.CreatedAt, &sqlModel.UpdatedAt)
}

func prepareForInsert(id *uint, status *int, createdAt *time.Time, updatedAt *time.Time) {
	now := time.Now().UTC()
	*id = 0
	*status = 1
	*createdAt = now
	*updatedAt = now
}
