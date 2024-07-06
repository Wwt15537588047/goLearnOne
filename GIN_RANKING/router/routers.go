package router

import (
	"gin_ranking/controllers"
	"gin_ranking/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	user := r.Group("/user")
	{
		/*		user.GET("/hello", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "Hello World！")
			})*/

		user.GET("/hello/:id", (&controllers.UserController{}).GetUserInfo)
		/**
		在Go语言中
		当你将一个方法绑定到一个路由处理器时，你实际上需要提供一个方法的引用，而不是调用这个方法。
		在 gin 框架中，路由处理器需要一个函数作为其回调，这个函数将在请求到达时被调用。
		*/
		user.GET("/list", (&controllers.UserController{}).GetUserList)

		user.POST("/add", (&controllers.UserController{}).AddUser)
		user.POST("/update", (&controllers.UserController{}).UpdateUser)
		user.POST("/delete", (&controllers.UserController{}).DeleteUser)

		user.POST("/post", func(context *gin.Context) {
			context.String(http.StatusOK, "User post!")
		})

		user.PUT("/put", func(context *gin.Context) {
			context.String(http.StatusOK, "User put!")
		})

		user.DELETE("/delete", func(context *gin.Context) {
			context.String(http.StatusOK, "User delete!")
		})
	}

	order := r.Group("/order")
	{
		order.GET("/hello", (&controllers.OrderController{}).GetOrderInfo)
		order.GET("/list", (&controllers.OrderController{}).GetOrderList)

		order.POST("/post", (&controllers.OrderController{}).PostOrderInfo)
	}
	//logger.LoggerToFile()
	return r
}
