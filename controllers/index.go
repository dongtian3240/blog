package controllers

import (
	"blog/common"
	"blog/models"
	"fmt"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index() {

	pager := common.NewPager()
	pager.PageNum = 1
	pager.PageSize = 10
	err := models.FindTopicListByPage(pager)
	if err != nil {
		this.Abort("501")
		return
	}
	fmt.Println(pager)
	this.Data["pager"] = pager
	user := this.GetSession("user")
	this.Data["user"] = user
	this.TplName = "index.tpl"
}
