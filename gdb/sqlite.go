package gdb

import (
	"github.com/eoe2005/g/gconf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectSqlite(dbconf *gconf.GDbYaml) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dbconf.Host), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
