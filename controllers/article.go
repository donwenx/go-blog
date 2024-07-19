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
func (a ArticleController) GetArticleById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	article, err := model.GetArticleById(id)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "查找失败")
		return
	}
	ReturnSuccess(c, 0, "查找成功", article)
}

// 查找keyword
func (a ArticleController) GetArticleByKeyword(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	article, err := model.GetArticleByKeyword(keyword)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "查找失败")
		return
	}
	ReturnSuccess(c, 0, "查找成功", article)
}

// 查找列表
func (a ArticleController) GetArticleList(c *gin.Context) {
	article, err := model.GetArticleList()
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "查找失败")
		return
	}
	ReturnSuccess(c, 0, "查找成功", article)
}

// 更新
func (a ArticleController) UpdateArticle(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "0")
	cidStr := c.DefaultPostForm("cid", "0")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	cid, _ := strconv.ParseInt(cidStr, 10, 64)
	title := c.DefaultPostForm("title", "")
	cover := c.DefaultPostForm("cover", "")
	content := c.DefaultPostForm("content", "")
	if id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	article, _ := model.GetArticleById(id)
	if article.Id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "文章不存在")
		return
	}
	// 更新数据库
	article, err := model.UpdateArticle(&model.UpdateArticleDto{
		Id:      id,
		Cid:     cid,
		Title:   title,
		Cover:   cover,
		Content: content,
	})
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "更新失败")
		return
	}
	ReturnSuccess(c, 0, "更新成功", article)
}

// 删除
