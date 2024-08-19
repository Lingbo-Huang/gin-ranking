package controllers

import (
	"gin-ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {
	// 获取用户id(投票人id)， 选手id
	userIdStr := c.DefaultPostForm("userId", "0")
	playerIdStr := c.DefaultPostForm("playerId", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)

	if userId == 0 || playerId == 0 {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}
	user, _ := models.GetUserInfo(userId)
	if user.Id == 0 {
		ReturnError(c, 4001, "投票用户不存在")
		return
	}
	player, _ := models.GetPlayerInfo(playerId)
	if player.Id == 0 {
		ReturnError(c, 4001, "参赛选手不存在")
		return
	}

	// 每个人只能给选手投一次票
	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.Id != 0 {
		ReturnError(c, 4001, "您已投过票")
		return
	}
	// 添加投票记录，并且增加score
	rs, err := models.AddVote(userId, playerId)
	if err == nil {
		// 如果添加成功，更新选手的score
		models.UpdatePlayerScore(playerId)
		ReturnSuccess(c, 0, "投票成功", rs, 1)
	} else {
		ReturnError(c, 4004, "投票失败，请联系管理员")
	}
}
