package model

import (
	"blog/dao"
	"time"
)

type Like struct {
	Id         int64  `json:"id"`
	Uid        int64  `json:"uid"`
	LikeId     int64  `json:"likeId"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"UpdateTime"`
	Type       string `json:"type"` // 评论，文章，(类型)
	State      int    `json:"state"`
}

func (Like) TableName() string {
	return "like"
}

type CreateLikeDto struct {
	Uid    int64  `json:"uid"`
	LikeId int64  `json:"likeId"`
	Type   string `json:"type"` // 评论，文章，(类型)
}

func CreateLike(data *CreateLikeDto) (Like, error) {
	like := Like{
		Uid:        data.Uid,
		LikeId:     data.LikeId,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		Type:       data.Type,
		State:      Valid,
	}
	err := dao.Db.Create(&like).Error
	return like, err
}

type GetLikeCountDto struct {
	LikeId int64  `json:"likeId"`
	Type   string `json:"type"` // 评论，文章，(类型)
}

func GetLikeCount(data *GetLikeCountDto) (int64, error) {
	like := Like{}
	var count int64
	err := dao.Db.Model(&like).Where("type = ? AND likeId = ?", data.LikeId, data.Type).Count(&count).Error
	return count, err
}

type UpdateLikeDto struct {
	Id    int64 `json:"id"`
	State int   `json:"state"`
}

func UpdateLike(data *UpdateLikeDto) (Like, error) {
	like := Like{
		Id:    data.Id,
		State: data.State,
	}
	err := dao.Db.Updates(&like).Error
	return like, err
}
