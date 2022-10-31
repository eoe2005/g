package gconf

var (
	localConf = AppConf{}
)

type AppConf struct {
	Mqs    []*GMqYaml    `yaml:"mq"`
	Dbs    []*GDbYaml    `yaml:"db"`
	Caches []*GCacheYaml `yaml:"cache"`
	Cfgs   []*GCfgYaml   `yaml:"cfg"`
	Mails  []*GMailYaml  `yaml:"mail"`
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
	UserPass        string `yaml:"userpass"`
	UserName        string `yaml:"username"`
	DbName          string `yaml:"dbname"`
	MaxIdleCons     int    `yaml:"max_idle_cons"`
	MaxOpenCons     int    `yaml:"max_open_cons"`
	MaxLifetime     int64  `yaml:"max_lifetime"`
	MaxIdleLifetime int64  `yaml:"max_idle_lifetime"`
}

type GCacheYaml struct {
	Name     string `yaml:"name"`
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	UserPass string `yaml:"userpass"`
	DB       int    `yaml:"db"`
}

type GCfgYaml struct {
	Name      string `yaml:"name"`
	Driver    string `yaml:"driver"`
	AppID     string `yaml:"appid"`
	Host      string `yaml:"host"`
	UserName  string `yaml:"user"`
	Namespace string `yaml:"namespace"`
	Cluster   string `yaml:"cluster"`
}

type GMailYaml struct {
	Name     string `yaml:"name"`
	Smtp     string `yaml:"host"`
	Prot     int    `yaml:"port"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
	FromName string `yaml:"fromname"`
	IsSsl    bool   `yaml:"ssl"`
}
