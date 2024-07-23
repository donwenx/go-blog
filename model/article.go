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

type UpdateArticleDto struct {
	Id           int64  `json:"id" bind:"required"`
	Uid          int64  `json:"uid"`
	Cid          int64  `json:"cid"`
	Title        string `json:"title"`
	Cover        string `json:"cover"`
	Content      string `json:"context"`
	AllowComment int    `json:"allowComment"`
	State        int    `json:"state"`
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

func GetArticleByKeyword(keyword string) ([]Article, error) {
	var article []Article
	err := dao.Db.Where("Title like concat('%',?,'%') AND state = ?", keyword, Valid).Find(&article).Error
	return article, err
}

func GetArticleById(id int64) (Article, error) {
	var article Article
	err := dao.Db.Where("id = ? AND state = ?", id, Valid).First(&article).Error
	return article, err
}

func GetArticleList() ([]Article, error) {
	var article []Article
	err := dao.Db.Where("state = ?", Valid).Find(&article).Error
	return article, err
}

// 更新
func UpdateArticle(data *UpdateArticleDto) (Article, error) {
	article := Article{Id: data.Id}
	err := dao.Db.Model(&article).Updates(Article{
		Cid:          data.Cid,
		Title:        data.Title,
		Cover:        data.Cover,
		Content:      data.Content,
		UpdateTime:   time.Now().Unix(),
		AllowComment: data.AllowComment,
		State:        data.State,
	}).Error
	article, _ = GetArticleById(article.Id)
	return article, err
}
