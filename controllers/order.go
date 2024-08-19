package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

// Search 定义一个结构体，用于接收 POST 请求体中的 JSON 数据
type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (o OrderController) GetList(c *gin.Context) {
	// 获取 POST 请求体中的表单数据用 PostForm 和 DefaultPostForm
	//cid := c.PostForm("cid")
	//name := c.DefaultPostForm("name", "wangwu")
	//ReturnSuccess(c, 0, cid, name, 1)

	// 接收json格式的数据要用map或者结构体
	// 用BindJSON方法把json数据绑定到map中
	//param := make(map[string]interface{})
	//err := c.BindJSON(&param)
	//if err == nil {
	//	//ReturnSuccess(c, 0, "success", param, 1)
	//	ReturnSuccess(c, 0, param["name"], param["cid"], 1)
	//	return
	//}
	//ReturnError(c, 4001, gin.H{"err": err})

	// 用BindJSON方法把json数据绑定到结构体中
	search := &Search{}
	err := c.BindJSON(search)
	if err == nil {
		//ReturnSuccess(c, 0, "success", param, 1)
		ReturnSuccess(c, 0, search.Name, search.Cid, 1)
		return
	}
	ReturnError(c, 4001, gin.H{"err": err})
}
