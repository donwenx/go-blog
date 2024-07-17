package modules

import (
	"blog/dao"
	"time"
)

type User struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Authority    string `json:"authority"`
	AllowPost    int    `json:"allowPost"`
	AllowComment int    `json:"allowComment"`
	AllowLogin   int    `json:"allowLogin"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
	State        int    `json:"state"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByUserName(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func AddUser(username string, password string) (int64, error) {
	user := User{Username: username, Password: password, AllowPost: 1, AllowComment: 1, AllowLogin: 1, CreateTime: time.Now().Unix(), UpdateTime: time.Now().Unix(), State: 1}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}
