package gdb

import (
	"time"

	"github.com/eoe2005/g/gconf"
	"github.com/eoe2005/g/gmonitor"
	"gorm.io/gorm"
)

var (
	_localDbMap  = map[string]*gorm.DB{}
	_localDbConf = map[string]*gconf.GDbYaml{}
)

func Register(dbList []*gconf.GDbYaml) {
	for _, dbConf := range dbList {
		_localDbConf[dbConf.Name] = dbConf

	}
}

func GetDB(key string, model ...interface{}) *gorm.DB {
	if r, ok := _localDbMap[key]; ok {
		if len(model) > 0 {
			return r.Model(model[0])
		}
		return r.Debug()
	}
	r := getDbCon(key)
	if len(model) > 0 {
		return r.Model(model[0]).Debug()
	}

	return r.Debug()

}

func getDbCon(name string) *gorm.DB {
	if dbConf, ok := _localDbConf[name]; ok {
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
		case "es":
			initEs(dbConf)
		}
		if con != nil {
			con.Use(gmonitor.NewGormPlugin(con, dbConf.Host, dbConf.DbName))
			db, e := con.DB()
			if e != nil {
				panic("链接数据库失败")
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
			return con
		} else {
			panic("不支持的数据库类型")
		}
	}
	panic("没有配置数据库")
}
