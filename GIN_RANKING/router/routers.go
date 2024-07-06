package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/hello", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Hello World！")
		})
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
	return r
}
