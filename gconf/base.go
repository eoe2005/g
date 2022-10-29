package gconf

type AppConf struct {
	Mqs    []*GMqYaml   `yaml:"mq"`
	Dbs    []*GDbYaml   `yaml:"db"`
	Caches []*GConfYaml `yaml:"cache"`
}
type GMqYaml struct {
	Name   string `yaml:"name"`
	Driver string `yaml:"driver"`
	Hosts  string `yaml:"hosts"`
	Topic  string `yaml:"topic"`
}

type GDbYaml struct {
	Name            string `yaml:"name"`
	Driver          string `yaml:"driver"`
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	UserPass        string `yaml:"pass"`
	UserName        string `yaml:"user"`
	DbName          string `yaml:"dbname"`
	MaxIdleCons     int    `yaml:"max_idle_cons"`
	MaxOpenCons     int    `yaml:"max_open_cons"`
	MaxLifetime     int64  `yaml:"max_lifetime"`
	MaxIdleLifetime int64  `yaml:"max_idle_lifetime"`
}

type GConfYaml struct {
	Name     string `yaml:"name"`
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	UserPass string `yaml:"pass"`
	DB       int    `yaml:"db"`
}
