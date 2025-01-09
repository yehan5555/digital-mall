package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	api "test_mysql/api/v1"
	"test_mysql/middleware"
)

// NewRouter 路由注册, 使用 gin 框架搭建的 web 路由

func NewRouter() *gin.Engine {

	//建一个WSGI应用程序实例

	r := gin.Default()

	//使用中间件来实现跨域
	r.Use(middleware.Cors())
	//加载静态文件,
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/api/v1")
	{
		//测试接口
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		//处理用户注册请求
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		//轮播图
		v1.GET("carousels", api.ListCarousels)

		// 验证 token 是 用户的 token
		// 需要登录保护
		//设置登录验证的路由组
		authed := v1.Group("/")
		//为 路由组添加中间件，JWT验证, 确保访问该路由下的所有路由都需要提供有效的 jwt 令牌，验证身份
		authed.Use(middleware.JWT())
		{
			//用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			// 显示金额
			authed.POST("money", api.ShowMoney)

			authed.POST("product", api.CreateProduct)
		}

	}
	return r
}
