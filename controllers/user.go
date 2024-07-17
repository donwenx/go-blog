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
