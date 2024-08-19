package dao

import (
	"gin-ranking/config"
	"gin-ranking/pkg/logger"
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
		logger.Error(map[string]interface{}{"database error": Db.Error})
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	Db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	Db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	Db.DB().SetConnMaxLifetime(time.Hour)
}
