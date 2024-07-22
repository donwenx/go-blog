package model

import "blog/dao"

type TagRelation struct {
	Id        int64 `json:"id"`
	ArticleId int64 `json:"articleId"`
	TagId     int64 `json:"TagId"`
	State     int   `json:"state"`
}

type CreateTagRelationDto struct {
	ArticleId int64 `json:"articleId"`
	TagId     int64 `json:"TagId"`
	State     int   `json:"state"`
}

type UpdateTagRelationDto struct {
	Id    int64 `json:"id"`
	State int   `json:"state"`
}

func (TagRelation) TableName() string {
	return "tag_relation"
}

func CreateTagRelation(data *CreateTagRelationDto) (TagRelation, error) {
	tagRelation := TagRelation{
		ArticleId: data.ArticleId,
		TagId:     data.TagId,
		State:     data.State,
	}
	err := dao.Db.Create(&tagRelation).Error
	return tagRelation, err
}

func GetTagRelationById(id int64) (TagRelation, error) {
	tagRelation := TagRelation{}
	err := dao.Db.Where("id = ? AND state = ?", id, Valid).First(&tagRelation).Error
	return tagRelation, err
}

func GetTagRelationByArticleId(id int64) ([]TagRelation, error) {
	tagRelation := []TagRelation{}
	err := dao.Db.Where("article_id = ? AND state = ?", id, Valid).Find(&tagRelation).Error
	return tagRelation, err
}

func UpdateTagRelation(data *UpdateTagRelationDto) (TagRelation, error) {
	tagRelation := TagRelation{
		State: data.State,
	}
	err := dao.Db.Where("id = ?", data.Id).Updates(&tagRelation).Error
	return tagRelation, err
}
