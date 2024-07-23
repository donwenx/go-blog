package controllers

import (
	"blog/errcode"
	"blog/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserResponse struct {
	Id           int64  `json:"id"`
	Avatar       string `json:"avatar"`
	Username     string `json:"username"`
	Authority    string `json:"authority"`
	AllowPost    int    `json:"allowPost"`
	AllowComment int    `json:"allowComment"`
	AllowLogin   int    `json:"allowLogin"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
	State        int    `json:"state"`
}

func (u UserController) Register(c *gin.Context) {
	param := model.CreateUserDto{}
	err := c.ShouldBind(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}
	user, _ := model.GetUserInfoByUserName(param.Username)
	if user.Id != 0 {
		ReturnError(c, 4001, "用户名已存在")
		return
	}

	// 注册user
	user, err = model.CreateUser(&model.CreateUserDto{
		Username: param.Username,
		Password: Md5(param.Password),
	})

	if err != nil {
		ReturnError(c, 4001, "用户注册失败")
		return
	}

	response := UserResponse{
		Id:           user.Id,
		Avatar:       user.Avatar,
		Username:     user.Username,
		Authority:    user.Authority,
		AllowPost:    user.AllowPost,
		AllowComment: user.AllowComment,
		AllowLogin:   user.AllowLogin,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
		State:        user.State,
	}
	ReturnSuccess(c, 0, "注册成功", response)
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
	param := model.CreateUserDto{}
	err := c.ShouldBind(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}
	user, _ := model.GetUserInfoByUserName(param.Username)
	if user.Id == 0 {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}
	if user.Password != Md5(param.Password) {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}
	if user.AllowLogin != model.Valid {
		ReturnError(c, 4001, "用户禁止登录")
		return
	}
	// 存入token
	token, _ := model.CreateToken(user.Id, 24*60*60)
	response := LoginResponse{
		Id:           user.Id,
		Username:     param.Username,
		Authority:    user.Authority,
		AllowPost:    user.AllowPost,
		AllowComment: user.AllowComment,
		AllowLogin:   user.AllowLogin,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
		Token:        token,
	}
	ReturnSuccess(c, 0, "登录成功", response)
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
	param := model.UpdateUserDto{}
	err := c.ShouldBind(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}

	user, _ := model.GetUserInfoById(param.Id)
	if user.Id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "用户不存在")
		return
	}
	// 更新数据库
	user, err = model.UpdateUser(&model.UpdateUserDto{
		Id:       param.Id,
		Username: param.Username,
		Avatar:   param.Avatar,
	})
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "更新失败")
		return
	}
	response := UserResponse{
		Id:           user.Id,
		Avatar:       user.Avatar,
		Username:     user.Username,
		Authority:    user.Authority,
		AllowPost:    user.AllowPost,
		AllowComment: user.AllowComment,
		AllowLogin:   user.AllowLogin,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
		State:        user.State,
	}
	ReturnSuccess(c, 0, "更新成功", response)
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
