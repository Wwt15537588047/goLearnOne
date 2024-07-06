package controllers

import (
	"gin_ranking/models"
	"gin_ranking/pkg/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
}

func (c UserController) GetUserInfo(ctx *gin.Context) {
	logger.Write("调用Hello", "user")
	idStr := ctx.Param("id")
	name := ctx.Param("name")
	id, _ := strconv.Atoi(idStr)

	user, _ := models.GetUserTest(id)
	ReturnSuccess(ctx, 200, name, user, 1)
}

func (u *UserController) GetUserList(ctx *gin.Context) {
	logger.Write("日志信息", "user")

	// 已经添加全局异常处理器-日志类进行panic异常捕获，不需要函数内部单独处理
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("程序崩溃，错误是：", err)
	//		ReturnError(ctx, 400, "msg get failure : GetUserList")
	//	}
	//}()

	num1 := 1
	num2 := 0
	num3 := num1 / num2

	ReturnSuccess(ctx, 200, num3, "Success", 1)
}

func (u *UserController) AddUser(ctx *gin.Context) {
	logger.Write("日志信息", "addUser")
	username := ctx.DefaultPostForm("username", "")
	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(ctx, 404, "插入数据库失败")
		return
	}
	ReturnSuccess(ctx, 200, id, "插入成功", 1)

}

func (U *UserController) UpdateUser(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	username := ctx.PostForm("username")
	models.UpdateUser(id, username)
	ReturnSuccess(ctx, 200, "username", "success", 1)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(ctx, 404, "删除数据库失败")
		return
	}
	ReturnSuccess(ctx, 200, "delete", "success", 1)
}
