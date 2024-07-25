package dao

import (
	"blog/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open(mysql.Open(config.Mysql), &gorm.Config{})
	Db = Db.Debug()

	if err != nil {
		logrus.WithError(err).Error("database error")
	}

	if Db.Error != nil {
		logrus.WithError(Db.Error).Error("database error")
	}
}
