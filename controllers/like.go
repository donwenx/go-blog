package controllers

import (
	"blog/errcode"
	"blog/model"

	"github.com/gin-gonic/gin"
)

type LikeController struct{}

func (LikeController) CreateLike(c *gin.Context) {
	param := model.CreateLikeDto{}
	err := c.ShouldBindUri(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}
	uid := c.GetInt64("userId")
	// 判断一下是否创建过
	like, err := model.GetLikeByUserId(&model.GetLikeByIdDto{Uid: uid, LikeId: param.LikeId, Type: param.Type})
	if err != nil || like.Id != 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "已点赞"+err.Error())
		return
	}
	like, err = model.CreateLike(&model.CreateLikeDto{
		Uid:    uid,
		LikeId: param.LikeId,
		Type:   param.Type,
	})
	if err != nil || like.Id == model.Invalid {
		ReturnError(c, errcode.ErrInvalidRequest, "更新失败"+err.Error())
		return
	}
	ReturnSuccess(c, 0, "成功", like)
}

type CancelLikeRequest struct {
	Id   int64  `json:"id" binding:"required" uri:"id"`     // like id
	Type string `json:"type" binding:"required" uri:"type"` // 评论，文章，(类型)
}

// 取消点赞
func (LikeController) CancelLike(c *gin.Context) {
	param := CancelLikeRequest{}
	err := c.ShouldBindUri(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}

	uid := c.GetInt64("userId")

	// 判断一下是否创建过
	like, err := model.GetLikeByLikeId(&model.GetLikeByIdDto{Uid: uid, LikeId: param.Id, Type: param.Type})
	if err != nil || like.State == model.Invalid {
		ReturnError(c, errcode.ErrInvalidRequest, "已取消")
		return
	}

	like, err = model.UpdateLike(&model.UpdateLikeDto{
		Id:    like.Id,
		State: model.Invalid,
	})
	if err != nil || like.Id == model.Invalid {
		ReturnError(c, errcode.ErrInvalidRequest, "更新失败"+err.Error())
		return
	}
	ReturnSuccess(c, 0, "成功", like)
}
