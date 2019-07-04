package models

import (
	"currency-exchange/common"
	"time"
	"github.com/satori/go.uuid"
)

type DailyRate struct {
	common.BaseModel
	ExchangeID   uuid.UUID `gorm:"type:char(36); not null"`
	ExchangeDate *time.Time
	Rate         float64 `gorm:"decimal"`
}