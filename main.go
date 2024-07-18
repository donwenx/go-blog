package main

import (
	"blog/dao"
	"blog/modules"
	"blog/router"
)

func main() {
	router := router.Router()
	dao.Db.AutoMigrate(modules.User{}, modules.Token{}, modules.Category{}) // 自动创建目录
	router.Run("127.0.0.1:8080")
}
