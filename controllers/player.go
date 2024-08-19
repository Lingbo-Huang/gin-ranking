package controllers

import (
	"fmt"
	"gin-ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
	rs, err := models.GetPlayers(aid)
	if err != nil {
		ReturnError(c, 4004, fmt.Sprintf("没有%d号活动的参赛选手相关信息", aid))
		return
	}
	ReturnSuccess(c, 0, "success", rs, 1)
}
