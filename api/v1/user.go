package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_mysql/pkg/util"
	"test_mysql/service"
)

//api 路由处理函数，处理用户注册请求
//绑定老是出问题，返回400错误

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogRusObj.Infoln("user register err", err)
		util.LogRusObj.Infoln(err)
		return
	} else {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}

}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogRusObj.Infoln(err)
		return
	} else {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}

}

func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogRusObj.Infoln(err)
		return
	} else {
		res := userUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}

}

func UploadAvatar(c *gin.Context) {
	// 上传头像
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatar); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogRusObj.Infoln(err)
		return
	} else {
		res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
		c.JSON(http.StatusOK, res)
	}
}

func SendEmail(c *gin.Context) {
	var sendEmail service.SendEmailService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendEmail); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogRusObj.Infoln(err)
		return
	} else {
		res := sendEmail.Send(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}

}

// ValidEmail 验证邮箱
func ValidEmail(c *gin.Context) {
	var validEmail service.ValidEmailService

	if err := c.ShouldBind(&validEmail); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogRusObj.Infoln(err)
		return
	} else {
		res := validEmail.Valid(c.Request.Context(), c.GetHeader("Authorization"))
		c.JSON(http.StatusOK, res)
	}

}

func ShowMoney(c *gin.Context) {
	// 显示余额
	var showMoney service.ShowMoneyService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoney); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogRusObj.Infoln(err)
		return
	} else {
		res := showMoney.ShowMoney(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}
}
