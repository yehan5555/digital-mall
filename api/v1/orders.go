package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_mysql/pkg/util"
	"test_mysql/service"
)

func CreateOrder(ctx *gin.Context) {
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	createOrderService := service.OrderService{}
	if err := ctx.ShouldBind(&createOrderService); err == nil {
		res := createOrderService.Create(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func ListOrder(ctx *gin.Context) {
	listOrderService := service.OrderService{}
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&listOrderService); err == nil {
		res := listOrderService.List(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func DeleteOrder(ctx *gin.Context) {
	listOrderService := service.OrderService{}
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&listOrderService); err == nil {
		res := listOrderService.Delete(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func ShowOrder(ctx *gin.Context) {
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	showOrderService := service.OrderService{}
	if err := ctx.ShouldBind(&showOrderService); err == nil {
		res := showOrderService.Show(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}
