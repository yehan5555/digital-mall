package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_mysql/pkg/util"
	"test_mysql/service"
)

//创建商品
// ctx 来请求数据，设置响应

func CreateProduct(ctx *gin.Context) {
	//获取上传的文件，multipart/form-data
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	//解析token
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

func ListProducts(ctx *gin.Context) {
	listProductService := service.ProductService{}
	if err := ctx.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func ShowProducts(ctx *gin.Context) {
	showProductService := service.ProductService{}
	if err := ctx.ShouldBind(&showProductService); err == nil {
		res := showProductService.Show(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
		util.LogRusObj.Infoln(err)
	}
}

func SearchProducts(ctx *gin.Context) {
	searchProductService := service.ProductService{}
	if err := ctx.ShouldBind(&searchProductService); err == nil {
		res := searchProductService.Search(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogRusObj.Infoln(err)
	}
}
