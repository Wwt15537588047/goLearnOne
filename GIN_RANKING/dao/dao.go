package dao

import (
	"gin_ranking/config"
	"gin_ranking/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", config.Mysqldb)
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}
	if Db.Error != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}

	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxIdleTime(time.Hour)
}
