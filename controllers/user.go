package controllers

import (
	"blog/modules"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Register(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}
	user, err := modules.GetUserInfoByUserName(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "用户名已存在")
	}

	// 注册user
	_, err = modules.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4001, "用户注册失败")
	}
	ReturnSuccess(c, 0, "注册成功", "")
}

type UserApi struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Authority    string `json:"authority"`
	AllowPost    int    `json:"allowPost"`
	AllowComment int    `json:"allowComment"`
	AllowLogin   int    `json:"allowLogin"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
	Token        string `json:"token"`
}

func (u UserController) Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}
	user, _ := modules.GetUserInfoByUserName(username)
	if user.Id == 0 {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}
	if user.Password != EncryMd5(password) {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}
	if user.AllowLogin != 1 {
		ReturnError(c, 4001, "用户禁止登录")
		return
	}
	// 存入token
	token, _ := modules.CreateToken(user.Id, 24*60*60)
	data := UserApi{
		Id:           user.Id,
		Username:     username,
		Authority:    user.Authority,
		AllowPost:    user.AllowPost,
		AllowComment: user.AllowComment,
		AllowLogin:   user.AllowLogin,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
		Token:        token}
	ReturnSuccess(c, 0, "登录成功", data)
}
