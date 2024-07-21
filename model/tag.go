package model

import (
	"blog/dao"
	"time"
)

type Tag struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Uid        int64  `json:"uid"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	State      int    `json:"state"`
}

type CreateTagDto struct {
	Name string `json:"name" form:"name"`
	Uid  int64  `json:"uid" form:"uid"`
}

type UpdateTagDto struct {
	Id    int64  `json:"id"  form:"id" binding:"required"`
	Name  string `json:"name"  form:"name"`
	State int    `json:"state"  form:"state"`
}

func (Tag) TableName() string {
	return "tag"
}

func CreateTag(data *CreateTagDto) (Tag, error) {
	tag := Tag{
		Name:       data.Name,
		Uid:        data.Uid,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		State:      Valid,
	}
	err := dao.Db.Create(&tag).Error
	return tag, err
}

func GetTagByName(name string) (Tag, error) {
	tag := Tag{}
	err := dao.Db.Where("name = ? AND state = ?", name, Valid).Find(&tag).Error
	return tag, err
}

func GetTagById(id int64) (Tag, error) {
	tag := Tag{}
	err := dao.Db.Where("id = ? AND state = ?", id, Valid).First(&tag).Error
	return tag, err
}

func GetTagList() ([]Tag, error) {
	tag := []Tag{}
	err := dao.Db.Find(&tag).Error
	return tag, err
}

func UpdateTag(data *UpdateTagDto) (Tag, error) {
	tag := Tag{
		Name:       data.Name,
		UpdateTime: time.Now().Unix(),
		State:      data.State,
	}
	err := dao.Db.Where("id = ?", data.Id).Updates(&tag).Error
	tag, _ = GetTagById(data.Id)
	return tag, err
}
