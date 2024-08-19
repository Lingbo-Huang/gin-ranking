package models

import (
	"gin-ranking/dao"
	"time"
)

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	AddTime    time.Time `json:"add_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func GetUserInfo(id int) (User, error) {
	var user User
	err := dao.Db.Where("id =?", id).First(&user).Error
	return user, err
}

func AddUser(username string, password string) (int, error) {
	user := User{
		Username:   username,
		Password:   password,
		AddTime:    time.Now(),
		UpdateTime: time.Now(),
	}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}
