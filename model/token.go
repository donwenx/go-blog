package model

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

func (Token) TableName() string {
	return "token"
}

func CreateToken(userId int64, expire int64) (string, error) {
	createTime := time.Now().Unix()
	time := createTime + expire
	token := Token{Uid: userId, CreateTime: createTime, Expire: time, Token: util.CreateNonceStr(32), State: Valid}
	err := dao.Db.Create(&token).Error
	return token.Token, err
}

// 查找token信息
func GetTokenInfo(tokenStr string) (Token, error) {
	token := Token{}
	err := dao.Db.Where("token = ?", tokenStr).First(&token).Error
	return token, err
}

// 设置token过期
func SetTokenOutLog(token string) {
	dao.Db.Model(&Token{}).Where("token = ?", token).UpdateColumn("state", Invalid)
}
