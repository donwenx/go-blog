package main

import (
	"blog/dao"
	"blog/model"
	"blog/router"
)

func main() {
	router := router.Router()
	dao.Db.AutoMigrate(model.User{}, model.Token{}, model.Category{}, model.Article{}) // 自动创建目录
	router.Run("127.0.0.1:8080")
}
