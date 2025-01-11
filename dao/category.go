package dao

import (
	"context"
	"gorm.io/gorm"
	"test_mysql/model"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

func (dao *CategoryDao) ListCategory() (Category []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&Category).Error
	return
}
