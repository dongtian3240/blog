package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index() {

	user := this.GetSession("user")
	this.Data["user"] = user
	this.TplName = "index.tpl"
}
