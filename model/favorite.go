package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	//会影响性能，不建议使用
	User      User `gorm:"ForeignKey:UserID"`
	UserID    uint `gorm:"not null"`
	Product   uint `gorm:"ForeignKey:ProductID"`
	ProductId uint `gorm:"not null"`
	Boss      User `gorm:"ForeignKey:BossID"`
	BossID    uint `gorm:"not null"`
}
