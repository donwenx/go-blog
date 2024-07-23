package controllers

import (
	"blog/errcode"
	"blog/model"

	"github.com/gin-gonic/gin"
)

type LikeController struct{}

func (LikeController) CreateLike(c *gin.Context) {
	param := model.CreateLikeDto{}
	err := c.ShouldBind(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}
	like, err := model.CreateLike(&param)
	if err != nil || like.Id == model.Invalid {
		ReturnError(c, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}
	ReturnSuccess(c, 0, "成功", like)
}

// 取消点赞
func (LikeController) CancelLike(c *gin.Context) {
	param := model.UpdateLikeDto{}
	err := c.ShouldBind(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}
	like, err := model.UpdateLike(&model.UpdateLikeDto{
		Id:    param.Id,
		State: model.Valid,
	})
	if err != nil || like.Id == model.Invalid {
		ReturnError(c, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}
	ReturnSuccess(c, 0, "成功", like)
}
