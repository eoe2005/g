package gconf

type AppConf struct {
	Kafkas       []*GKafkaYaml    `yaml:"kafka"`
	Dbs          []*GDbYaml       `yaml:"db"`
	RedisCluster []*GRedisYaml    `yaml:"redis_cluter"`
	Redis        []*GRedisYaml    `yaml:"redis"`
	Memcache     []*GMemcacheYaml `yaml:"memcache"`
}
type GKafkaYaml struct {
	Name  string `yaml:"name"`
	Hosts string `yaml:"hosts"`
	Topic string `yaml:"topic"`
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

type GRedisYaml struct {
}
type GMemcacheYaml struct {
}
