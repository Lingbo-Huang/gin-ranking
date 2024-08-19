package controllers

import "github.com/gin-gonic/gin"

/*
把公共的文件放在controllers中
和业务相关的也都会在controllers包下，这样就不用引用包了
*/

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"` // 总数，一般从mysql查询出来都是int64类型
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JsonStruct{
		Code: code,
		Msg:  msg,
	}
	c.JSON(404, json)
}
