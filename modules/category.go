package modules

import (
	"blog/dao"
	"time"
)

type Category struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime"`
	State      int    `json:"state" gorm:"default:1"`
}

func (Category) TableName() string {
	return "category"
}

func CreateCategory(name string) (Category, error) {
	category := Category{Name: name, CreateTime: time.Now().Unix()}
	err := dao.Db.Create(&category).Error
	return category, err
}
