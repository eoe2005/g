package gdb

import (
	"fmt"

	"github.com/eoe2005/g/gconf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectMysql(dbconf *gconf.GDbYaml) *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconf.UserName, dbconf.UserPass, dbconf.Host, dbconf.Port, dbconf.DbName)), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
