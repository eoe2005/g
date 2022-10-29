package gdb

import (
	"fmt"

	"github.com/eoe2005/g/gconf"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func connectClickHouse(dbconf *gconf.GDbYaml) *gorm.DB {
	db, err := gorm.Open(clickhouse.Open(fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s&read_timeout=10&write_timeout=20", dbconf.Host, dbconf.Port, dbconf.DbName, dbconf.UserName, dbconf.UserPass)), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
