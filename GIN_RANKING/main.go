package main

import (
	"gin_ranking/router"
)

func main() {
	//fmt.Println("Hello World ...")
	//r := gin.Default()
	r := router.Router()
	/*r.GET("hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World！")
	})

	r.POST("/user/list", func(context *gin.Context) {
		context.String(http.StatusOK, "User post!")
	})

	r.PUT("/user/put", func(context *gin.Context) {
		context.String(http.StatusOK, "User put!")
	})

	r.DELETE("/user/delete", func(context *gin.Context) {
		context.String(http.StatusOK, "User delete!")
	})*/

	r.Run(":9999")
}
