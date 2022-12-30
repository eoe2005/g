package gmonitor

import (
	"time"

	"gorm.io/gorm"
)

const msql_key = "metrics_key"

type gormMetrics struct {
}
type gormConf struct {
	Host   string
	DBName string
}
type gormOpt struct {
	start time.Time
	cmd   string
}

func NewGormPlugin(db *gorm.DB, host, dbname string) *gormMetrics {
	db.Set("gormConf", gormConf{
		Host:   host,
		DBName: dbname,
	})
	return &gormMetrics{}

}
func (s *gormMetrics) Name() string {
	return "gorm:metrics_mysql"
}
func (s *gormMetrics) Initialize(db *gorm.DB) (err error) {
	// before database operation
	db.Callback().Create().Before("gorm:create").Register("metrics_create_span", s.BeforeCallback("create"))
	db.Callback().Query().Before("gorm:query").Register("metrics_create_span", s.BeforeCallback("query"))
	db.Callback().Update().Before("gorm:update").Register("metrics_create_span", s.BeforeCallback("update"))
	db.Callback().Delete().Before("gorm:delete").Register("metrics_create_span", s.BeforeCallback("delete"))
	db.Callback().Row().Before("gorm:row").Register("metrics_create_span", s.BeforeCallback("row"))
	db.Callback().Raw().Before("gorm:raw").Register("metrics_create_span", s.BeforeCallback("raw"))

	// after database operation
	db.Callback().Create().After("gorm:create").Register("metrics_end_span", s.AfterCallback())
	db.Callback().Query().After("gorm:query").Register("metrics_end_span", s.AfterCallback())
	db.Callback().Update().After("gorm:update").Register("metrics_end_span", s.AfterCallback())
	db.Callback().Delete().After("gorm:delete").Register("metrics_end_span", s.AfterCallback())
	db.Callback().Row().After("gorm:row").Register("metrics_end_span", s.AfterCallback())
	db.Callback().Raw().After("gorm:raw").Register("metrics_end_span", s.AfterCallback())

	return
}

func (s *gormMetrics) BeforeCallback(operation string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		db.Set(msql_key, gormOpt{start: time.Now(), cmd: operation})
	}

}

func (s *gormMetrics) AfterCallback() func(db *gorm.DB) {
	return func(db *gorm.DB) {
		optObj, _ := db.Get(msql_key)
		opt, ok := optObj.(gormOpt)
		if !ok {
			// fmt.Println("mysqlErr1", optObj)
			return
		}
		confObj, _ := db.Get("gormConf")
		conf, ok := confObj.(gormConf)
		if !ok {
			// fmt.Println("mysqlErr2", confObj)
			return
		}
		tags := map[string]string{
			"db_host": conf.Host,
			"command": opt.cmd,
			"db_name": conf.DBName,
		}
		Counter("mysql_total", "mysql_total", 1, tags)
		Summary("mysql_cost", "mysql_cost", float64(time.Now().Sub(opt.start)), tags)
	}
}
