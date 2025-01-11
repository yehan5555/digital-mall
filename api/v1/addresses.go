package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_mysql/pkg/util"
	"test_mysql/service"
)

func CreateAddresses(ctx *gin.Context) {
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	createAddressService := service.AddressService{}
	if err := ctx.ShouldBind(&createAddressService); err == nil {
		res := createAddressService.Create(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func ListAddresses(ctx *gin.Context) {
	listAddressService := service.AddressService{}
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&listAddressService); err == nil {
		res := listAddressService.List(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func DeleteAddresses(ctx *gin.Context) {
	listAddressService := service.AddressService{}
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&listAddressService); err == nil {
		res := listAddressService.Delete(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func ShowAddresses(ctx *gin.Context) {
	showAddressService := service.AddressService{}
	if err := ctx.ShouldBind(&showAddressService); err == nil {
		res := showAddressService.Show(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func UpdateAddresses(ctx *gin.Context) {
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	createAddressService := service.AddressService{}
	if err := ctx.ShouldBind(&createAddressService); err == nil {
		res := createAddressService.Update(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}
