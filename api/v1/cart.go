package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_mysql/pkg/util"
	"test_mysql/service"
)

func CreateCart(ctx *gin.Context) {
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	createCartService := service.CartService{}
	if err := ctx.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func ListCart(ctx *gin.Context) {
	listCartService := service.CartService{}
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&listCartService); err == nil {
		res := listCartService.List(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func DeleteCart(ctx *gin.Context) {
	listCartService := service.CartService{}
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&listCartService); err == nil {
		res := listCartService.Delete(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func UpdateCart(ctx *gin.Context) {
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	createCartService := service.CartService{}
	if err := ctx.ShouldBind(&createCartService); err == nil {
		res := createCartService.Update(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}
