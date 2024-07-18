package controllers

import (
	"blog/errcode"
	"blog/modules"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

// 创建
func (c CategoryController) CreateCategory(ctx *gin.Context) {
	name := ctx.DefaultPostForm("name", "")
	if name == "" {
		ReturnError(ctx, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	// 创建前判断是否已存在
	category, _ := modules.GetCategoryByName(name)
	if category.Id != 0 {
		ReturnError(ctx, errcode.ErrInvalidRequest, "名称已存在")
		return
	}
	_, err := modules.CreateCategory(&modules.CreateCategoryDto{
		Name: name,
	})
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "创建失败")
		return
	}
	ReturnSuccess(ctx, 0, "创建成功", "")
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
	category, err := modules.GetCategoryById(id)
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "获取失败")
		return
	}
	ReturnSuccess(ctx, 0, "查询成功", category)
}

// 获取列表
func (c CategoryController) GetCateGoryList(ctx *gin.Context) {
	category, err := modules.GetCateGoryList()
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "查询失败")
		return
	}
	ReturnSuccess(ctx, 0, "查询成功", category)
}

// 更新
func (c CategoryController) UpdateCategory(ctx *gin.Context) {
	idStr := ctx.DefaultPostForm("id", "0")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	name := ctx.DefaultPostForm("name", "")
	stateStr := ctx.DefaultPostForm("state", "1")
	state, _ := strconv.Atoi(stateStr)

	if id == 0 || name == "" {
		ReturnError(ctx, errcode.ErrInvalidRequest, "请输入正确信息")
		return
	}
	// 判断分类是否存在
	category, _ := modules.GetCategoryByName(name)
	if category.Id != 0 {
		ReturnError(ctx, errcode.ErrInvalidRequest, "分类已存在")
		return
	}
	// 更新数据库
	category, err := modules.UpdateCategory(&modules.UpdateCategoryDto{
		Id:    id,
		Name:  name,
		State: state,
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
	category, _ := modules.GetCategoryById(id)
	if category.State == 0 {
		ReturnError(ctx, errcode.ErrInvalidRequest, "分类已删除")
		return
	}
	category, err := modules.UpdateCategory(&modules.UpdateCategoryDto{
		Id:    id,
		State: 0,
	})
	if err != nil {
		ReturnError(ctx, errcode.ErrInvalidRequest, "删除失败")
		return
	}
	ReturnSuccess(ctx, 0, "删除成功", "")
}
