package dao

import (
	"context"
	"gorm.io/gorm"
	"test_mysql/model"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoById(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) error {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

func (dao *ProductImgDao) ListProductImg(id uint) (productImg []*model.ProductImg, err error) {
	err = dao.DB.Model(model.ProductImg{}).Where("product_id = ?", id).Find(&productImg).Error
	return
}
