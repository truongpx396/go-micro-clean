package common

import (
	"time"
)

type SQLModelCreate struct {
	ID        uint      `json:"-" gorm:"column:id;primaryKey"`
	Status    int       `json:"status" gorm:"column:status;"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at;"`
	UpdateAt  time.Time `json:"-" gorm:"column:updated_at;"`
}

func (sqlModel *SQLModelCreate) PrepareForInsert() {
	prepareForInsert(&sqlModel.ID, &sqlModel.Status, &sqlModel.CreatedAt, &sqlModel.UpdateAt)
}

type SQLModel struct {
	ID        uint       `json:"id" gorm:"column:id;primaryKey"`
	Status    int        `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;default:NULL"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;default:NULL"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index;default:NULL"`
}

func (sqlModel *SQLModel) PrepareForInsert() {
	prepareForInsert(&sqlModel.ID, &sqlModel.Status, sqlModel.CreatedAt, sqlModel.UpdatedAt)
}

func (sqlModel *SQLModel) BeforeUpdate() {
	now := time.Now().UTC()
	if sqlModel.CreatedAt == nil {
		sqlModel.CreatedAt = &now
	}
	// if sqlModel.UpdatedAt == nil {
	sqlModel.UpdatedAt = &now
	// }
}

func prepareForInsert(id *uint, status *int, createdAt *time.Time, updatedAt *time.Time) {
	now := time.Now().UTC()
	*id = 0
	*status = 1
	*createdAt = now
	*updatedAt = now
}
