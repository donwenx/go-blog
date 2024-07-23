package model

import (
	"blog/dao"
	"time"
)

type Comment struct {
	Id         int64  `json:"id"`
	Uid        int64  `json:"uid"`
	Aid        int64  `json:"Aid"`
	ParentId   int64  `json:"parentId"`
	Content    string `json:"context"`
	Like       int64  `json:"like"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	State      int    `json:"state"`
}

type CreateCommentDto struct {
	Uid      int64  `json:"uid" binding:"required"`
	Aid      int64  `json:"Aid" binding:"required"`
	ParentId int64  `json:"parentId" binding:"required"`
	Content  string `json:"context" binding:"required"`
}

type UpdateCommentDto struct {
	Id         int64 `json:"id" binding:"required"`
	UpdateTime int64 `json:"updateTime"`
	State      int   `json:"state"`
}

func (Comment) TableName() string {
	return "comment"
}

func CreateComment(data *CreateCommentDto) (Comment, error) {
	comment := Comment{
		Uid:        data.Uid,
		Aid:        data.Aid,
		ParentId:   data.ParentId,
		Content:    data.Content,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		State:      Valid,
	}
	err := dao.Db.Create(&comment).Error
	return comment, err
}

func GetCommentById(id int64) (Comment, error) {
	var comment Comment
	err := dao.Db.Where("id = ? AND state = ?", id, Valid).First(&comment).Error
	return comment, err
}

func UpdateComment(data *UpdateCommentDto) (Comment, error) {
	comment := Comment{Id: data.Id}
	err := dao.Db.Model(&comment).Updates(Comment{
		UpdateTime: time.Now().Unix(),
		State:      data.State,
	}).Error
	comment, _ = GetCommentById(comment.Id)
	return comment, err
}
