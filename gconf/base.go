package gconf

var (
	localConf = AppConf{}
)

func GetAppConf() AppConf {
	return localConf
}

type AppConf struct {
	Mqs     []*GMqYaml      `yaml:"mq"`
	Dbs     []*GDbYaml      `yaml:"db"`
	Caches  []*GCacheYaml   `yaml:"cache"`
	Cfgs    []*GCfgYaml     `yaml:"cfg"`
	Mails   []*GMailYaml    `yaml:"mail"`
	Web     *GWebYaml       `yaml:"web"`
	Storage []*GStorageYaml `yaml:"storage"`
	Log     *GLogYaml       `yaml:"log"`
	Base    []*BaseYaml     `yaml:"app"`
	Wx      []*WxYaml       `yaml:"wx"`
}
type WxYaml struct {
	Name      string `yaml:"name"`
	AppID     string `yaml:"appid"`
	Scope     string `yaml:"scope"`
	AppSecret string `yaml:"secret"`
	RetureUrl string `yaml:"return_url"`
}
type BaseYaml struct {
	Name string            `yaml:"name"`
	Data map[string]string `yaml:"data"`
}
type GLogYaml struct {
	Dir         string `yaml:"log_dir"`
	SplitType   string `yaml:"split_type"`
	MaxFileSize int64  `yaml:"file_size"`
}

type GStorageYaml struct {
	Name   string `yaml:"name"`
	Driver string `yaml:"driver"`
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
	Bucket string `yaml:"bucket"`
	Domain string `yaml:"domain"`
}
type GMqYaml struct {
	Name     string `yaml:"name"`
	Driver   string `yaml:"driver"`
	Hosts    string `yaml:"hosts"`
	Topic    string `yaml:"topic"`
	RefRedis string `yaml:"ref"`
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

type GDingDingYaml struct {
	Name  string `yaml:"name"`
	Token string `yaml:"token"`
}
type GWebYaml struct {
	MiddleWare []*GWebMiddleWareYaml `'yaml:"middleware"`
}

type GWebMiddleWareYaml struct {
	Name       string `yaml:"name"`
	Driver     string `yaml:"driver"`
	Ref        string `yaml:"ref_redis"` // driver = session_redis | session_redis_cluster
	AuthKey    string `yaml:"auth_key"`  // driver = gwt
	IsHeader   bool   `yaml:"is_header"`
	IsCookie   bool   `yaml:"is_cookie"`
	SendName   string `yaml:"outer_name"`
	TimeOut    int    `yaml:"timeout"`
	Key        string `yaml:"key"`         // driver = aes
	PublicKey  string `yaml:"public_key"`  // driver= rsa
	PrivateKey string `yaml:"private_key"` //driver = rsa
}
