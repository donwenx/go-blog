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

func GetTagList(id int64) ([]Tag, error) {
	tag := []Tag{}
	err := dao.Db.Where("id = ? AND state = ?", id, Valid).Find(&tag).Error
	return tag, err
}
