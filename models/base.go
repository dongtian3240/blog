package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/blog?charset=utf8")
	orm.RegisterModel(new(Topic), new(User), new(Label), new(Comment), new(TopicLabel))
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
}
