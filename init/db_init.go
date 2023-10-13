package init

import (
	"fmt"

	"github.com/beego/beego"
	"github.com/beego/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type dbModel string

const (
	dbModelMaster dbModel = "master"
	dbModelSlave  dbModel = "slave"
)

type dbOpt struct {
	Model  dbModel
	DBName string
}

func (d dbOpt) String() string {
	return fmt.Sprintf("db_%s_%s", d.Model, d.DBName)
}

func dbInit(opts ...dbOpt) {
	orm.Debug = beego.AppConfig.String("runmode") == beego.DEV

	for i, opt := range opts {
		if i == 0 {
			orm.RegisterDataBase("default", "mysql", beego.AppConfig.String(opt.String()))
			if opt.Model == dbModelMaster {
				orm.RunSyncdb("default", false, orm.Debug)
			}
		}

		orm.RegisterDataBase(opt.DBName, "mysql", beego.AppConfig.String(opt.String()))
		if opt.Model == dbModelMaster {
			orm.RunSyncdb(opt.DBName, false, orm.Debug)
		}
	}
}
