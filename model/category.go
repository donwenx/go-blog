package model

import (
	"blog/dao"
	"time"
)

// modules -> 模块
// model -> 模型

type Category struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime"`
	State      int    `json:"state"`
}

type CreateCategoryDto struct {
	Name string `json:"name"`
}

type UpdateCategoryDto struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	State int    `json:"state"`
}

func (Category) TableName() string {
	return "category"
}

func CreateCategory(data *CreateCategoryDto) (Category, error) {
	category := Category{Name: data.Name, CreateTime: time.Now().Unix(), State: Valid}
	err := dao.Db.Create(&category).Error
	return category, err
}

func GetCategoryByName(name string) (Category, error) {
	category := Category{Name: name}
	err := dao.Db.Where("name = ? AND state = ?", name, Valid).First(&category).Error
	return category, err
}

func GetCategoryById(id int64) (Category, error) {
	var category Category
	err := dao.Db.Where("id = ? AND state = ?", id, Valid).Find(&category).Error
	return category, err
}

func GetCateGoryList() ([]Category, error) {
	var category []Category
	err := dao.Db.Where("state = ?", Valid).Find(&category).Error
	return category, err
}

func UpdateCategory(data *UpdateCategoryDto) (Category, error) {
	category := Category{Id: data.Id}
	err := dao.Db.Model(&category).Updates(Category{
		Name:  data.Name,
		State: data.State,
	}).Error
	category, _ = GetCategoryById(category.Id)
	return category, err
}
