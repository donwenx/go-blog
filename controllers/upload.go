package controllers

import (
	"blog/errcode"
	"blog/model"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UploadController struct{}

func (UploadController) CreateUpload(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "创建失败"+err.Error())
		return
	}
	path := "./upload/" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, path)

	param := model.CreateUploadDto{
		Name:     file.Filename,
		Path:     path,
		Size:     file.Size,
		MimeType: file.Header.Get("content-type"),
	}
	upload, err := model.CreateUpload(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "创建失败"+err.Error())
		return
	}
	ReturnSuccess(c, 0, "注册成功", upload)
}

func (UploadController) DeleteUpload(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	upload, _ := model.GetUploadById(id)
	if upload.State == model.Invalid {
		ReturnError(c, errcode.ErrInvalidRequest, "文件不存在")
		return
	}
	upload, err := model.UpdateUpload(&model.UpdateUploadDto{
		Id:    id,
		State: model.Invalid,
	})
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "删除失败")
		return
	}
	ReturnSuccess(c, 0, "删除成功", "")
}

func (UploadController) GetUploadById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	upload, _ := model.GetUploadById(id)
	if upload.State == model.Invalid {
		ReturnError(c, errcode.ErrInvalidRequest, "文件不存在")
		return
	}
	ReturnSuccess(c, 0, "查询成功", upload)
}

func (UploadController) Download(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	upload, err := model.GetUploadById(id)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "读取失败"+err.Error())
		return
	}
	if upload.State == model.Invalid {
		ReturnError(c, errcode.ErrInvalidRequest, "文件不存在")
		return
	}
	file, err := os.ReadFile(upload.Path)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "读取失败"+err.Error())
		return
	}
	// 返回文件数据
	c.Data(http.StatusOK, upload.MimeType, file)
}
