package controllers

import (
	"blog/errcode"
	"blog/model"
	"strconv"

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
	user, _ := model.GetUserInfoByUserName(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "用户名已存在")
		return
	}

	// 注册user
	_, err := model.AddUser(&model.AddUserDto{
		Username: username,
		Password: Md5(password),
	})

	if err != nil {
		ReturnError(c, 4001, "用户注册失败")
		return
	}
	ReturnSuccess(c, 0, "注册成功", "")
}

type LoginResponse struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
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
	user, _ := model.GetUserInfoByUserName(username)
	if user.Id == 0 {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}
	if user.Password != Md5(password) {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}
	if user.AllowLogin != model.Valid {
		ReturnError(c, 4001, "用户禁止登录")
		return
	}
	// 存入token
	token, _ := model.CreateToken(user.Id, 24*60*60)
	data := LoginResponse{
		Id:           user.Id,
		Username:     username,
		Authority:    user.Authority,
		AllowPost:    user.AllowPost,
		AllowComment: user.AllowComment,
		AllowLogin:   user.AllowLogin,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
		Token:        token,
	}
	ReturnSuccess(c, 0, "登录成功", data)
}

func (u UserController) LogOut(c *gin.Context) {
	token := c.GetHeader("token")
	model.SetTokenOutLog(token)
}

func (u UserController) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	user, err := model.GetUserInfoById(id)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "用户不存在")
		return
	}
	ReturnSuccess(c, 0, "查询成功", user)
}

func (u UserController) GetUserList(c *gin.Context) {
	user, err := model.GetUserList()
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "查询失败")
		return
	}
	ReturnSuccess(c, 0, "查询成功", user)
}

func (u UserController) UpdateUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "0")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	username := c.DefaultPostForm("username", "")
	avatar := c.DefaultPostForm("avatar", "")
	if id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}

	user, _ := model.GetUserInfoById(id)
	if user.Id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "用户不存在")
		return
	}
	// 更新数据库
	user, err := model.UpdateUser(&model.UpdateUserDto{
		Id:       id,
		Username: username,
		Avatar:   avatar,
	})
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "更新失败")
		return
	}
	ReturnSuccess(c, 0, "更新成功", user)
}

func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	user, _ := model.GetUserInfoById(id)
	if user.Id == 0 || user.State == model.Invalid {
		ReturnError(c, errcode.ErrInvalidRequest, "用户不存在")
		return
	}
	user, err := model.UpdateUser(&model.UpdateUserDto{
		Id:    id,
		State: model.Invalid,
	})
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "删除失败")
		return
	}
	ReturnSuccess(c, 0, "删除成功", "")
}
