package controllers

import (
	"blog/errcode"
	"blog/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

type CreateArticleRequest struct {
	Uid     int64   `json:"uid" binding:"required"`
	Cid     int64   `json:"cid" binding:"required"`
	Title   string  `json:"title" binding:"required"`
	Cover   string  `json:"cover" binding:"required"`
	Content string  `json:"content" binding:"required"`
	TagId   []int64 `json:"tagId"`
}

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id        int64  `json:"id"`
	Avatar    string `json:"avatar"`
	Username  string `json:"username"`
	Authority string `json:"authority"`
}

type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ArticleResponse struct {
	User     User     `json:"user"`
	Category Category `json:"category"`
	Title    string   `json:"title"`
	Cover    string   `json:"cover"`
	Content  string   `json:"context"`
	Tag      []Tag    `json:"tag"`
}

// 创建
func (a ArticleController) CreateArticle(c *gin.Context) {
	param := CreateArticleRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}
	article, err := model.CreateArticle(&model.CreateArticleDto{
		Uid:     param.Uid,
		Cid:     param.Cid,
		Title:   param.Title,
		Cover:   param.Cover,
		Content: param.Content,
	})
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "创建失败")
		return
	}
	for k, _ := range param.TagId {
		_, err := model.CreateTagRelation(&model.CreateTagRelationDto{
			ArticleId: article.Id,
			TagId:     param.TagId[k],
		})

		if err != nil {
			ReturnError(c, errcode.ErrInvalidRequest, "创建失败")
			return
		}
	}
	// 查文章
	article, err = model.GetArticleById(article.Id)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "创建失败")
		return
	}

	response, err := CreateArticleResponse(&article)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "创建失败")
		return
	}
	ReturnSuccess(c, 0, "成功", response)
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
	response, err := CreateArticleResponse(&article)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "创建失败")
		return
	}
	ReturnSuccess(c, 0, "成功", response)
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

	response := make([]*ArticleResponse, len(article))
	for k := range article {
		response[k], err = CreateArticleResponse(&article[k])
		if err != nil {
			ReturnError(c, errcode.ErrInvalidRequest, "创建失败")
			return
		}
	}
	ReturnSuccess(c, 0, "查找成功", response)
}

// 查找列表
func (a ArticleController) GetArticleList(c *gin.Context) {
	article, err := model.GetArticleList()
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "查找失败")
		return
	}

	response := make([]*ArticleResponse, len(article))
	for k := range article {
		response[k], err = CreateArticleResponse(&article[k])
		if err != nil {
			ReturnError(c, errcode.ErrInvalidRequest, "创建失败")
			return
		}
	}
	ReturnSuccess(c, 0, "查找成功", response)
}

// 更新
func (a ArticleController) UpdateArticle(c *gin.Context) {
	param := model.UpdateArticleDto{}
	err := c.ShouldBind(&param)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}
	article, _ := model.GetArticleById(param.Id)
	if article.Id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "文章不存在")
		return
	}
	// 更新数据库
	article, err = model.UpdateArticle(&model.UpdateArticleDto{
		Id:      param.Id,
		Cid:     param.Cid,
		Title:   param.Title,
		Cover:   param.Cover,
		Content: param.Content,
	})
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "更新失败")
		return
	}
	response, err := CreateArticleResponse(&article)
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "创建失败")
		return
	}
	ReturnSuccess(c, 0, "成功", response)
}

// 删除
func (a ArticleController) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id == 0 {
		ReturnError(c, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	article, _ := model.GetArticleById(id)
	if article.Id == 0 || article.State == model.Invalid {
		ReturnError(c, errcode.ErrInvalidRequest, "文章不存在")
		return
	}
	article, err := model.UpdateArticle(&model.UpdateArticleDto{
		Id:    id,
		State: model.Invalid,
	})
	if err != nil {
		ReturnError(c, errcode.ErrInvalidRequest, "删除失败")
		return
	}
	ReturnSuccess(c, 0, "删除成功", "")
}

func CreateArticleResponse(data *model.Article) (*ArticleResponse, error) {

	// 查用户
	user, err := model.GetUserInfoById(data.Uid)
	if err != nil {
		return nil, err
	}
	// 查分类
	category, err := model.GetCategoryById(data.Cid)
	if err != nil {
		return nil, err
	}
	// 查tag
	tagRelation, err := model.GetTagRelationByArticleId(data.Id)
	if err != nil {
		return nil, err
	}
	tag := make([]model.Tag, len(tagRelation))
	for k := range tagRelation {
		tag[k], err = model.GetTagById(tagRelation[k].TagId)
		if err != nil {
			return nil, err
		}
	}

	userResp := User{Id: user.Id, Avatar: user.Avatar, Username: user.Username, Authority: user.Authority}
	categoryResp := Category{Id: category.Id, Name: category.Name}
	tagResp := make([]Tag, len(tag))
	for k := range tag {
		tagResp[k] = Tag{Id: tag[k].Id, Name: tag[k].Name}
	}

	response := &ArticleResponse{
		User:     userResp,
		Category: categoryResp,
		Title:    data.Title,
		Cover:    data.Cover,
		Content:  data.Content,
		Tag:      tagResp,
	}
	return response, nil
}
