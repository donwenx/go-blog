package controllers

import (
	"blog/errcode"
	"blog/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func (CommentController) CreateComment(ctx *gin.Context) {
	uidStr := ctx.DefaultPostForm("uid", "0")
	aidStr := ctx.DefaultPostForm("aid", "0")
	parentIdStr := ctx.DefaultPostForm("parentId", "0")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	aid, _ := strconv.ParseInt(aidStr, 10, 64)
	parentId, _ := strconv.ParseInt(parentIdStr, 10, 64)
	context := ctx.DefaultPostForm("context", "")
	if uid == 0 || aid == 0 || context == "" {
		ReturnError(ctx, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	comment, err := model.CreateComment(&model.CreateCommentDto{
		Uid:      uid,
		Aid:      aid,
		Content:  context,
		ParentId: parentId,
	})
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
