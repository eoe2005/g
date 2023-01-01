package gdb

import (
	"context"
	"fmt"

	"github.com/eoe2005/g/gconf"
	"github.com/eoe2005/g/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func connectMysql(dbconf *gconf.GDbYaml) *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconf.UserName, dbconf.UserPass, dbconf.Host, dbconf.Port, dbconf.DbName)), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.New(glog.GetLog("sql"), logger.Config{}),
	})
	if err != nil {
		glog.Error(context.Background(), "链接数据库失败[%s]%s", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconf.UserName, dbconf.UserPass, dbconf.Host, dbconf.Port, dbconf.DbName), err.Error())
		return nil
	}
	return db
}
