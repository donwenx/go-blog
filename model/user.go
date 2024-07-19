package model

import (
	"blog/dao"
	"time"
)

type User struct {
	Id           int64  `json:"id"`
	Avatar       string `json:"avatar"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Authority    string `json:"authority"`
	AllowPost    int    `json:"allowPost" gorm:"default:1"`
	AllowComment int    `json:"allowComment" gorm:"default:1"`
	AllowLogin   int    `json:"allowLogin" gorm:"default:1"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
	State        int    `json:"state"`
}

type AddUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserDto struct {
	Id       int64  `json:"id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Password string `json:"password"`
	State    int    `json:"state"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByUserName(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func GetUserInfoById(id int64) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func GetUserList() ([]User, error) {
	var user []User
	err := dao.Db.Find(&user).Error
	return user, err
}

func AddUser(data *AddUserDto) (int64, error) {
	user := User{
		Username:  data.Username,
		Password:  data.Password,
		AllowPost: Valid, AllowComment: Valid, AllowLogin: Valid,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		State:      Valid,
	}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

func UpdateUser(data *UpdateUserDto) (User, error) {
	user := User{Id: data.Id}
	err := dao.Db.Model(&user).Updates(User{
		Password: data.Password,
		Avatar:   data.Avatar,
		Username: data.Username,
		State:    data.State,
	}).Error
	user, _ = GetUserInfoById(user.Id)
	return user, err
}
