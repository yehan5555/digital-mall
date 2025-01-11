package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"test_mysql/dao"
	"test_mysql/model"
	"test_mysql/pkg/e"
	"test_mysql/pkg/util"
	"test_mysql/serializer"
)

type OrderPay struct {
	OrderId   uint   `json:"order_id" form:"order_id"`
	Money     int    `json:"money" form:"money"`
	OrderNo   string `json:"order_no" form:"order_no"`
	ProductId uint   `json:"product_id" form:"product_id"`
	PayTime   string `json:"pay_time" form:"pay_time"`
	Sign      string `json:"sign" form:"sign"`
	BossId    uint   `json:"boss_id" form:"boss_id"`
	BossName  string `json:"boss_name" form:"boss_name"`
	Num       int    `json:"num" form:"num"`
	Key       string `json:"key" form:"key"`
}

func (service *OrderPay) PayDown(ctx context.Context, uId uint) serializer.Response {
	util.Encrypt.SetKey(service.Key)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	//使用事务处理，一旦发生错误，事务会自动回滚
	tx := orderDao.Begin()
	order, err := orderDao.GetOrderById(service.OrderId, uId)
	if err != nil {
		util.LogRusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	money := order.Money
	num := order.Num
	money = money * float64(num)

	//  用户扣钱
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		util.LogRusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 对钱进行解密，减去订单，再加密保存
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)

	if moneyFloat-money < 0 {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("余额不足").Error(),
		}
	}
	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(finMoney)

	userDao = dao.NewUserDaoByDB(userDao.DB)
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 商家加钱
	var boss *model.User
	boss, err = userDao.GetUserById(service.BossId)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	moneyStr = util.Encrypt.AesDecoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finMoney)
	err = userDao.UpdateUserById(service.BossId, boss)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = userDao.UpdateUserById(boss.ID, boss)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 对应的商品数目减一
	var product *model.Product
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(service.ProductId)

	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if product == nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "product not found",
		}
	}
	product.Num -= num

	err = productDao.UpdateProduct(service.ProductId, product)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 订单删除, 这里似乎有点问题，需要看 github 上的源码
	err = orderDao.DeleteOrderByOrderId(service.OrderId, uId)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 自己的商品加1 ， 同一件商品就会有点问题，需要看 github 上的源码
	productUser := model.Product{

		Name:          product.Name,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		Price:         product.Price,
		DiscountPrice: product.Price,
		OnSale:        false,
		Num:           1,
		BossId:        uId,
		BossName:      user.Username,
		BossAvatar:    user.Avatar,
	}
	err = productDao.CreateProduct(&productUser)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	tx.Commit()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
