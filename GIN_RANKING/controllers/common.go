package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义返回的数据格式Json格式
type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type JsonStructFailure struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func ReturnSuccess(ctx *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	ctx.JSON(http.StatusOK, json)
}

func ReturnError(ctx *gin.Context, code int, msg interface{}) {
	json := &JsonStructFailure{
		Code: code,
		Msg:  msg,
	}
	ctx.JSON(http.StatusOK, json)
}
