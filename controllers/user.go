package controllers

import (
	"gin-ranking/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
同一个包下重复文件名的话，会报错
采用结构体的方式解决
*/
type UserController struct{}

func (u UserController) Register(c *gin.Context) {
	// 注册页面三个表单项 用户名 密码 确认密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")

	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	if password != confirmPassword {
		ReturnError(c, 4001, "两次密码不一致")
		return
	}

	user, err := models.GetUserInfoByUsername(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "用户名已存在")
		return
	}

	_, err = models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4001, "保存失败，请联系管理员")
		return
	}

	ReturnSuccess(c, 200, "注册成功", nil, 0)
}

// UserApi 专门定义一个结构体用来返回用户信息，不用model里的User结构体是为了防止返回的结构体里有密码
type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	user, _ := models.GetUserInfoByUsername(username)
	if user.Id == 0 {
		ReturnError(c, 4004, "用户名不存在")
		return
	}
	if user.Password != EncryMd5(password) {
		ReturnError(c, 4004, "密码错误")
		return
	}

	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	session.Save()

	data := UserApi{
		Id:       user.Id,
		Username: user.Username,
	}
	ReturnSuccess(c, 0, "恭喜，登录成功", data, 1)
}
