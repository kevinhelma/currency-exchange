package common

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:char(36); primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

//BeforeCreate : Gorm callback to generate ID before create model
func (b *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	if b.ID == uuid.Nil {
		return scope.SetColumn("ID", interface{}(uuid.NewV4()))
	}

	return nil
}
