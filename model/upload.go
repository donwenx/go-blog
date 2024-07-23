package model

import (
	"blog/dao"
	"time"
)

type Upload struct {
	Id         int64  `json:"id"`
	Uid        int64  `json:"uid"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Size       int64  `json:"size"`
	MimeType   string `json:"mimeType"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	State      int    `json:"state"`
}

func (Upload) TableName() string {
	return "upload"
}

type CreateUploadDto struct {
	Uid      int64  `json:"uid"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Size     int64  `json:"size"`
	MimeType string `json:"mimeType"`
}

func CreateUpload(data *CreateUploadDto) (Upload, error) {
	upload := Upload{
		Uid:        data.Uid,
		Name:       data.Name,
		Path:       data.Path,
		Size:       data.Size,
		MimeType:   data.MimeType,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		State:      Valid,
	}
	err := dao.Db.Create(&upload).Error
	return upload, err
}

func GetUploadById(id int64) (Upload, error) {
	upload := Upload{}
	err := dao.Db.Where("id = ?", id).Find(&upload).Error
	return upload, err
}

type UpdateUploadDto struct {
	Id       int64  `json:"id" binding:"required"`
	Uid      int64  `json:"uid"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Size     int64  `json:"size"`
	MimeType string `json:"mimeType"`
	State    int    `json:"state"`
}

func UpdateUpload(data *UpdateUploadDto) (Upload, error) {
	upload := Upload{
		Uid:        data.Uid,
		Name:       data.Name,
		Path:       data.Path,
		Size:       data.Size,
		MimeType:   data.MimeType,
		State:      data.State,
		UpdateTime: time.Now().Unix(),
	}
	err := dao.Db.Where("id = ?", data.Id).Updates(&upload).Error
	upload, _ = GetUploadById(data.Id)
	return upload, err
}
