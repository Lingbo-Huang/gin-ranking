package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

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

func EncryMd5(s string) string {
	// 初始化一个MD5哈希对象
	ctx := md5.New()
	// 将字符串写入ctx 对象的内部缓冲区
	ctx.Write([]byte(s))
	// 计算MD5哈希值, 并将结果转换为十六进制编码的字符串
	return hex.EncodeToString(ctx.Sum(nil))
}
