package controllers

import "github.com/gin-gonic/gin"

type OrderController struct {
}

type Search struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (o *OrderController) GetOrderInfo(ctx *gin.Context) {
	ReturnSuccess(ctx, 200, "success", "Hello World Order.", 1)
}

func (o *OrderController) GetOrderList(ctx *gin.Context) {
	ReturnError(ctx, 400, "msg get failure : GetOrderList")
}

func (o *OrderController) PostOrderInfo(ctx *gin.Context) {
	// 使用普通格式进行处理
	//id := ctx.PostForm("id")
	//name := ctx.DefaultPostForm("name", "wanger")
	//ReturnSuccess(ctx, 200, id, name, 1)

	// 使用JSON格式进行调用传递
	//param := make(map[string]interface{})
	//err := ctx.BindJSON(&param)
	//if err == nil {
	//	ReturnSuccess(ctx, 200, param["id"], param["name"], 1)
	//	return
	//}
	//ReturnError(ctx, 400, "Failure.")

	// 使用struct结构体进行接收
	search := &Search{}
	err := ctx.BindJSON(&search)
	if err == nil {
		ReturnSuccess(ctx, 200, search.Id, search.Name, 1)
		return
	}
	ReturnError(ctx, 400, "Failure")
}
