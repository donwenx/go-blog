package model

import (
	"blog/dao"
	"time"
)

type Article struct {
	Id           int64  `json:"id"`
	Uid          int64  `json:"uid"`
	Cid          int64  `json:"cid"`
	Title        string `json:"title"`
	Cover        string `json:"cover"`
	Content      string `json:"context"`
	Like         int64  `json:"like"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
	AllowComment int    `json:"allowComment"`
	State        int    `json:"state"`
}

type CreateArticleDto struct {
	Uid     int64  `json:"uid"`
	Cid     int64  `json:"cid"`
	Title   string `json:"title"`
	Cover   string `json:"cover"`
	Content string `json:"context"`
}

func (Article) TableName() string {
	return "article"
}

func CreateArticle(data *CreateArticleDto) (Article, error) {
	article := Article{
		Uid:          data.Uid,
		Cid:          data.Cid,
		Title:        data.Title,
		Cover:        data.Cover,
		Content:      data.Content,
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now().Unix(),
		AllowComment: Valid,
		State:        Valid,
	}
	err := dao.Db.Create(&article).Error
	return article, err
}
