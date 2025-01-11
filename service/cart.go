package service

import (
	"context"
	"strconv"
	"test_mysql/dao"
	"test_mysql/model"
	"test_mysql/pkg/e"
	"test_mysql/serializer"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossID    uint `json:"boss_id" form:"boss_id"`
	ProductID uint `json:"product_id" form:"product_id"`
	Num       int  `json:"num" form:"num"`
}

func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	var cart *model.Cart
	code := e.Success

	//判断有没有商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductID)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	cartDao := dao.NewCartDao(ctx)
	cart = &model.Cart{
		UserID:    uId,
		ProductID: service.ProductID,
		BossID:    service.BossID,
		Num:       uint(service.Num),
	}

	err = cartDao.CreateCart(cart)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(service.BossID)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product, boss),
	}

}

func (service *CartService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	carts, err := cartDao.ListCartByUserId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarts(ctx, carts),
	}
}

func (service *CartService) Update(ctx context.Context, uId uint, aId string) serializer.Response {
	code := e.Success
	CartDao := dao.NewCartDao(ctx)
	CartId, _ := strconv.Atoi(aId)
	err := CartDao.UpdateCartNumById(uint(CartId), service.Num)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}

func (service *CartService) Delete(ctx context.Context, uId uint, cId string) serializer.Response {
	cartId, _ := strconv.Atoi(cId)
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	err := cartDao.DeleteCartByCartId(uint(cartId), uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
