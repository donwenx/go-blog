package controllers

import (
	"blog/errcode"
	"blog/model"

	"github.com/gin-gonic/gin"
)

type TagController struct{}

func (t TagController) CreateTag(c *gin.Context) {
	param := model.CreateTagDto{}
	err := c.ShouldBind(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	tag, _ := model.GetTagByName(param.Name)
	if tag.Id != 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "名字已存在")
		return
	}
	tag, err = model.CreateTag(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "创建失败")
		return
	}
	ReturnSuccess(c, 0, "创建成功", tag)
}

func (t TagController) GetTagList(c *gin.Context) {
	tag, err := model.GetTagList()
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "查找失败")
		return
	}
	ReturnSuccess(c, 0, "创建成功", tag)
}

type UpdateTagRequest struct {
	Id   int64  `json:"id"  form:"id" binding:"required"`
	Name string `json:"name"  form:"name" binding:"required"`
}

func (t TagController) UpdateTag(c *gin.Context) {
	param := UpdateTagRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息:"+err.Error())
		return
	}
	tag, err := model.UpdateTag(&model.UpdateTagDto{
		Id:   param.Id,
		Name: param.Name,
	})
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "更新失败")
		return
	}
	ReturnSuccess(c, 0, "更新成功", tag)
}

type DeleteTagRequest struct {
	Id int64 `json:"id" uri:"id" binding:"required"`
}

func (t TagController) DeleteTag(c *gin.Context) {
	param := DeleteTagRequest{}
	err := c.ShouldBindUri(&param)
	if param.Id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息"+err.Error())
		return
	}
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息:"+err.Error())
		return
	}
	tag, _ := model.GetTagById(param.Id)
	if tag.State == model.Invalid {
		ReturnError(c, errcode.ErrInvalidRequest, "已删除")
		return
	}
	_, err = model.UpdateTag(&model.UpdateTagDto{
		Id:    param.Id,
		State: model.Invalid,
	})
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "删除失败"+err.Error())
		return
	}
	ReturnSuccess(c, 0, "删除成功", "")
}
