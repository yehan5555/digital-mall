package dao

import (
	"context"
	"gorm.io/gorm"
	"test_mysql/model"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Create(product).Error
}
