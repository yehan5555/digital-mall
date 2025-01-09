package model

import "gorm.io/gorm"

// 轮播图，为了宣传

type Carousel struct {
	gorm.Model
	ImgPath   string `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
}
