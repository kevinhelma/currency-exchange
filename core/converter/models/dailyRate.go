package models

import "currency-exchange/common"

type Exchange struct {
	common.BaseModel
	Source string `gorm:"type:char(5)"`
	Target string `gorm:"type:char(5)"`
}