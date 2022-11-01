package g

import (
	"github.com/eoe2005/g/gcache"
	"github.com/eoe2005/g/gconf"
	"github.com/eoe2005/g/gconfcenter"
	"github.com/eoe2005/g/gdb"
	"github.com/eoe2005/g/gmail"
	"github.com/eoe2005/g/gmq"
)

func initConfig() {
	gconf.Load()
	gdb.Register(gconf.GetAooConf().Dbs)
	gmail.Register(gconf.GetAooConf().Mails)
	gcache.Register(gconf.GetAooConf().Caches)
	gconfcenter.Register(gconf.GetAooConf().Cfgs)
	gmq.Register(gconf.GetAooConf().Mqs)
}
