package controllers

import (
	"gin-ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
同一个包下重复文件名的话，会报错
采用结构体的方式解决
*/
type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	// 通过解析请求的 URL 参数，获取用户的唯一标识符（ID）
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	name := c.Param("name")

	user, _ := models.GetUserTest(id)

	ReturnSuccess(c, 0, name, user, 1)
}

func (u UserController) AddUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(c, 4002, "AddUser 保存失败")
	} else {
		ReturnSuccess(c, 0, "AddUser 保存成功", id, 1)
	}
}

func (u UserController) UpdateUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	models.UpdateUser(id, username)
	ReturnSuccess(c, 0, "UpdateUser 更新成功", id, 1)

}

func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 4003, "DeleteUser 删除失败")
	} else {
		ReturnSuccess(c, 0, "DeleteUser 删除成功", true, 1)
	}
}

func (u UserController) GetList(c *gin.Context) {
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	ReturnError(c, 4004, num3)
}

func (u UserController) GetUserListTest(c *gin.Context) {
	users, err := models.GetUserListTest()
	if err != nil {
		ReturnError(c, 4005, "GetUserListTest 获取失败")
	} else {
		ReturnSuccess(c, 0, "GetUserListTest 获取成功", users, 1)
	}
}
