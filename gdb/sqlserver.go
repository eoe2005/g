package gdb

import (
	"fmt"

	"github.com/eoe2005/g/gconf"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func connectSqlServer(dbconf *gconf.GDbYaml) *gorm.DB {
	db, err := gorm.Open(sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", dbconf.UserName, dbconf.UserPass, dbconf.Host, dbconf.Port, dbconf.DbName)), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
