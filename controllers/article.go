package controllers

import (
	"blog/errcode"
	"blog/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

// 创建
func (a ArticleController) CreateArticle(c *gin.Context) {
	uidStr := c.DefaultPostForm("uid", "0")
	cidStr := c.DefaultPostForm("cid", "0")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	cid, _ := strconv.ParseInt(cidStr, 10, 64)
	title := c.DefaultPostForm("title", "")
	cover := c.DefaultPostForm("cover", "")
	content := c.DefaultPostForm("content", "")
	if uid == 0 || cid == 0 || title == "" || cover == "" || content == "" {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	article, err := model.CreateArticle(&model.CreateArticleDto{
		Uid:     uid,
		Cid:     cid,
		Title:   title,
		Cover:   cover,
		Content: content,
	})
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "创建失败")
		return
	}
	ReturnSuccess(c, 0, "创建成功", article)
}

// 查找id
// 查找列表
// 更新
// 删除
