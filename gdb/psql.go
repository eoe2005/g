package gdb

import (
	"fmt"

	"github.com/eoe2005/g/gconf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectPsql(dbconf *gconf.GDbYaml) *gorm.DB {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", dbconf.Host, dbconf.UserName, dbconf.UserPass, dbconf.DbName, dbconf.Port)), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
