package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model `gorm:"not null"`
	UserID     uint `gorm:"not null"`
	ProductID  uint `gorm:"not null"`
	BossID     uint `gorm:"not null"`
	Num        uint `gorm:"not null"`
	MaxNum     uint `gorm:"not null"`
	Check      bool `gorm:"not null"`
}
