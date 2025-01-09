package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_mysql/pkg/util"
	"test_mysql/service"
)

//创建商品

func CreateProduct(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	if err := ctx.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(ctx.Request.Context(), claims.ID, files)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}
