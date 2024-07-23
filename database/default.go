package database

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
	"src/config"
)

func init() {
	dsn := config.DBConnectString()
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dsn)
}
