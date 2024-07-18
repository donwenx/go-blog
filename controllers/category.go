package controllers

import (
	"blog/errcode"
	"blog/modules"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

func (c CategoryController) CreateCategory(ctx *gin.Context) {
	name := ctx.DefaultPostForm("name", "")
	_, err := modules.CreateCategory(name)
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "创建失败")
		return
	}
	ReturnSuccess(ctx, 0, "创建成功", "")
}
