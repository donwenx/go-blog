package modules

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
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByUserName(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func AddUser(data *AddUserDto) (int64, error) {
	user := User{
		Username:  data.Username,
		Password:  data.Password,
		AllowPost: 1, AllowComment: 1, AllowLogin: 1,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		State:      1,
	}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

func UpdateUser(data *UpdateUserDto) error {
	err := dao.Db.Model(User{Id: data.Id}).Updates(User{
		Password: data.Password,
	}).Error
	return err
}
