package controllers

import (
	"fmt"
	"gin-ranking/cache"
	"gin-ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
	rs, err := models.GetPlayers(aid, "id asc")
	if err != nil {
		ReturnError(c, 4004, fmt.Sprintf("没有%d号活动的参赛选手相关信息", aid))
		return
	}
	ReturnSuccess(c, 0, "success", rs, 1)
}

func (p PlayerController) GetRanking(c *gin.Context) {
	// 先从redis获取数据，如果redis没有数据，则从mysql获取数据，然后存入redis
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
	// 从redis获取数据
	var redisKey string
	redisKey = "ranking:" + aidStr
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()
	if err == nil && len(rs) > 0 {
		var players []models.Player
		for _, value := range rs {
			id, _ := strconv.Atoi(value)
			rsInfo, _ := models.GetPlayerInfo(id)
			if rsInfo.Id > 0 { // 取到值就放players里，详细信息方便传给前端
				players = append(players, rsInfo)
			}
		}
		ReturnSuccess(c, 0, "success", players, 1)
		return
	}
	// 从mysql获取数据
	rsDb, errDb := models.GetPlayers(aid, "score desc")
	if errDb == nil {
		// 保存到redis
		for _, value := range rsDb {
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.Id, value.Score)).Err()

		}
		// 设置过期时间
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)

		ReturnSuccess(c, 0, "success", rsDb, 1)
	} else {
		ReturnError(c, 4004, fmt.Sprintf("没有%d号活动的参赛选手相关信息", aid))
	}
}
