package modules

import (
	"blog/dao"
	"blog/util"
	"time"
)

type Token struct {
	Id         int64  `json:"id"`
	Uid        int64  `json:"uid"`
	CreateTime int64  `json:"createTime"`
	Expire     int64  `json:"expire"`
	Token      string `json:"token"`
	State      int64  `json:"state"`
}

func TableName() string {
	return "token"
}

func CreateToken(userId int64, expire int64) (string, error) {
	createTime := time.Now().Unix()
	time := createTime + expire
	token := Token{Uid: userId, CreateTime: createTime, Expire: time, Token: util.CreateNonceStr(32), State: 1}
	err := dao.Db.Create(&token).Error
	return token.Token, err
}
