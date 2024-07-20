package controllers

import (
	"blog/errcode"
	"blog/model"
	"fmt"

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
	fmt.Println(param)
	if param.Name == "" || param.Uid == model.Invalid {
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
