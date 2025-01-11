package dao

import (
	"context"
	"gorm.io/gorm"
	"test_mysql/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// 根据  username 查询用户是否存在

func (dao *UserDao) ExistOrNotByUsername(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("username =?", userName).Count(&count).Find(&user).Error
	if count == 0 {

		return nil, false, err
	}
	return user, true, nil
}

//插入数据

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(user).Error
}

// GetUserById 根据id查询用户信息
func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id =?", id).First(&user).Error
	return
}

// UpdateUser 更新用户信息

func (dao *UserDao) UpdateUserById(uId uint, user *model.User) error {

	return dao.DB.Model(&model.User{}).Where("id =?", uId).Updates(user).Error
}
