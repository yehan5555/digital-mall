package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	api "test_mysql/api/v1"
	"test_mysql/middleware"
)

// NewRouter 路由注册, 使用 gin 框架搭建的 web 路由

func NewRouter() *gin.Engine {

	//建一个应用程序实例

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

		// 商品操作
		v1.GET("products", api.ListProducts)
		v1.GET("products/:id", api.ShowProducts)
		v1.GET("img/:id", api.ListProductImg)
		v1.GET("categories", api.ListCategory)

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

			// 创建商品
			authed.POST("product", api.CreateProduct)
			//搜索商品
			authed.POST("products", api.SearchProducts)

			//收藏夹操作, 有点小问题，待解决
			authed.GET("favorites", api.ListFavorite)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites/:id", api.DeleteFavorite)

			//地址操作
			authed.POST("addresses", api.CreateAddresses)
			authed.GET("addresses/:id", api.ShowAddresses)
			authed.GET("addresses", api.ListAddresses)
			authed.PUT("addresses/:id", api.UpdateAddresses)
			authed.DELETE("addresses/:id", api.DeleteAddresses)

			// 购物车操作
			authed.POST("carts", api.CreateCart)
			authed.GET("carts", api.ListCart)
			authed.PUT("carts/:id", api.UpdateCart)
			authed.DELETE("carts/:id", api.DeleteCart)

			// 订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("orders", api.ListOrder)
			authed.GET("orders/:id", api.ShowOrder)
			authed.DELETE("orders/:id", api.DeleteOrder)

			// 支付功能
			authed.POST("paydown", api.OrderPay)

		}

	}
	return r
}
