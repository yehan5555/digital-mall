package dao

import (
	"context"
	"gorm.io/gorm"
	"test_mysql/model"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func (dao *CartDao) CreateCart(in *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&in).Error
}

func (dao *CartDao) GetCartByAid(aId uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id = ?", aId).First(&cart, aId).Error
	return
}

func (dao *CartDao) ListCartByUserId(uId uint) (carts []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id = ?", uId).Find(&carts).Error
	return
}

func (dao *CartDao) UpdateCartById(aId uint, cart *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Where("id = ?", aId).Updates(&cart).Error
}

func (dao *CartDao) DeleteCartByCartId(aId uint, uId uint) error {
	return dao.DB.Model(&model.Cart{}).
		Where("id = ? AND user_id = ?", aId, uId).
		Delete(&model.Cart{}).Error

}

func (dao *CartDao) UpdateCartNumById(cId uint, num int) error {
	return dao.DB.Model(&model.Cart{}).
		Where("id = ?", cId).
		Update("num", num).Error
}
