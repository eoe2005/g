package gdb

import (
	"time"

	"github.com/eoe2005/g/gconf"
	"gorm.io/gorm"
)

var (
	_localDbMap = map[string]*gorm.DB{}
)

func Register(dbList []*gconf.GDbYaml) {
	for _, dbConf := range dbList {
		var con *gorm.DB
		switch dbConf.Driver {
		case "mysql":
			con = connectMysql(dbConf)
		case "sqlite":
			con = connectSqlite(dbConf)
		case "psql":
			con = connectPsql(dbConf)
		case "sqlserver":
			con = connectSqlServer(dbConf)
		case "clickhouse":
			con = connectClickHouse(dbConf)
		}
		if con != nil {
			db, e := con.DB()
			if e != nil {
				continue
			}
			if dbConf.MaxIdleCons > 0 {
				db.SetMaxIdleConns(dbConf.MaxIdleCons)
			}
			if dbConf.MaxOpenCons > 0 {
				db.SetMaxOpenConns(dbConf.MaxOpenCons)
			}
			if dbConf.MaxLifetime > 0 {
				db.SetConnMaxLifetime(time.Duration(dbConf.MaxLifetime) * time.Second)
			}
			if dbConf.MaxIdleLifetime > 0 {
				db.SetConnMaxIdleTime(time.Duration(dbConf.MaxIdleLifetime) * time.Second)
			}
			_localDbMap[dbConf.Name] = con
		}
	}
}

func GetDB(key string, model ...interface{}) *gorm.DB {
	if r, ok := _localDbMap[key]; ok {
		if len(model) > 0 {
			return r.Model(model[0])
		}
		return r
	}
	return nil
}
