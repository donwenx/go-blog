package controllers

import (
	"blog/errcode"
	"blog/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

// 创建
func (c CategoryController) CreateCategory(ctx *gin.Context) {
	param := model.CreateCategoryDto{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}

	// 创建前判断是否已存在
	category, _ := model.GetCategoryByName(param.Name)
	if category.Id != 0 {
		ReturnError(ctx, errcode.ErrInvalidRequest, "名称已存在")
		return
	}
	category, err = model.CreateCategory(&param)
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "创建失败")
		return
	}
	ReturnSuccess(ctx, 0, "创建成功", category)
}

// 获取1个，根据id查找
func (c CategoryController) GetCategoryById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	// 判断id是否存在
	if id == 0 {
		ReturnError(ctx, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	category, err := model.GetCategoryById(id)
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "获取失败")
		return
	}
	ReturnSuccess(ctx, 0, "查询成功", category)
}

// 获取列表
func (c CategoryController) GetCateGoryList(ctx *gin.Context) {
	category, err := model.GetCateGoryList()
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "查询失败")
		return
	}
	ReturnSuccess(ctx, 0, "查询成功", category)
}

// 更新
func (c CategoryController) UpdateCategory(ctx *gin.Context) {
	param := model.UpdateCategoryDto{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "绑定失败"+err.Error())
		return
	}

	// 判断分类是否存在
	category, _ := model.GetCategoryById(param.Id)
	if category.Id == 0 {
		ReturnError(ctx, errcode.ErrInvalidRequest, "分类不存在")
		return
	}
	// 更新数据库
	category, err = model.UpdateCategory(&model.UpdateCategoryDto{
		Id:   param.Id,
		Name: param.Name,
	})
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "更新失败")
		return
	}
	ReturnSuccess(ctx, 0, "更新成功", category)
}

// 删除
func (c CategoryController) DeleteCategory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id == 0 {
		ReturnError(ctx, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	category, _ := model.GetCategoryById(id)
	if category.State == model.Invalid {
		ReturnError(ctx, errcode.ErrInvalidRequest, "分类已删除")
		return
	}
	category, err := model.UpdateCategory(&model.UpdateCategoryDto{
		Id:    id,
		State: model.Invalid,
	})
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "删除失败")
		return
	}
	ReturnSuccess(ctx, 0, "删除成功", "")
}
