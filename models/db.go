package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var ormObj orm.Ormer

//Connect DB ORM PostgreSQL
func ConnDB() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=m435tr0 dbname=go-gin-postgre host=localhost sslmode=disable")
	orm.RegisterModel(new(Users))
	ormObj = orm.NewOrm()
}

func GetOrmObj() orm.Ormer {
	return ormObj
}
