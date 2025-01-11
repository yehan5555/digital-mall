package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_mysql/pkg/util"
	"test_mysql/service"
)

func CreateFavorite(ctx *gin.Context) {
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	createFavoriteService := service.FavoriteService{}
	if err := ctx.ShouldBind(&createFavoriteService); err == nil {
		res := createFavoriteService.Create(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func ListFavorite(ctx *gin.Context) {
	listFavoriteService := service.FavoriteService{}
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&listFavoriteService); err == nil {
		res := listFavoriteService.List(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func DeleteFavorite(ctx *gin.Context) {
	listFavoriteService := service.FavoriteService{}
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&listFavoriteService); err == nil {
		res := listFavoriteService.Delete(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}
