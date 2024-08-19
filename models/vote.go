package models

import (
	"gin-ranking/dao"
	"time"
)

type Vote struct {
	Id       int       `json:"id"`
	UserId   int       `json:"userId"`
	PlayerId int       `json:"playerId"`
	AddTime  time.Time `json:"addTime"`
}

func (Vote) TableName() string {
	return "vote"
}

func GetVoteInfo(userId, playerId int) (Vote, error) {
	var vote Vote
	err := dao.Db.Where("user_id = ? AND player_id = ?", userId, playerId).First(&vote).Error
	return vote, err
}

func AddVote(userId, playerId int) (int, error) {
	vote := Vote{
		UserId:   userId,
		PlayerId: playerId,
		AddTime:  time.Now(),
	}
	err := dao.Db.Create(&vote).Error
	return vote.Id, err
}
