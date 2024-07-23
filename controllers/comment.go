package controllers

import (
	"blog/errcode"
	"blog/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func (CommentController) CreateComment(ctx *gin.Context) {
	param := model.CreateCommentDto{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}

	comment, err := model.CreateComment(&param)
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "创建失败")
		return
	}
	ReturnSuccess(ctx, 0, "创建成功", comment)
}

func (CommentController) DeleteComment(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id == 0 {
		ReturnError(ctx, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	comment, _ := model.GetCommentById(id)
	if comment.State == model.Invalid {
		ReturnError(ctx, errcode.ErrInvalidRequest, "已删除")
		return
	}
	comment, err := model.UpdateComment(&model.UpdateCommentDto{
		Id:    id,
		State: model.Invalid,
	})
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "删除失败")
		return
	}
	ReturnSuccess(ctx, 0, "删除成功", "")
}
