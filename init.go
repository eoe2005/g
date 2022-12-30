package g

import (
	"flag"

	"github.com/eoe2005/g/gcache"
	"github.com/eoe2005/g/gconf"
	"github.com/eoe2005/g/gconfcenter"
	"github.com/eoe2005/g/gdb"
	"github.com/eoe2005/g/genv"
	"github.com/eoe2005/g/gmail"
	"github.com/eoe2005/g/gmq"
	"github.com/eoe2005/g/gstorage"
	"github.com/eoe2005/g/gweb"
)

var (
	localCall = []func(){}
)

func Use(call ...func()) {
	localCall = append(localCall, call...)
}
func initConfig() {
	_e := flag.String("e", "dev", "请输入运行环境")
	flag.Parse()
	genv.SetRunEnv(*_e)
	gconf.Load()
	gdb.Register(gconf.GetAppConf().Dbs)
	gmail.Register(gconf.GetAppConf().Mails)
	gcache.Register(gconf.GetAppConf().Caches)
	gconfcenter.Register(gconf.GetAppConf().Cfgs)
	gmq.Register(gconf.GetAppConf().Mqs)
	gweb.Register(gconf.GetAppConf().Web)
	gstorage.Register(gconf.GetAppConf().Storage)
}
